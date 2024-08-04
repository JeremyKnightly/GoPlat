package runtime

import (
	levels "GoPlat/components/levels"
	sprites "GoPlat/components/sprites"
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

func DrawPlayer(player *sprites.Player, screen *ebiten.Image) {
	playerDrawOptions := ebiten.DrawImageOptions{}

	if !player.IsMovingRight {
		playerDrawOptions.GeoM.Scale(-1, 1)

		var frameWidth float64
		if player.IsIdle {
			frameWidth = player.IdleAnimation.MaxFrameWidth
		} else {
			frameWidth = player.ActionAnimations[player.CurrentAnimationIndex].MaxFrameWidth
		}
		playerDrawOptions.GeoM.Translate(frameWidth, 0)
	}

	playerDrawOptions.GeoM.Translate(player.X, player.Y)
	if player.IsIdle {
		currentFrame := player.IdleAnimation.Animate()
		screen.DrawImage(currentFrame, &playerDrawOptions)
		return
	}
	if player.ActionAnimations[player.CurrentAnimationIndex].AnimationComplete {
		player.IsAnimationLocked = false
		player.ActionAnimations[player.CurrentAnimationIndex].AnimationComplete = false
		player.CurrentAnimationIndex = 0
	}

	currentFrame, frameVector, canCancel := player.ActionAnimations[player.CurrentAnimationIndex].AnimateAction()

	if player.IsMovingRight {
		newVec := frameVector.Add(player.X, player.Y)
		player.X = newVec.DeltaX
		player.Y = newVec.DeltaY
	} else {
		newVec := frameVector.InvertX(player.X, player.Y)
		player.X = newVec.DeltaX
		player.Y = newVec.DeltaY
	}
	player.CanAnimationCancel = canCancel
	playerDrawOptions.GeoM.Translate(frameVector.DeltaX, frameVector.DeltaY)
	screen.DrawImage(currentFrame, &playerDrawOptions)
}
