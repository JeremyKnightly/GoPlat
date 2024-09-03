package runtime

import (
	"GoPlat/Engine/camera"
	sprites "GoPlat/gameComponents/sprites"

	"github.com/hajimehoshi/ebiten/v2"
)

func DrawSprites(npsm *sprites.NPSpriteManager, screen *ebiten.Image, cam *camera.Camera) {
	for _, interactable := range npsm.ExistingInteractables {
		spriteOpts := ebiten.DrawImageOptions{}
		x, y := interactable.GetPosition()
		spriteOpts.GeoM.Translate(x, y)
		spriteOpts.GeoM.Translate(cam.X, cam.Y)
		spriteImg := interactable.GetSpriteFrameImage()
		if spriteImg != nil {
			screen.DrawImage(spriteImg, &spriteOpts)
		} else {
			println("SpriteFrameNotLoaded: ", interactable.GetCurrentSpriteName())
		}
	}
}
