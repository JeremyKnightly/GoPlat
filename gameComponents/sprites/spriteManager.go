package sprites

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type NPSpriteManager struct {
	ExistingInteractables []interactableSprite
	SpriteDB              []NamedSprite
}

type NamedSprite struct {
	Sprite *ebiten.Image
	Name   string
}

func NewNPSpriteManager() *NPSpriteManager {
	return &NPSpriteManager{}
}

func (npsm *NPSpriteManager) ClearInteractables() {
	npsm.ExistingInteractables = []interactableSprite{}
}

func (npsm *NPSpriteManager) GetImageFromNamedSprite(imageName string) *ebiten.Image {
	for _, namedSprite := range npsm.SpriteDB {
		if namedSprite.Name == imageName {
			return namedSprite.Sprite
		}
	}
	return nil
}
