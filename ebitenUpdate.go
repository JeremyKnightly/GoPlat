package main

import (
	"GoPlat/engine/collision"
	movement "GoPlat/engine/movement"
	"GoPlat/gameComponents/animations"
	levels "GoPlat/gameComponents/levels"
	sprites "GoPlat/gameComponents/sprites"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Update() error {
	if g.gameState == 0 { //start screen
		g.StartScreen()
		g.currentBGMIdx = 1
	} else if g.gameState == 1 { //game
		g.PlayGame()
		g.currentBGMIdx = 0
	} else if g.gameState == 2 { //credits
		g.RollCredits()
		g.currentBGMIdx = 2
	} else if g.gameState == 3 { //High Scores
		g.ShowScores()
		g.currentBGMIdx = 3
	} else if g.gameState == 4 { //options
		g.ShowOptions()
		g.currentBGMIdx = 4
	}

	g.loopBGM(g.currentBGMIdx)

	return nil
}

func (g *Game) StartScreen() {
	g.currentLevel = g.startScreen
	g.setPlayerPositionWithInput()
	g.checkPlayerClicks()
	g.setPlayerFrame()
	g.playSFX()
}

func (g *Game) PlayGame() {
	if g.Player.CurrentCheckpointIndex == 9999 {
		g.GoToLevel(g.currentLevelIndex + 1)
	}
	g.currentLevel = g.levels[g.currentLevelIndex]

	g.setPlayerPositionWithInput()
	g.setCameraPosition()
	g.setPlayerFrame()
	g.playSFX()
}

func (g *Game) RollCredits() {
	g.currentLevel = g.startScreen
}

func (g *Game) ShowScores() {
	g.currentLevel = g.startScreen
}

func (g *Game) ShowOptions() {
	g.currentLevel = g.startScreen
}

func (g *Game) GoToLevel(levelNum int) {
	g.currentLevelIndex = levelNum
	if g.currentLevelIndex > len(g.levels) {
		g.currentLevelIndex = 0
	}
	g.Player.CurrentCheckpointIndex = 0
	x, y := g.levels[g.currentLevelIndex].GetCheckpointXY(g.Player.CurrentCheckpointIndex)
	g.Player.X = x
	g.Player.Y = y
}

func (g *Game) checkPlayerClicks() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		changeState, stateInt := collision.MenuStateChange(g.currentLevel)
		if changeState {
			g.gameState = stateInt
			g.Player.Status.Score.SetGameStartTime()
			g.GoToLevel(0)
		}
	}
}

func (g *Game) setPlayerPositionWithInput() {
	inputVector := movement.HandleMovementCalculations(g.Gamepad, g.Player, g.controls, g.currentLevel)

	collision.EnsureValidMove(g.currentLevel, g.Player, inputVector)
}

func (g *Game) setCameraPosition() {
	g.camera.FollowTarget(g.Player.X+g.tileSize/2, g.Player.Y+g.tileSize/2, g.defaultScreenWidth, g.defaultScreenHeight)
	g.camera.Constrain(g.currentLevel.Layers[0].Height*g.tileSize,
		g.currentLevel.Layers[0].Width*g.tileSize,
		g.defaultScreenWidth,
		g.defaultScreenHeight,
	)
}

func (g *Game) setPlayerFrame() {

	g.Player.Frame.ImageOptions = ebiten.DrawImageOptions{}
	g.Player.Frame.EffectOptions = ebiten.DrawImageOptions{}
	if g.Player.IsIdle && !g.Player.IsDead {
		setPlayerIdleFrameAndPosition(g.Player)
	} else {
		g.setPlayerActiveFrame(g.Player, g.currentLevel)
	}
	translatePlayerDrawOptions(g.Player, g.Player.X+g.camera.X, g.Player.Y+g.camera.Y)
}

func setPlayerIdleFrameAndPosition(player *sprites.Player) {
	player.Frame.ImageToDraw = player.IdleAnimation.Animate()
	prepSpriteNoEffect(player)
}

func prepSpriteNoEffect(player *sprites.Player) {
	if !player.IsMovingRight {
		adjustDrawOptionsForLeftMove(
			&player.Frame.ImageOptions,
			getAnimationMaxFrameWidth(player),
		)
	}
}

