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
	for _, layer := range level.Layers {
		for index, id := range layer.Data {
			x := index % layer.Width
			y := index / layer.Width

			x *= 16
			y *= 16

			srcX := (id - 1) % 15
			srcY := (id - 1) / 15

			srcX *= 16
			srcY *= 16

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
	currentFrame,frameVector,_ := player.ActionAnimations[player.CurrentAnimationIndex].Animate()
	if player.ActionAnimations[player.CurrentAnimationIndex].AnimationComplete {
		player.IsAnimationLocked = false
		player.ActionAnimations[player.CurrentAnimationIndex].AnimationComplete = false
	}
	if player.IsMovingRight {
		newVec := frameVector.Add(player.X, player.Y)
		player.X = newVec.DeltaX
		player.Y = newVec.DeltaY
	} else {
		newVec := frameVector.InvertX(player.X, player.Y)
		player.X = newVec.DeltaX
		player.Y = newVec.DeltaY
	}
	playerDrawOptions.GeoM.Translate(frameVector.DeltaX, frameVector.DeltaY)
	screen.DrawImage(currentFrame, &playerDrawOptions)

}