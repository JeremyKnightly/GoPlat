package movement

import (
	"GoPlat/engine/collision"
	"GoPlat/engine/physics"
	controls "GoPlat/gameComponents/controls"
	"GoPlat/gameComponents/gamepad"
	"GoPlat/gameComponents/levels"
	"GoPlat/gameComponents/sprites"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func HandleMovementCalculations(gp *gamepad.Gamepad, p *sprites.Player, playerControls []controls.Control, lvl *levels.Level) controls.Vector {
	if p.IsDead {
		return controls.GetBlankVector()
	}
	if p.IsAnimationLocked {
		AnimationLockWallOverride(p, lvl)
	}

	directions := GetControlsPressed(gp, playerControls)
	isLocked := p.IsAnimationLocked && (!p.CanAnimationCancel || !IsAnimationCancelling(p, directions))
	xMulti := GetXMulti(p, directions, isLocked)
	playerVector, specialAction := GetMovementVector(directions, xMulti)

	if len(specialAction.Name) > 0 && !isLocked && DoSpecialAction(p, specialAction.Name) {
		p.IsAnimationLocked = true
	} else if !p.IsAnimationLocked && !p.IsAirborn {
		p.CurrentAnimationIndex = 0
	}
	p.IsIdle = IsIdle(p, playerVector, specialAction)
	p.IsMovingRight = IsMovingRight(p, playerVector)
	p.IsPhysicsLocked = p.IsAnimationLocked

	if !p.IsPhysicsLocked {
		physics.HandlePhysics(p, lvl, &playerVector)
	}

	return playerVector
}

func GetXMulti(p *sprites.Player, directions []controls.Direction, isLocked bool) float64 {
	movementMulti := 1.0
	animIDX := p.CurrentAnimationIndex
	if animIDX == 1 || animIDX == 6 || animIDX == 4 {
		movementMulti = 0
	} else if isLocked {
		movementMulti = .35
	}

	return movementMulti
}

func GetControlsPressed(gp *gamepad.Gamepad, controlSlice []controls.Control) []controls.Direction {
	var directions []controls.Direction
	for _, control := range controlSlice {
		controlActivated := true
		if control.GetType() == "Gamepad" {
			if control.InputType == "AXIS" {
				controlActivated = gp.AxisMatchesInput(control)
			} else if control.InputType == "BTN" {
				controlActivated = gp.IsButtonPressed(control.Input)
			}
		} else if control.GetType() == "Keyboard" {
			if len(control.Keys) > 0 {
				for _, key := range control.Keys {
					if !ebiten.IsKeyPressed(key) {
						controlActivated = false
					}
				}
			} else {
				controlActivated = ebiten.IsKeyPressed(control.Key)
			}
		}
		if controlActivated {
			directions = append(directions, control.Direction)
		}

	}

	return directions
}

func GetMovementVector(directions []controls.Direction, XMulti float64) (controls.Vector, controls.Direction) {
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

	vector.ScaleX(XMulti)

	return vector, specialAction
}

func IsMovingRight(player *sprites.Player, vector controls.Vector) bool {
	if player.IsAnimationLocked {
		return player.IsMovingRight
	}
	if !player.IsWallSliding {
		return vector.DeltaX > 0 || (player.IsMovingRight && vector.DeltaX == 0)
	} else {
		return vector.DeltaX < 0 || (!player.IsMovingRight && vector.DeltaX == 0)
	}
}

func IsIdle(p *sprites.Player, vector controls.Vector, specialAction controls.Direction) bool {
	if len(specialAction.Name) > 0 || p.IsWallHanging {
		return false
	} else if vector.DeltaX != 0 || vector.DeltaY != 0 {
		return false
	} else if p.IsAirborn {
		return false
	} else if p.IsAnimationLocked || p.IsPhysicsLocked {
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
