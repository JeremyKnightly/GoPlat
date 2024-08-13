package movement

import (
	"GoPlat/Engine/collision"
	controls "GoPlat/gameComponents/controls"
	"GoPlat/gameComponents/levels"
	"GoPlat/gameComponents/sprites"

	"github.com/hajimehoshi/ebiten/v2"
)

func shouldCancelPlayerInput(p *sprites.Player, lvl *levels.Level, directions []controls.Direction) bool {
	if p.IsAnimationLocked {
		AnimationLockWallOverride(p, lvl)
	}
	if p.IsAnimationLocked && !p.CanAnimationCancel {
		return true
	}
	if p.IsAnimationLocked && !IsAnimationCancelling(p, directions) {
		return true
	}
	return false
}
func HandleMovementCalculations(p *sprites.Player, playerControls []controls.Control, lvl *levels.Level) {
	directions := GetControlsPressed(playerControls)
	cancel := shouldCancelPlayerInput(p, lvl, directions)
	if cancel {
		return
	}

	//handle movement
	netInputForces, specialActionTriggered := GetNetForces_Input(directions)
	p.IsMovingRight = IsMovingRight(p, netInputForces)
	p.IsIdle = IsIdle(p, netInputForces, specialActionTriggered)

	if len(specialActionTriggered.Name) > 0 {
		validMove := HandleSpecialAction(p, specialActionTriggered.Name)
		if validMove {
			p.IsAnimationLocked = true
		} else {
			//if they can animation cancel, the animation is not over
			//this prevents user from cancelling animation lock on accident
			if p.IsAnimationLocked {
				p.IsGravityLocked = true
			} else {
				p.IsGravityLocked = false
				//physics.HandlePhysics(p, lvl, &rtnVector)
			}
			return
		}
	} else {
		p.IsGravityLocked = false
		p.CurrentAnimationIndex = 0
	}
	p.Physics.NetForce.X = netInputForces.X
	p.Physics.NetForce.Y = netInputForces.Y
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

func GetNetForces_Input(directions []controls.Direction) (controls.Vector2, controls.Direction) {
	specialDirections := []controls.Direction{
		controls.JUMP,
		controls.DASHLEFT,
		controls.DASHRIGHT,
	}
	var netForce controls.Vector2
	var specialAction controls.Direction
	for _, direction := range directions {
		netForce.X += direction.ForceX
		netForce.Y += direction.ForceY
		for _, special := range specialDirections {
			if direction == special {
				specialAction = direction
			}
		}
	}

	return netForce, specialAction
}

func IsMovingRight(player *sprites.Player, vector controls.Vector2) bool {
	if !player.IsAnimationLocked {
		return vector.X >= 0
	}
	if player.IsMovingRight {
		return true
	}
	return false
}

func IsIdle(p *sprites.Player, velocity controls.Vector2, specialAction controls.Direction) bool {
	if len(specialAction.Name) > 0 || p.IsWallHanging {
		return false
	}
	if velocity.X != 0 || velocity.Y != 0 {
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
	nearWall, _ := collision.DetectWall(p, lvl)
	nearGround := collision.DetectGround(p, lvl)
	if !nearWall {
		p.IsGravityLocked = true
		return
	}
	//if it isn't a wall or hurt or death animation, cancel that animation
	if (p.CurrentAnimationIndex < 4 || p.CurrentAnimationIndex > 8) ||
		(nearGround && p.CurrentAnimationIndex == 7) {
		p.ActionAnimations[p.CurrentAnimationIndex].CurrentFrameIndex = 0
		p.IsAnimationLocked = false
	}
}
