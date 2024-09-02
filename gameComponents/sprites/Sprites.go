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

type interactableSprite interface {
	SetPosition(float64, float64)
	GetMessageText() string
	SetMessageText(string)
	SetNextSpriteName(string)
	GoToNextSprite()
	GetSpawnType() string
	SetSpawnType(string)
	GetCurrentSpriteName() string
	SetSpriteFrameImage(*ebiten.Image)
	AddToPlayerStatus(*Player)
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
