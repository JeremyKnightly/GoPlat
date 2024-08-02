package startup

import (
	"GoPlat/Components/animations"
	"log"
)

func GetPlayerActionAnimations() []*animations.ActionAnimation {
	playerWalk := animations.GeneratePlayerWalkNoWeapon()
	playerJump := animations.GeneratePlayerJumpNoWeapon()

	animations := []*animations.ActionAnimation{
		playerWalk,
		playerJump,
	}

	if len(animations) == 0 {
		log.Fatal("FATAL ERROR: COULD NOT GENERATE PLAYER ACTIONS!")
	}

	return animations
}

func GetPlayerIdleAnimation() *animations.Animation {
	playerIdle := animations.GeneratePlayerIdleNoWeapon()

	return playerIdle
}