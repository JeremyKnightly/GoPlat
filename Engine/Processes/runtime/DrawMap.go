package runtime

import (
	"GoPlat/Engine/camera"
	levels "GoPlat/gameComponents/levels"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

const tileSize = 16

func DrawLevel_FirstDraw(level *levels.Level, screen *ebiten.Image, cam *camera.Camera) {
	for idx, layer := range level.Layers {
		if !layer.FirstDraw {
			continue
		}
		DrawLayer(level, idx, screen, cam)
	}
}

func DrawLevel_SecondDraw(level *levels.Level, screen *ebiten.Image, cam *camera.Camera) {
	for idx, layer := range level.Layers {
		if layer.FirstDraw {
			continue
		}
		DrawLayer(level, idx, screen, cam)
	}
}

func DrawLayer(level *levels.Level, layerIdx int, screen *ebiten.Image, cam *camera.Camera) {
	mapDrawOptions := ebiten.DrawImageOptions{}
	layer := level.Layers[layerIdx]
	for index, id := range layer.Data {
		if id == 0 {
			continue
		}
		tilesInRow := level.TilemapImage.Bounds().Dx() / tileSize

		x := index % int(layer.Width)
		y := index / int(layer.Width)

		x *= tileSize
		y *= tileSize

		srcX := (id - 1) % tilesInRow
		srcY := (id - 1) / tilesInRow

		srcX *= tileSize
		srcY *= tileSize

		mapDrawOptions.GeoM.Translate(float64(x), float64(y))
		mapDrawOptions.GeoM.Translate(cam.X, cam.Y)

		//tilesInRow := spriteSheet.Bounds().Dx() / layer.TileSize
		screen.DrawImage(level.TilemapImage.SubImage(image.Rect(srcX, srcY, srcX+tileSize, srcY+tileSize)).(*ebiten.Image),
			&mapDrawOptions,
		)

		mapDrawOptions.GeoM.Reset()
	}
}
