package startup

import (
	"GoPlat/gameComponents/sprites"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func CreateDefaultPlayer() *sprites.Player {
	player := &sprites.Player{
		BioSprite: &sprites.BioSprite{
			Sprite: &sprites.Sprite{
				X: 50,
				Y: 100,
				Frame: &sprites.Frame{
					ImageOptions:  ebiten.DrawImageOptions{},
					EffectOptions: ebiten.DrawImageOptions{},
					HasEffect:     false,
				},
			},
			ActionAnimations:      GetPlayerActionAnimations(),
			IsMovingRight:         true,
			IdleAnimation:         GetPlayerIdleAnimation(),
			IsIdle:                true,
			CurrentAnimationIndex: 0,
		},
		DashCooldowntime:   time.Millisecond * 3000,
		DashLastUsed:       time.Date(1994, time.April, 4, 2, 0, 0, 0, time.Now().Local().Location()),
		CanJump:            true,
		HasSecondJump:      true,
		IsWallSliding:      false,
		IsAnimationLocked:  false,
		IsPhysicsLocked:    false,
		CanAnimationCancel: false,
		IsAirborn:          false,
	}

	return player
}
