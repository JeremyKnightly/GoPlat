package runtime

import (
	"GoPlat/gameComponents/animations"
	levels "GoPlat/gameComponents/levels"
	sprites "GoPlat/gameComponents/sprites"
	"GoPlat/engine/collision"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

func DrawLevel(level *levels.Level, screen *ebiten.Image) {
	mapDrawOptions := ebiten.DrawImageOptions{}
	//loop over layers
	tileSize := 16
	for _, layer := range level.Layers {
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

			screen.DrawImage(level.TilemapImage.SubImage(image.Rect(srcX, srcY, srcX+16, srcY+16)).(*ebiten.Image),
				&mapDrawOptions,
			)

			mapDrawOptions.GeoM.Reset()
		}
	}

}

func stopCompletedAnimations(player *sprites.Player) {
	if player.ActionAnimations[player.CurrentAnimationIndex].AnimationComplete {
		player.IsAnimationLocked = false
		player.ActionAnimations[player.CurrentAnimationIndex].AnimationComplete = false
		player.CurrentAnimationIndex = 0
	}
}

func DrawPlayer(player *sprites.Player, screen *ebiten.Image, lvl *levels.Level) {
	playerDrawOptions := ebiten.DrawImageOptions{}

	if player.IsIdle {
		handleIdleDraw(player, screen, &playerDrawOptions)
	} else {
		handleActiveDraw(player, screen, lvl)
	}
}

func handleActiveDraw(player *sprites.Player, screen *ebiten.Image, lvl *levels.Level) {
	playerDrawOptions := ebiten.DrawImageOptions{}
	effectDrawOptions := ebiten.DrawImageOptions{}
	stopCompletedAnimations(player)
	hasAnimationEffect := player.ActionAnimations[player.CurrentAnimationIndex].HasEffect

	if !hasAnimationEffect {
		prepSpriteNoEffect(player, &playerDrawOptions)
	} else {
		prepSpriteWithEffect(player,&playerDrawOptions, &effectDrawOptions)
	}

	currentAnimation := player.ActionAnimations[player.CurrentAnimationIndex]
	currentFrame, frameVector, canCancel := currentAnimation.AnimateAction()
	player.CanAnimationCancel = canCancel

	newPosition := frameVector.PlayerMove(player.X, player.Y, player.IsMovingRight)
	validMove := collision.IsValidMove(lvl, player, newPosition)
	
	if validMove {
		player.X = newPosition.DeltaX
		player.Y = newPosition.DeltaY
		playerDrawOptions.GeoM.Translate(frameVector.DeltaX, frameVector.DeltaY)
	}

	if hasAnimationEffect{
		effectFrame := currentAnimation.Effect.Frames[currentAnimation.CurrentFrameIndex]
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
	options.GeoM.Scale(-1,1)
	options.GeoM.Translate(maxWidth, 0)
} 

func prepSpriteNoEffect(p *sprites.Player, options *ebiten.DrawImageOptions) {
	if !p.IsMovingRight {
		frameWidth := getAnimationMaxFrameWidth(p)
		adjustDrawOptionsForLeftMove(options,frameWidth)
	}
	options.GeoM.Translate(p.X, p.Y)
}

func prepSpriteWithEffect(p *sprites.Player, options *ebiten.DrawImageOptions, effectOptions *ebiten.DrawImageOptions) {
	effect := p.ActionAnimations[p.CurrentAnimationIndex].Effect
	if !p.IsMovingRight {
		effectWidth := getEffectMaxFrameWidth(&effect)
		frameWidth := getAnimationMaxFrameWidth(p)
		adjustDrawOptionsForLeftMove(effectOptions,effectWidth + frameWidth)
		adjustDrawOptionsForLeftMove(options,frameWidth)
		effectOptions.GeoM.Translate(-effect.Offset, 0)
	} else {
		effectOptions.GeoM.Translate(effect.Offset, 0)
	}
	options.GeoM.Translate(p.X, p.Y)
	effectOptions.GeoM.Translate(p.X, p.Y)
}

func handleIdleDraw(player *sprites.Player, screen *ebiten.Image, options *ebiten.DrawImageOptions) {
	prepSpriteNoEffect(player, options)
	currentFrame := player.IdleAnimation.Animate()
	screen.DrawImage(currentFrame, options)
}