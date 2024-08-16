package movement

import (
	"GoPlat/engine/collision"
	"GoPlat/engine/physics"
	controls "GoPlat/gameComponents/controls"
	"GoPlat/gameComponents/levels"
	"GoPlat/gameComponents/sprites"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func HandleMovementCalculations(p *sprites.Player, playerControls []controls.Control, lvl *levels.Level) controls.Vector {
	rtnVector := controls.Vector{
		DeltaX: 0,
		DeltaY: 0,
	}

	if p.IsAnimationLocked {
		AnimationLockWallOverride(p, lvl)
	}
	if (p.IsAnimationLocked && !p.CanAnimationCancel) || p.IsDead {
		return rtnVector
	}
	directions := GetControlsPressed(playerControls)
	if p.IsAnimationLocked && !IsAnimationCancelling(p, directions) {
		return rtnVector
	}

	//handle movement
	playerVector, specialAction := GetMovementVector(directions)

	p.IsMovingRight = IsMovingRight(p, playerVector)
	p.IsIdle = IsIdle(p, playerVector, specialAction)
	//Idle Detection

	var validMove bool
	if len(specialAction.Name) > 0 {
		validMove = HandleSpecialAction(p, specialAction.Name)
		if validMove {
			p.IsAnimationLocked = true
		} else {
			//if they can animation cancel, the animation is not over
			//this prevents user from cancelling animation lock on accident
			if p.IsAnimationLocked {
				p.IsPhysicsLocked = true
			} else {
				p.IsPhysicsLocked = false
				physics.HandlePhysics(p, lvl, &rtnVector)
			}
			return rtnVector
		}
	} else {
		p.IsPhysicsLocked = false
		p.CurrentAnimationIndex = 0
	}

	if !p.IsPhysicsLocked {
		physics.HandlePhysics(p, lvl, &playerVector)
	}

	rtnVector.DeltaX = playerVector.DeltaX
	rtnVector.DeltaY = playerVector.DeltaY

	return rtnVector
}

func GetControlsPressed(controlSlice []controls.Control) []controls.Direction {
	var directions []controls.Direction
	for _, control := range controlSlice {
		controlActivated := true
		if len(control.Keys) > 0 {
			for _, key := range control.Keys {
				if !ebiten.IsKeyPressed(key) {
					controlActivated = false
				}
			}
		} else {
			controlActivated = ebiten.IsKeyPressed(control.Key)
		}
		if controlActivated {
			directions = append(directions, control.Direction)
		}
	}

	return directions
}

func GetMovementVector(directions []controls.Direction) (controls.Vector, controls.Direction) {
	specialDirections := []controls.Direction{
		controls.JUMP,
		controls.DASHLEFT,
		controls.DASHRIGHT,
	}
	var vector controls.Vector
	var specialAction controls.Direction
	for _, direction := range directions {
		if math.Abs(direction.DeltaX) > math.Abs(vector.DeltaX) {
			vector.DeltaX = direction.DeltaX
		} else if math.Abs(direction.DeltaY) > math.Abs(vector.DeltaY) {
			vector.DeltaY = direction.DeltaY
		}
		for _, special := range specialDirections {
			if direction == special {
				specialAction = direction
			}
		}
	}

	return vector, specialAction
}

func IsMovingRight(player *sprites.Player, vector controls.Vector) bool {
	if !player.IsAnimationLocked {
		return vector.DeltaX >= 0
	}
	return player.IsMovingRight
}

func IsIdle(p *sprites.Player, vector controls.Vector, specialAction controls.Direction) bool {
	if len(specialAction.Name) > 0 || p.IsWallHanging {
		return false
	} else if vector.DeltaX != 0 || vector.DeltaY != 0 {
		return false
	}

	return true
}

func IsAnimationCancelling(p *sprites.Player, input []controls.Direction) bool {
	for _, direction := range input {
		for _, cancelDirection := range p.ActionAnimations[p.CurrentAnimationIndex].AllowCancelOnDirections {
			if direction == cancelDirection {
				return true
			}
		}
	}
	return false
}

func AnimationLockWallOverride(p *sprites.Player, lvl *levels.Level) {
	nearWall, _, _ := collision.DetectWall(p, lvl)
	nearGround := collision.DetectGround(p, lvl)
	if !nearWall {
		p.IsPhysicsLocked = true
	} else {
		//if it isn't a wall or hurt or death animation, cancel that animation
		if (p.CurrentAnimationIndex < 1 || p.CurrentAnimationIndex > 8) ||
			(nearGround && p.CurrentAnimationIndex == 7) {
			p.ActionAnimations[p.CurrentAnimationIndex].CurrentFrameIndex = 0
			p.IsAnimationLocked = false
		}
	}
}
