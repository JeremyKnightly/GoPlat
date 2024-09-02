package sprites

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type NPSpriteManager struct {
	ExistingInteractables []*interactableSprite
	SpriteDB              []NamedSprite
}

type NamedSprite struct {
	sprite *ebiten.Image
	name   string
}

func NewNPSpriteManager() *NPSpriteManager {
	return &NPSpriteManager{}
}
