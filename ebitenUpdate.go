package main

import (
	"GoPlat/Engine/camera"
	"GoPlat/Engine/collision"
	movement "GoPlat/engine/movement"
	"GoPlat/gameComponents/animations"
	levels "GoPlat/gameComponents/levels"
	sprites "GoPlat/gameComponents/sprites"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Update() error {
	g.currentLevel = g.levels[g.currentLevelIndex]

	movement.HandleMovementCalculations(g.Player, g.controls, g.currentLevel)
	g.Player.Physics.UpdatePhysics(1 / ebiten.ActualFPS())
	for !collision.IsValidMove(g.currentLevel, g.Player) {
		collision.ResolveCollisions(g.currentLevel, g.Player)
	}
	g.setCameraPosition()
	g.setPlayerFrame()

	return nil
}

func (g *Game) setCameraPosition() {
	g.camera.FollowTarget(g.Player.Physics.Position.X+g.tileSize/2, g.Player.Physics.Position.Y+g.tileSize/2, g.screenWidth, g.screenHeight)
	g.camera.Constrain(g.currentLevel.Layers[0].Height,
		g.currentLevel.Layers[0].Width,
		g.screenWidth,
		g.screenHeight,
	)
}

func (g *Game) setPlayerFrame() {

	g.Player.Frame.ImageOptions = ebiten.DrawImageOptions{}
	g.Player.Frame.EffectOptions = ebiten.DrawImageOptions{}
	if g.Player.IsIdle {
		setPlayerIdleFrameAndPosition(g.Player, g.camera)
	} else {
		setPlayerActiveFrameAndPosition(g.Player, g.currentLevel, g.camera)
	}
}

func setPlayerIdleFrameAndPosition(player *sprites.Player, cam *camera.Camera) {
	prepSpriteNoEffect(player)
	player.Frame.ImageToDraw = player.IdleAnimation.Animate()
	player.Frame.ImageOptions.GeoM.Translate(cam.X, cam.Y)
}

func prepSpriteNoEffect(player *sprites.Player) {
	if !player.IsMovingRight {
		adjustDrawOptionsForLeftMove(
			&player.Frame.ImageOptions,
			getAnimationMaxFrameWidth(player),
		)
	}
	player.Frame.ImageOptions.GeoM.Translate(player.Physics.Position.X, player.Physics.Position.Y)
}

func prepSpriteWithEffect(p *sprites.Player) {
	effect := p.ActionAnimations[p.CurrentAnimationIndex].Effect
	if !p.IsMovingRight {
		effectWidth := getEffectMaxFrameWidth(&effect)
		frameWidth := getAnimationMaxFrameWidth(p)
		adjustDrawOptionsForLeftMove(&p.Frame.EffectOptions, effectWidth+frameWidth)
		adjustDrawOptionsForLeftMove(&p.Frame.ImageOptions, frameWidth)
	}

	p.Frame.ImageOptions.GeoM.Translate(p.Physics.Position.X, p.Physics.Position.Y)
	p.Frame.EffectOptions.GeoM.Translate(p.Physics.Position.X, p.Physics.Position.Y)
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

func stopCompletedAnimations(player *sprites.Player) {
	currentAnimation := player.ActionAnimations[player.CurrentAnimationIndex]
	if !currentAnimation.AnimationComplete {
		return
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

func setPlayerActiveFrameAndPosition(player *sprites.Player, lvl *levels.Level, cam *camera.Camera) {
	stopCompletedAnimations(player)
	prepActiveSprite(player)

	currentAnimation := player.ActionAnimations[player.CurrentAnimationIndex]

	currentFrame, canCancel := currentAnimation.AnimateAction()
	player.Frame.HasEffect = currentAnimation.HasEffect
	player.Frame.ImageToDraw = currentFrame
	player.CanAnimationCancel = canCancel

	if player.Frame.HasEffect && !currentAnimation.AnimationComplete {
		player.Frame.EffectImageToDraw = currentAnimation.Effect.Frames[currentAnimation.CurrentFrameIndex]
	}
	translatePlayerDrawOptions(player, cam.X, cam.Y)
}

func translatePlayerDrawOptions(p *sprites.Player, DeltaX float64, DeltaY float64) {
	p.Frame.ImageOptions.GeoM.Translate(DeltaX, DeltaY)
	if p.Frame.HasEffect {
		p.Frame.EffectOptions.GeoM.Translate(DeltaX, DeltaY)
	}
}
