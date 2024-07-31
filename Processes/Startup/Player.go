package startup

import sprites "GoPlat/Components/Sprites"

func CreateDefaultPlayer() (*sprites.Player) {
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
		HasSecondJump:     true,
		IsWallSliding:     false,
		IsAnimationLocked: false,
	}

	return player
}