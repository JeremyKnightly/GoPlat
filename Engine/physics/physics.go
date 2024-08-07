package physics

import (
	"GoPlat/engine/collision"
	controls "GoPlat/gameComponents/controls"
	"GoPlat/gameComponents/levels"
	"GoPlat/gameComponents/sprites"
)

//need position data.
//Need to find out if player is on the ground.
//Also need to find out if player is against a wall
//wall sliding is airborn + on wall


func HandlePhysics(player *sprites.Player, lvl *levels.Level, pVector *controls.Vector) {
	onGround := collision.DetectGround(player, lvl)
	nearWall, wallPlayerLeft := collision.DetectWall(player, lvl)
	if !onGround {

		//player.IsAirborn = true
		//player.TimeWentAirborn = time.Now()
	} else {
		//player.IsAirborn = false
	}
	if nearWall{
		println("There is a wall to my")
		if wallPlayerLeft {
			print(" left")
		} else {
			print(" right")
		}
	}
}