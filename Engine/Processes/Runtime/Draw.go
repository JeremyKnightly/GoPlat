package runtime

import (
	"GoPlat/engine/camera"
	"GoPlat/engine/collision"
	"GoPlat/engine/movement"
	"GoPlat/gameComponents/animations"
	levels "GoPlat/gameComponents/levels"
	sprites "GoPlat/gameComponents/sprites"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

func DrawLevelFirstDraw(level *levels.Level, screen *ebiten.Image, cam *camera.Camera) {
	tileSize := 16
	for idx, layer := range level.Layers {
		if !layer.FirstDraw {
			continue
		}
		DrawLayer(level, idx, screen, cam, tileSize)
	}
}

func DrawLevelSecondDraw(level *levels.Level, screen *ebiten.Image, cam *camera.Camera) {
	tileSize := 16
	for idx, layer := range level.Layers {
		if layer.FirstDraw {
			continue
		}
		DrawLayer(level, idx, screen, cam, tileSize)
	}
}

func DrawLayer(level *levels.Level, layerIdx int, screen *ebiten.Image, cam *camera.Camera, tileSize int) {
	mapDrawOptions := ebiten.DrawImageOptions{}
	layer := level.Layers[layerIdx]
	for index, id := range layer.Data {
		if id == 0 {
			continue
		}

		x := index % int(layer.Width)
		y := index / int(layer.Width)

		x *= tileSize
		y *= tileSize

		srcX := (id - 1) % 15
		srcY := (id - 1) / 15

		srcX *= tileSize
		srcY *= tileSize

		mapDrawOptions.GeoM.Translate(float64(x), float64(y))
		mapDrawOptions.GeoM.Translate(cam.X, cam.Y)

		screen.DrawImage(level.TilemapImage.SubImage(image.Rect(srcX, srcY, srcX+16, srcY+16)).(*ebiten.Image),
			&mapDrawOptions,
		)

		mapDrawOptions.GeoM.Reset()
	}
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
	currentAnimation.AnimationComplete = false
}

func DrawPlayer(player *sprites.Player, screen *ebiten.Image, lvl *levels.Level, cam *camera.Camera) {
	playerDrawOptions := ebiten.DrawImageOptions{}

	if player.IsIdle {
		handleIdleDraw(player, screen, &playerDrawOptions, cam)
	} else {
		handleActiveDraw(player, screen, lvl, cam)
	}
}

func handleActiveDraw(player *sprites.Player, screen *ebiten.Image, lvl *levels.Level, cam *camera.Camera) {
	playerDrawOptions := ebiten.DrawImageOptions{}
	effectDrawOptions := ebiten.DrawImageOptions{}
	stopCompletedAnimations(player)
	hasAnimationEffect := player.ActionAnimations[player.CurrentAnimationIndex].HasEffect

	if !hasAnimationEffect {
		prepSpriteNoEffect(player, &playerDrawOptions)
	} else {
		prepSpriteWithEffect(player, &playerDrawOptions, &effectDrawOptions)
	}

	currentAnimation := player.ActionAnimations[player.CurrentAnimationIndex]
	currentFrame, frameVector, canCancel := currentAnimation.AnimateAction()
	player.CanAnimationCancel = canCancel

	if !player.IsMovingRight {
		frameVector.DeltaX *= -1
	}
	finalVec := movement.HandleAnimationVectorCalculations(lvl, player, frameVector)
	newPosition := finalVec.PlayerMove(player.X, player.Y, player.IsMovingRight)
	playerDrawOptions.GeoM.Translate(cam.X, cam.Y)
	//if invalid move, draw frame as is and return
	if !collision.IsValidMove(lvl, player, finalVec) {
		screen.DrawImage(currentFrame, &playerDrawOptions)
		return
	}

	player.X = newPosition.DeltaX
	player.Y = newPosition.DeltaY
	playerDrawOptions.GeoM.Translate(finalVec.DeltaX, finalVec.DeltaY)

	if hasAnimationEffect {
		effectFrame := currentAnimation.Effect.Frames[currentAnimation.CurrentFrameIndex]
		effectDrawOptions.GeoM.Translate(cam.X, cam.Y)
		screen.DrawImage(effectFrame, &effectDrawOptions)
	}

	screen.DrawImage(currentFrame, &playerDrawOptions)
}

func getAnimationMaxFrameWidth(p *sprites.Player) float64 {
	if p.IsIdle {
		return p.IdleAnimation.MaxFrameWidth
	} else {
		return p.ActionAnimations[p.CurrentAnimationIndex].MaxFrameWidth
	}
}

func getEffectMaxFrameWidth(effect *animations.Effect) float64 {
	return effect.MaxFrameWidth
}

func adjustDrawOptionsForLeftMove(options *ebiten.DrawImageOptions, maxWidth float64) {
	options.GeoM.Scale(-1, 1)
	options.GeoM.Translate(maxWidth, 0)
}

func prepSpriteNoEffect(p *sprites.Player, options *ebiten.DrawImageOptions) {
	if !p.IsMovingRight {
		frameWidth := getAnimationMaxFrameWidth(p)
		adjustDrawOptionsForLeftMove(options, frameWidth)
	}
	options.GeoM.Translate(p.X, p.Y)
}

func prepSpriteWithEffect(p *sprites.Player, options *ebiten.DrawImageOptions, effectOptions *ebiten.DrawImageOptions) {
	effect := p.ActionAnimations[p.CurrentAnimationIndex].Effect
	if !p.IsMovingRight {
		effectWidth := getEffectMaxFrameWidth(&effect)
		frameWidth := getAnimationMaxFrameWidth(p)
		adjustDrawOptionsForLeftMove(effectOptions, effectWidth+frameWidth)
		adjustDrawOptionsForLeftMove(options, frameWidth)
		//effectOptions.GeoM.Translate(-effect.Offset, 0)
	}

	options.GeoM.Translate(p.X, p.Y)
	effectOptions.GeoM.Translate(p.X, p.Y)
}

func handleIdleDraw(player *sprites.Player, screen *ebiten.Image, options *ebiten.DrawImageOptions, cam *camera.Camera) {
	prepSpriteNoEffect(player, options)
	currentFrame := player.IdleAnimation.Animate()
	options.GeoM.Translate(cam.X, cam.Y)
	screen.DrawImage(currentFrame, options)
}
