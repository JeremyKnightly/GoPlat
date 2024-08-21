package physics

import (
	"GoPlat/engine/collision"
	controls "GoPlat/gameComponents/controls"
	"GoPlat/gameComponents/levels"
	"GoPlat/gameComponents/sprites"
)

func HandlePhysics(player *sprites.Player, lvl *levels.Level, pVector *controls.Vector) {
	onGround := collision.DetectGround(player, lvl)
	stopPhysicsCalcs := adjustedAirbornStatus(player, onGround)
	if stopPhysicsCalcs {
		player.CanDash = true
		player.CanJump = true
		return
	}
	player.IsIdle = false

	nearWall, wallPlayerLeft, canSlide := collision.DetectWall(player, lvl)

	if nearWall {
		player.CanJump = true
		eventHandled := handleWallLogic(player, lvl, wallPlayerLeft, canSlide)
		if eventHandled {
			return
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
	} else if canWallSlide {
		player.CurrentAnimationIndex = 7
		player.IsMovingRight = wallPlayerLeft
		player.IsWallSliding = true
	} else {
		return false
	}
	return true
}

func canWallHang(player *sprites.Player, lvl *levels.Level, wallPlayerLeft bool) bool {
	playerRect := collision.GetPlayerRect(player)

	//takes a rectangle above players head and in
	//specified direction to see if there is a ledge
	playerRect.Y -= (4 + playerRect.Height)
	if wallPlayerLeft {
		playerRect.X -= playerRect.Width / 2
	} else {
		playerRect.X += playerRect.Width / 2
	}

	validMove := collision.IsValidMoveRect(lvl, playerRect)
	nearGround := collision.DetectGroundRect(playerRect, lvl)

	return validMove && nearGround
}

func adjustedAirbornStatus(player *sprites.Player, onGround bool) bool {
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
	pVector.DeltaY += 1.4
}
