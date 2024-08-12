package sprites

import (
	"GoPlat/gameComponents/animations"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	Image *ebiten.Image
	Frame *Frame
	X, Y  float64
}

type Frame struct {
	ImageToDraw       *ebiten.Image
	EffectImageToDraw *ebiten.Image
	HasEffect         bool
	ImageOptions      *ebiten.DrawImageOptions
	EffectOptions     *ebiten.DrawImageOptions
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
	DashCooldowntime   time.Duration
	DashLastUsed       time.Time
	CanJump            bool
	IsPhysicsLocked    bool
	CanAnimationCancel bool
	IsAnimationLocked  bool
	IsAirborn          bool
	HasSecondJump      bool
	IsWallSliding      bool
	IsWallHanging      bool
}
