package movement

import (
<<<<<<< HEAD
	"GoPlat/engine/collision"
	"GoPlat/engine/physics"
=======
	"GoPlat/Engine/collision"
>>>>>>> 16f3c53bd428513a1c986c0fe9c23443b6469d9c
	controls "GoPlat/gameComponents/controls"
	"GoPlat/gameComponents/levels"
	"GoPlat/gameComponents/sprites"

	"github.com/hajimehoshi/ebiten/v2"
)

<<<<<<< HEAD
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
=======
func shouldCancelPlayerInput(p *sprites.Player, lvl *levels.Level, directions []controls.Direction) bool {
	if p.IsAnimationLocked {
		AnimationLockWallOverride(p, lvl)
	}
	if p.IsAnimationLocked && !p.CanAnimationCancel {
		return true
>>>>>>> 16f3c53bd428513a1c986c0fe9c23443b6469d9c
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
<<<<<<< HEAD
	playerVector, specialAction := GetMovementVector(directions)

	var validMove bool
	if len(specialAction.Name) > 0 {
		validMove = HandleSpecialAction(p, specialAction.Name)
=======
	netInputForces, specialActionTriggered := GetNetForces_Input(directions)
	p.IsMovingRight = IsMovingRight(p, netInputForces)
	p.IsIdle = IsIdle(p, netInputForces, specialActionTriggered)

	if len(specialActionTriggered.Name) > 0 {
		validMove := HandleSpecialAction(p, specialActionTriggered.Name)
>>>>>>> 16f3c53bd428513a1c986c0fe9c23443b6469d9c
		if validMove {
			p.IsAnimationLocked = true
			p.IsIdle = false
		} else {
			//if they can animation cancel, the animation is not over
			//this prevents user from cancelling animation lock on accident
			if p.IsAnimationLocked {
				p.IsPhysicsLocked = true
			} else {
				p.IsPhysicsLocked = false
				physics.HandlePhysics(p, lvl, &rtnVector)
			}
<<<<<<< HEAD
=======
			return
>>>>>>> 16f3c53bd428513a1c986c0fe9c23443b6469d9c
		}
		return rtnVector
	} else {
		p.IsMovingRight = IsMovingRight(p, playerVector)
		p.IsIdle = IsIdle(p, playerVector, specialAction)
		p.IsPhysicsLocked = false
		p.CurrentAnimationIndex = 0
	}
<<<<<<< HEAD

	if !p.IsPhysicsLocked {
		physics.HandlePhysics(p, lvl, &playerVector)
	}

	rtnVector.DeltaX = playerVector.DeltaX
	rtnVector.DeltaY = playerVector.DeltaY

	return rtnVector
=======
	p.Physics.NetForce.X = netInputForces.X
	p.Physics.NetForce.Y = netInputForces.Y
>>>>>>> 16f3c53bd428513a1c986c0fe9c23443b6469d9c
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

<<<<<<< HEAD
func GetMovementVector(directions []controls.Direction) (controls.Vector, controls.Direction) {
=======
func GetNetForces_Input(directions []controls.Direction) (controls.Vector2, controls.Direction) {
>>>>>>> 16f3c53bd428513a1c986c0fe9c23443b6469d9c
	specialDirections := []controls.Direction{
		controls.JUMP,
		controls.DASHLEFT,
		controls.DASHRIGHT,
	}
<<<<<<< HEAD
	var vector controls.Vector
	var specialAction controls.Direction
	for _, direction := range directions {
		if math.Abs(direction.DeltaX) > math.Abs(vector.DeltaX) {
			vector.DeltaX = direction.DeltaX
		} else if math.Abs(direction.DeltaY) > math.Abs(vector.DeltaY) {
			vector.DeltaY = direction.DeltaY
		}
=======
	var netForce controls.Vector2
	var specialAction controls.Direction
	for _, direction := range directions {
		netForce.X += direction.ForceX
		netForce.Y += direction.ForceY
>>>>>>> 16f3c53bd428513a1c986c0fe9c23443b6469d9c
		for _, special := range specialDirections {
			if direction == special {
				specialAction = direction
			}
		}
	}

<<<<<<< HEAD
	return vector, specialAction
=======
	return netForce, specialAction
>>>>>>> 16f3c53bd428513a1c986c0fe9c23443b6469d9c
}

func IsMovingRight(player *sprites.Player, vector controls.Vector) bool {
	if !player.IsAnimationLocked {
		return vector.DeltaX >= 0
	}
	return player.IsMovingRight
}

<<<<<<< HEAD
func IsIdle(p *sprites.Player, vector controls.Vector, specialAction controls.Direction) bool {
	if len(specialAction.Name) > 0 || p.IsWallHanging {
		return false
	} else if vector.DeltaX != 0 || vector.DeltaY != 0 {
		return false
	} else if p.IsAirborn {
=======
func IsIdle(p *sprites.Player, velocity controls.Vector2, specialAction controls.Direction) bool {
	if len(specialAction.Name) > 0 || p.IsWallHanging {
		return false
	}
	if velocity.X != 0 || velocity.Y != 0 {
>>>>>>> 16f3c53bd428513a1c986c0fe9c23443b6469d9c
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
<<<<<<< HEAD
		p.IsPhysicsLocked = true
	} else {
		//if it isn't a wall or hurt or death animation, cancel that animation
		if (p.CurrentAnimationIndex < 1 || p.CurrentAnimationIndex > 8) ||
			(nearGround && p.CurrentAnimationIndex == 7) {
			p.ActionAnimations[p.CurrentAnimationIndex].CurrentFrameIndex = 0
			p.IsAnimationLocked = false
		}
=======
		p.IsGravityLocked = true
		return
	}
	//if it isn't a wall or hurt or death animation, cancel that animation
	if (p.CurrentAnimationIndex < 4 || p.CurrentAnimationIndex > 8) ||
		(nearGround && p.CurrentAnimationIndex == 7) {
		p.ActionAnimations[p.CurrentAnimationIndex].CurrentFrameIndex = 0
		p.IsAnimationLocked = false
>>>>>>> 16f3c53bd428513a1c986c0fe9c23443b6469d9c
	}
}
