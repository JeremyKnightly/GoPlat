package physics

import (
	"GoPlat/engine/collision"
	controls "GoPlat/gameComponents/controls"
	"GoPlat/gameComponents/levels"
	"GoPlat/gameComponents/sprites"
	"time"
)

//need position data.
//Need to find out if player is on the ground.
//Also need to find out if player is against a wall
//wall sliding is airborn + on wall


func HandlePhysics(player *sprites.Player, lvl *levels.Level, pVector *controls.Vector) {
	onGround := collision.DetectGround(player, lvl)
	nearWall, wallPlayerLeft := collision.DetectWall(player, lvl)
	
	if player.IsAirborn {
		if onGround {
			player.IsAirborn = false
			return
			//if player is not airborn, no actions will be overridden
		}
	} else {
		if onGround { 
			return
		}
		player.IsAirborn = true
		player.TimeWentAirborn = time.Now()
	}
	println("things can happen")
	
	if nearWall{
		println("There is a wall to my")
		if wallPlayerLeft {
			print(" left")
		} else {
			print(" right")
		}
	}
}