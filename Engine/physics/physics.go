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
	stopPhysicsCalcs := playerOnGround(player, onGround)
	if stopPhysicsCalcs {return}

	nearWall, wallPlayerLeft := collision.DetectWall(player, lvl)
	
	if !nearWall {
		handleFall(player, pVector)
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

func playerOnGround(player *sprites.Player, onGround bool) bool {
	if player.IsAirborn {
		if onGround {
			player.IsAirborn = false
			return true
		}
	} else {
		if onGround { 
			return true
		}
		player.IsAirborn = true
	}
	return false
}

func handleFall(player *sprites.Player, pVector *controls.Vector) {
	player.CurrentAnimationIndex = 9
	player.IsIdle = false
	pVector.DeltaY += 1.1
}