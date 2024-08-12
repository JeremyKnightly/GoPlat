package startup

import (
	controls "GoPlat/gameComponents/controls"
	"GoPlat/gameComponents/sprites"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func CreateDefaultPlayer() *sprites.Player {
	player := &sprites.Player{
		BioSprite: &sprites.BioSprite{
			Sprite: &sprites.Sprite{
				Physics: &controls.PhysicsObj{
					Position: controls.Vector2{
						X: 50.0,
						Y: 100.0,
					},
					Mass: 1,
				},

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
		IsGravityLocked:    false,
		CanAnimationCancel: false,
		IsAirborn:          false,
	}

	return player
}
