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
		DashCooldowntime:  time.Millisecond * 4000,
		DashLastUsed:      time.Now(),
		JumpCooldownTime:  time.Millisecond * 2500,
		JumpLastUsed:      time.Now(),
		HasSecondJump:     true,
		IsWallSliding:     false,
		IsAnimationLocked: false,
	}

	return player
}
