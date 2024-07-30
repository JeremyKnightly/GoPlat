package sprites

import (
	"GoPlat/Components/animations"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	Image *ebiten.Image
	X, Y  float64
}

type BioSprite struct {
	*Sprite
	ActionAnimations      []*animations.ActionAnimation
	IdleAnimation         *animations.Animation
	IsIdle                bool
	IsMovingRight         bool
	CurrentAnimationIndex uint16
}

type Player struct {
	*BioSprite
	HasSecondJump bool
	IsWallSliding bool
}
