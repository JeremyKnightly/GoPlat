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
	
	stopPhysicsCalcs := playerOnGround(player, onGround)
	if stopPhysicsCalcs {return}
	pVector.DeltaY += .9
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

func playerOnGround(player *sprites.Player, onGround bool) bool {
	if player.IsAirborn {
		if onGround {
			player.IsAirborn = false
			return true
			//if player is not airborn, no actions will be overridden
		}
	} else {
		if onGround { 
			return true
		}
		player.IsAirborn = true
	}
	return false
}