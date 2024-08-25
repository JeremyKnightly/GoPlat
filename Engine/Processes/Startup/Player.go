package startup

import (
	"GoPlat/gameComponents/levels"
	"GoPlat/gameComponents/sprites"

	"github.com/hajimehoshi/ebiten/v2"
)

func CreateDefaultPlayer(lvl *levels.Level) *sprites.Player {
	checkX, checkY := lvl.GetCheckpointXY(0)
	player := &sprites.Player{
		BioSprite: &sprites.BioSprite{
			Sprite: &sprites.Sprite{
<<<<<<< HEAD
				X: checkX,
				Y: checkY,
=======
				Physics: &controls.PhysicsObj{
					Position: controls.Vector2{
						X: 50.0,
						Y: 200.0,
					},
					Velocity: controls.Vector2{
						X: 0,
						Y: 0,
					},
					Acceleration: controls.Vector2{
						X: 0,
						Y: 0,
					},
					Mass: .012,
				},
>>>>>>> 16f3c53bd428513a1c986c0fe9c23443b6469d9c
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
		CanJump:                true,
		HasSecondJump:          true,
		CanDash:                true,
		IsWallSliding:          false,
		IsAnimationLocked:      false,
		IsPhysicsLocked:        false,
		CanAnimationCancel:     false,
		IsAirborn:              false,
		CurrentCheckpointIndex: 0,
	}

	return player
}
