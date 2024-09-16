package physics

import (
	"GoPlat/engine/collision"
	controls "GoPlat/gameComponents/controls"
	"GoPlat/gameComponents/levels"
	"GoPlat/gameComponents/sprites"
	"math"
)

func HandlePhysics(player *sprites.Player, lvl *levels.Level, pVector *controls.Vector) {
	onGround := collision.DetectGround(player, lvl)
	stopPhysics := shouldStopPhysics(player, onGround)
	if stopPhysics {
		player.CanDash = true
		player.CanJump = true
		player.IsWallSliding = false
		return
	}
	player.IsIdle = false

	nearWall, wallPlayerLeft, canSlide := collision.DetectWall(player, lvl)

	if nearWall {
		player.CanJump = true
		eventHandled := handleWallLogic(player, lvl, wallPlayerLeft, canSlide)
		if eventHandled {
			return
		} else {
		}
	}

	handleFall(player, pVector)
}

func handleWallLogic(player *sprites.Player, lvl *levels.Level, wallPlayerLeft, canWallSlide bool) bool {
	if canWallHang(player, lvl, wallPlayerLeft) {
		player.IsMovingRight = !wallPlayerLeft
		player.CurrentAnimationIndex = 6
		player.IsWallHanging = true
		player.IsAnimationLocked = true
		player.IsWallSliding = false
	} else if canWallSlide {
		player.CurrentAnimationIndex = 7
		player.IsMovingRight = wallPlayerLeft
		player.IsWallSliding = true
	} else {
		player.IsWallSliding = false
		return false
	}
	return true
}

func canWallHang(player *sprites.Player, lvl *levels.Level, wallPlayerLeft bool) bool {
	playerRect := collision.GetPlayerRect(player)

	//takes a rectangle above players head and in
	//specified direction to see if there is a ledge
	playerRect.Y -= (playerRect.Height)
	if wallPlayerLeft {
		playerRect.X -= playerRect.Width / 2
	} else {
		playerRect.X += playerRect.Width / 2
	}
	yCoord, xCoord := collision.GetWallHangCoords(lvl, playerRect, wallPlayerLeft)
	if yCoord == 0 {
		return false
	} else {
		//constraint to prevent shenanigans
		if math.Abs(player.Y-yCoord) <= 14 {
			player.Y = yCoord - 4
			player.X = xCoord - 8
		}
	}

	//nearGround := collision.DetectGroundRect(playerRect, lvl)

	return true
}

func shouldStopPhysics(player *sprites.Player, onGround bool) bool {
	stopPhysics := false
	// isAirborn checks if player WAS airborn before this frame
	// onGround checks current state
	if onGround {
		player.IsAirborn = false
		stopPhysics = true
	} else {
		player.IsAirborn = true
	}

	return stopPhysics
}

func handleFall(player *sprites.Player, pVector *controls.Vector) {
	player.CurrentAnimationIndex = 9
	player.IsWallSliding = false
	player.IsIdle = false
	pVector.DeltaY += 1.6
}