func prepSpriteWithEffect(p *sprites.Player) {
	effect := p.ActionAnimations[p.CurrentAnimationIndex].Effect
	if !p.IsMovingRight {
		effectWidth := getEffectMaxFrameWidth(&effect)
		frameWidth := getAnimationMaxFrameWidth(p)
		adjustDrawOptionsForLeftMove(&p.Frame.EffectOptions, effectWidth+frameWidth)
		adjustDrawOptionsForLeftMove(&p.Frame.ImageOptions, frameWidth)
	}
}

func getAnimationMaxFrameWidth(player *sprites.Player) float64 {
	if player.IsIdle {
		return player.IdleAnimation.MaxFrameWidth
	} else {
		return player.ActionAnimations[player.CurrentAnimationIndex].MaxFrameWidth
	}
}

func adjustDrawOptionsForLeftMove(options *ebiten.DrawImageOptions, maxWidth float64) {
	options.GeoM.Scale(-1, 1)
	options.GeoM.Translate(maxWidth, 0)
}

func (g *Game) stopCompletedAnimations(player *sprites.Player) {
	currentAnimation := player.ActionAnimations[player.CurrentAnimationIndex]
	if !currentAnimation.AnimationComplete {
		return
	}

	//if player dead
	if player.CurrentAnimationIndex == 8 {
		player.Resurrect(g.levels[g.currentLevelIndex])
	}
	//if I shouldn't loop, set index to walk
	if !currentAnimation.LoopAnimation {
		player.CurrentAnimationIndex = 0
	}
	player.IsAnimationLocked = false
	player.Frame.HasEffect = false
	currentAnimation.AnimationComplete = false
}

func getEffectMaxFrameWidth(effect *animations.Effect) float64 {
	return effect.MaxFrameWidth
}

func prepActiveSprite(player *sprites.Player) {
	hasAnimationEffect := player.ActionAnimations[player.CurrentAnimationIndex].HasEffect

	if !hasAnimationEffect {
		prepSpriteNoEffect(player)
	} else {
		prepSpriteWithEffect(player)
		player.Frame.HasEffect = true
	}
}

func (g *Game) setPlayerActiveFrame(player *sprites.Player, lvl *levels.Level) {
	g.stopCompletedAnimations(player)
	prepActiveSprite(player)

	currentAnimation := player.ActionAnimations[player.CurrentAnimationIndex]

	frameDurationInSeconds := currentAnimation.FrameDuration.Seconds()
	currentAnimation.TicksPerFrame = frameDurationInSeconds * ebiten.ActualTPS()

	currentFrame, frameVector, canCancel, _ := currentAnimation.AnimateAction()
	player.Frame.HasEffect = currentAnimation.HasEffect
	player.Frame.ImageToDraw = currentFrame
	player.Frame.EffectOffset = currentAnimation.Effect.OffsetX
	player.Frame.EffectOffsetOneWay = currentAnimation.Effect.OffsetOneWay
	player.Frame.EffectOffsetRight = currentAnimation.Effect.OffsetRightOnly
	player.CanAnimationCancel = canCancel

	if player.Frame.HasEffect && !currentAnimation.AnimationComplete {
		player.Frame.EffectImageToDraw = currentAnimation.Effect.Frames[currentAnimation.CurrentFrameIndex]
	}

	if !player.IsMovingRight {
		frameVector.DeltaX *= -1
		player.Frame.EffectOffset *= -1
	}

	collision.EnsureValidMove(lvl, player, frameVector)
}

func translatePlayerDrawOptions(p *sprites.Player, DeltaX float64, DeltaY float64) {
	p.Frame.ImageOptions.GeoM.Translate(DeltaX, DeltaY)
	if p.Frame.HasEffect {
		p.Frame.EffectOptions.GeoM.Translate(DeltaX, DeltaY)
		if p.Frame.EffectOffsetOneWay {
			if p.Frame.EffectOffsetRight && p.IsMovingRight {
				p.Frame.EffectOptions.GeoM.Translate(p.Frame.EffectOffset, 0)
			} else if !p.Frame.EffectOffsetRight && !p.IsMovingRight {
				p.Frame.EffectOptions.GeoM.Translate(p.Frame.EffectOffset, 0)
			}
			return
		}
		p.Frame.EffectOptions.GeoM.Translate(p.Frame.EffectOffset, 0)
	}
}
