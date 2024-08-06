package startup

import (
	"GoPlat/gameComponents/animations"
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
	Fall := animations.GeneratePlayerFallNoWeapon()
	Land := animations.GeneratePlayerLandNoWeapon()

	animations := []*animations.ActionAnimation{
		Walk,	//0
		Run,	//1
		Jump,	//2
		DblJump,	//3
		EdgeClimb,	//4
		Hurt,	//5	
		WallGrab,	//6
		WallSlide,	//7
		Death,	//8
		Fall,	//9
		Land,	//10
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
