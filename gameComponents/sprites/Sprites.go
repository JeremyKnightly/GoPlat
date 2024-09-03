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
	GetPosition() (float64, float64)
	GetMessageText() string
	SetMessageText(string)
	SetNextSpriteName(string)
	GoToNextSprite()
	GetSpawnType() string
	SetSpawnType(string)
	GetCurrentSpriteName() string
	SetSpriteFrameImage(*ebiten.Image)
	GetSpriteFrameImage() *ebiten.Image
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

func (s *Sprite) GetPosition() (float64, float64) {
	return s.X, s.Y
}

func (s *Sprite) SetPosition(x, y float64) {
	s.X = x
	s.Y = y
}

func GetNewFrame() *Frame {
	return &Frame{
		ImageToDraw: nil,
	}
}
