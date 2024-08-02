package sprites

import (
	"GoPlat/Components/animations"
	"time"

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
	DashCooldowntime  time.Duration
	DashLastUsed      time.Time
	JumpCooldownTime  time.Duration
	JumpLastUsed      time.Time
	IsAnimationLocked bool
	HasSecondJump     bool
	IsWallSliding     bool
}
