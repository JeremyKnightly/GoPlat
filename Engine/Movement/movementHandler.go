package movement

import (
	"GoPlat/engine/collision"
	controls "GoPlat/gameComponents/controls"
	"GoPlat/gameComponents/levels"
	"GoPlat/gameComponents/sprites"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

/*func HandleAnimationVectorCalculations(lvl *levels.Level, p *sprites.Player, frameVec controls.Vector) controls.Vector {
	validMove := collision.IsValidMove(lvl, p, frameVec)
	if validMove {
		return frameVec
	}

	//while the animation vector is invalid, loop and
	//adjust vector until it is valid
	collisionData := collision.ExtractCollisionData(lvl)
	playerRect := collision.GetPlayerRect(p)

	var tempRect collision.Rect
	for !validMove {
		tempRect = playerRect
		if p.IsMovingRight {
			tempRect.X += frameVec.DeltaX
		} else {
			tempRect.X += frameVec.DeltaX
		}
		tempRect.Y += frameVec.DeltaY

		for _, coll := range collisionData {
			collisionLeft := collision.CheckXCollisionPlayerLeft(playerRect, coll)
			collisionRight := collision.CheckXCollisionPlayerRight(playerRect, coll)
			collisionTop := collision.CheckYCollisionPlayerTop(playerRect, coll)
			collisionBottom := collision.CheckYCollisionPlayerBottom(playerRect, coll)

			if collisionLeft && collisionRight { //if Y collision
				frameVec.BumpY()
				break
			} else if collisionTop && collisionBottom { //if X collision
				frameVec.BumpX(p.IsMovingRight)
				break
			} else if collisionLeft || collisionRight {
				frameVec.BumpX(p.IsMovingRight)
				break
			} else if collisionTop || collisionBottom {
				frameVec.BumpY()
				break
			}
		}

		validMove = collision.IsValidMove(lvl, p, frameVec)
	}

	return frameVec
}*/

func HandleMovementCalculations(p *sprites.Player, playerControls []controls.Control, lvl *levels.Level) controls.Vector2 {
	rtnVector := controls.Vector2{
		X: 0,
		Y: 0,
	}
	if p.IsAnimationLocked {
		AnimationLockWallOverride(p, lvl)
	}
	if p.IsAnimationLocked && !p.CanAnimationCancel {
		return rtnVector
	}
	directions := GetControlsPressed(playerControls)
	if p.IsAnimationLocked && !IsAnimationCancelling(p, directions) {
		return rtnVector
	}

	//handle movement
	playerVelocity, playerAcceleration, specialAction := GetMovementVector(directions)

	p.IsMovingRight = IsMovingRight(p)
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
				p.IsGravityLocked = true
			} else {
				p.IsGravityLocked = false
				//physics.HandlePhysics(p, lvl, &rtnVector)
			}
			return rtnVector
		}
	} else {
		p.IsGravityLocked = false
		p.CurrentAnimationIndex = 0
	}

	/*if !p.IsGravityLocked {
		physics.HandlePhysics(p, lvl, &playerVector)
	}*/

	rtnVector.X = playerVector.X
	rtnVector.Y = playerVector.Y

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

func GetMovementVector(directions []controls.Direction) (controls.Vector2, controls.Vector2, controls.Direction) {
	specialDirections := []controls.Direction{
		controls.JUMP,
		controls.DASHLEFT,
		controls.DASHRIGHT,
	}
	var velocity controls.Vector2
	var acceleration controls.Vector2
	var specialAction controls.Direction
	for _, direction := range directions {
		if math.Abs(direction.VelX) > math.Abs(velocity.X) {
			velocity.X = direction.VelX
		} else if math.Abs(direction.VelY) > math.Abs(velocity.Y) {
			velocity.Y = direction.VelY
		}
		for _, special := range specialDirections {
			if direction == special {
				specialAction = direction
			}
		}
		acceleration.X += direction.AccX
		acceleration.Y += direction.AccY
	}

	return velocity, acceleration, specialAction
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

func IsIdle(p *sprites.Player, vector controls.Vector2, specialAction controls.Direction) bool {
	if len(specialAction.Name) > 0 || p.IsWallHanging {
		return false
	} else if vector.X != 0 || vector.Y != 0 {
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
	} else {
		//if it isn't a wall or hurt or death animation, cancel that animation
		if (p.CurrentAnimationIndex < 4 || p.CurrentAnimationIndex > 8) ||
			(nearGround && p.CurrentAnimationIndex == 7) {
			p.ActionAnimations[p.CurrentAnimationIndex].CurrentFrameIndex = 0
			p.IsAnimationLocked = false
		}
	}
}
