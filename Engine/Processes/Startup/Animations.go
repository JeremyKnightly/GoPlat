package startup

import (
	"GoPlat/Components/animations"
	"log"
)

func GetPlayerActionAnimations() []*animations.ActionAnimation {
	Walk := animations.GeneratePlayerWalkNoWeapon()
	Jump := animations.GeneratePlayerJumpNoWeapon()
	Run := animations.GeneratePlayerDashNoWeapon()
	DblJump := animations.GeneratePlayerDoubleJumpNoWeapon()
	EdgeClimb := animations.GeneratePlayerEdgeClimbNoWeapon()
	Hurt := animations.GeneratePlayerHurtNoWeapon()
	WallGrab := animations.GeneratePlayerWallGrabNoWeapon()
	WallSlide := animations.GeneratePlayerWallSlideNoWeapon()
	Death := animations.GeneratePlayerDeathNoWeapon()

	animations := []*animations.ActionAnimation{
		Walk,	
		Run,
		Jump,
		DblJump,
		EdgeClimb,
		Hurt,
		WallGrab,
		WallSlide,
		Death,
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
