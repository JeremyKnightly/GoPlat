package startup

import (
	score "GoPlat/gameComponents/PlayerStatus/Score"
	"GoPlat/gameComponents/levels"
	"GoPlat/gameComponents/sprites"

	"github.com/hajimehoshi/ebiten/v2"
)

func CreateDefaultPlayer(lvl *levels.Level) *sprites.Player {
	checkX, checkY := lvl.GetCheckpointXY(0)
	player := &sprites.Player{
		BioSprite: &sprites.BioSprite{
			Sprite: &sprites.Sprite{
				X: checkX,
				Y: checkY,
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
		Score:                  &score.Score{},
	}
	player.Score.ResetScore()

	return player
}
