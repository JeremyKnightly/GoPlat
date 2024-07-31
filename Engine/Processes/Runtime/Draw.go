package runtime

import (
	levels "GoPlat/Components/Levels"
	sprites "GoPlat/Components/Sprites"
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

	if player.CurrentAnimationIndex == 1 {
		playerDrawOptions.GeoM.Translate(0, -16)
	}
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
	var currentFrame *ebiten.Image
	if player.IsIdle {
		currentFrame = player.IdleAnimation.Animate()
	} else {
		currentFrame = player.ActionAnimations[player.CurrentAnimationIndex].Animate()
	}
	screen.DrawImage(currentFrame, &playerDrawOptions)
}