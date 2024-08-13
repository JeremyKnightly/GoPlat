package sprites

import (
	"GoPlat/gameComponents/animations"
	"GoPlat/gameComponents/controls"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	Image   *ebiten.Image
	Frame   *Frame
	Physics *controls.PhysicsObj
}

type Frame struct {
	ImageToDraw       *ebiten.Image
	EffectImageToDraw *ebiten.Image
	HasEffect         bool
	ImageOptions      ebiten.DrawImageOptions
	EffectOptions     ebiten.DrawImageOptions
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
	IsGravityLocked    bool
	CanAnimationCancel bool
	IsAnimationLocked  bool
	IsAirborn          bool
	HasSecondJump      bool
	IsWallSliding      bool
	IsWallHanging      bool
}
