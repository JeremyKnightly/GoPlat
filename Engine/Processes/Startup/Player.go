package startup

import (
	"GoPlat/components/sprites"
	"time"
)

func CreateDefaultPlayer() *sprites.Player {
	player := &sprites.Player{
		BioSprite: &sprites.BioSprite{
			Sprite: &sprites.Sprite{
				Image: nil,
				X:     50,
				Y:     150,
			},
			ActionAnimations:      GetPlayerActionAnimations(),
			IsMovingRight:         true,
			IdleAnimation:         GetPlayerIdleAnimation(),
			IsIdle:                true,
			CurrentAnimationIndex: 0,
		},
		DashCooldowntime:  time.Millisecond * 3000,
		DashLastUsed:      time.Date(1994,time.April,4,2,0,0,0,time.Now().Local().Location()),
		JumpCooldownTime:  time.Millisecond * 1800,
		JumpLastUsed:      time.Date(1994,time.April,4,2,0,0,0,time.Now().Local().Location()),
		HasSecondJump:     true,
		IsWallSliding:     false,
		IsAnimationLocked: false,
		IsPhysicsLocked: false,
		CanAnimationCancel: false,
	}

	return player
}
