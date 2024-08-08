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
	if stopPhysicsCalcs {
		player.CanJump = true
		return
	}
	player.IsIdle = false

	nearWall, wallPlayerLeft := collision.DetectWall(player, lvl)
	
	if !nearWall {
		handleFall(player, pVector)
		return
	}
	handleWallLogic(player, lvl, wallPlayerLeft, pVector)
}

func handleWallLogic(player *sprites.Player, lvl *levels.Level, wallPlayerLeft bool, pVector *controls.Vector) {
	if canWallHang(player, lvl, wallPlayerLeft) {
		player.IsMovingRight = !wallPlayerLeft
		player.CurrentAnimationIndex = 6
		player.IsWallHanging = true
		player.IsAnimationLocked = true
	} else {
		player.CurrentAnimationIndex = 7
		player.IsMovingRight = wallPlayerLeft
		player.IsWallSliding = true
	}
}


func canWallHang(player *sprites.Player, lvl *levels.Level, wallPlayerLeft bool) bool {
	playerRect := collision.GetPlayerRect(player)
	
	//takes a rectangle above players head and in 
	//specified direction to see if there is a ledge
	playerRect.Y -= (4 + playerRect.Height)
	if wallPlayerLeft {
		playerRect.X -= playerRect.Width/2
	} else {
		playerRect.X += playerRect.Width/2
	}

	validMove := collision.IsValidMoveRect(lvl, playerRect)
	
	return validMove
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
	pVector.DeltaY += 1
}