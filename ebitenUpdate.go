package main

import (
<<<<<<< HEAD
	"GoPlat/engine/collision"
=======
	"GoPlat/Engine/camera"
	"GoPlat/Engine/collision"
>>>>>>> 16f3c53bd428513a1c986c0fe9c23443b6469d9c
	movement "GoPlat/engine/movement"
	"GoPlat/gameComponents/animations"
	levels "GoPlat/gameComponents/levels"
	sprites "GoPlat/gameComponents/sprites"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Update() error {
	if g.Player.CurrentCheckpointIndex == 9999 {
		g.GoNextLevel()
	}
	g.currentLevel = g.levels[g.currentLevelIndex]

<<<<<<< HEAD
	g.setPlayerPositionWithInput()
	g.setCameraPosition()
	g.setPlayerFrame()
	g.playSFX()
	g.loopBGM(g.currentBGMIdx)
=======
	movement.HandleMovementCalculations(g.Player, g.controls, g.currentLevel)
	g.Player.Physics.UpdatePhysics(1 / ebiten.ActualFPS())
	//println(g.Player.Physics.Velocity.X, "  ", g.Player.Physics.Velocity.Y, "\n")
	for !collision.IsValidMove(g.currentLevel, g.Player) {
		collision.ResolveCollisions(g.currentLevel, g.Player)
	}
	//println(g.Player.Physics.Position.X, " , ", g.Player.Physics.Position.Y)
	g.setCameraPosition()
	g.setPlayerFrame()
>>>>>>> 16f3c53bd428513a1c986c0fe9c23443b6469d9c

	return nil
}

<<<<<<< HEAD
func (g *Game) GoNextLevel() {
	g.currentLevelIndex++
	if g.currentLevelIndex > len(g.levels) {
		g.currentLevelIndex = 0
	}
	g.Player.CurrentCheckpointIndex = 0
	x, y := g.levels[g.currentLevelIndex].GetCheckpointXY(g.Player.CurrentCheckpointIndex)
	g.Player.X = x
	g.Player.Y = y
}

func (g *Game) setPlayerPositionWithInput() {
	inputVector := movement.HandleMovementCalculations(g.Player, g.controls, g.currentLevel)

	collision.EnsureValidMove(g.currentLevel, g.Player, inputVector)
}

=======
>>>>>>> 16f3c53bd428513a1c986c0fe9c23443b6469d9c
func (g *Game) setCameraPosition() {
	g.camera.FollowTarget(g.Player.X+g.tileSize/2, g.Player.Y+g.tileSize/2, g.screenWidth, g.screenHeight)
	g.camera.Constrain(g.currentLevel.Layers[0].Height*g.tileSize,
		g.currentLevel.Layers[0].Width*g.tileSize,
		g.screenWidth,
		g.screenHeight,
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
	prepSpriteNoEffect(player)
	player.Frame.ImageToDraw = player.IdleAnimation.Animate()
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
<<<<<<< HEAD

	if !player.IsMovingRight {
		frameVector.DeltaX *= -1
		player.Frame.EffectOffset *= -1
	}

	collision.EnsureValidMove(lvl, player, frameVector)
=======
	translatePlayerDrawOptions(player, cam.X, cam.Y)
>>>>>>> 16f3c53bd428513a1c986c0fe9c23443b6469d9c
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
