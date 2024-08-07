package movement

import (
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
		p.IsPhysicsLocked = true
	}
	if p.IsAnimationLocked && !p.CanAnimationCancel {
		return rtnVector
	}
	directions := GetControlsPressed(playerControls)
	if p.IsAnimationLocked && !IsAnimationCancelling(p, directions) {
		return rtnVector
	}

	//handle movement
	playerVector, specialAction := GetMovementVector(directions)

	p.IsMovingRight = IsMovingRight(playerVector)
	p.IsIdle = IsIdle(playerVector, specialAction)
	//Idle Detection

	if len(specialAction.Name) > 0 {
		validMove := HandleSpecialAction(p, specialAction.Name)
		if !validMove {return rtnVector}
		p.IsAnimationLocked = true
	} else {
		p.CurrentAnimationIndex = 0
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
		} else if ebiten.IsKeyPressed(control.Key) {
			controlActivated = true
		} else {
			controlActivated = false
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
	for _,direction := range directions { 
		if math.Abs(direction.DeltaX) > math.Abs(vector.DeltaX) {
			vector.DeltaX = direction.DeltaX
		} else if math.Abs(direction.DeltaY) > math.Abs(vector.DeltaY) {
			vector.DeltaY = direction.DeltaY
		}
		for _,special := range specialDirections {
			if direction == special {
				specialAction = direction
			}
		}
	}

	return vector, specialAction
}

func IsMovingRight(vector controls.Vector) bool {
	return vector.DeltaX >= 0
}

func IsIdle(vector controls.Vector,specialAction controls.Direction) bool {
	if len(specialAction.Name) > 0 {
		return false
	} else if vector.DeltaX != 0 || vector.DeltaY != 0 {
		return false
	}

	return true
}

func IsAnimationCancelling (p *sprites.Player, input []controls.Direction) (bool) {
	for _,direction := range input {
		for _,cancelDirection := range 
		p.ActionAnimations[p.CurrentAnimationIndex].AllowCancelOnDirections {
			if direction == cancelDirection {
				return true
			}
		}
	}
	return false
}
