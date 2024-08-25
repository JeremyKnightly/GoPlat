package sprites

import (
	"GoPlat/gameComponents/animations"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	Image *ebiten.Image
	Frame *Frame
	X, Y  float64
}

type Frame struct {
	ImageToDraw        *ebiten.Image
	EffectImageToDraw  *ebiten.Image
	HasEffect          bool
	ImageOptions       ebiten.DrawImageOptions
	EffectOptions      ebiten.DrawImageOptions
	EffectOffset       float64
	EffectOffsetOneWay bool
	EffectOffsetRight  bool
}

type BioSprite struct {
	*Sprite
	ActionAnimations      []*animations.ActionAnimation
	IdleAnimation         *animations.Animation
	IsIdle                bool
	IsMovingRight         bool
	CurrentAnimationIndex uint16
	IsDead                bool
}
