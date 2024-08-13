package controls

import (
	"math"
)

const Gravity = 9.81 * 2
const FrictionCoefficient = .85

type Vector2 struct {
	X, Y float64
}

type PhysicsObj struct {
	Position     Vector2
	Velocity     Vector2
	Acceleration Vector2
	Mass         float64
}

func (obj *PhysicsObj) UpdatePhysics(deltaTime float64) {
	if deltaTime <= 0 {
		deltaTime = 1 / 60
	}
	obj.Acceleration.Y += Gravity

	deltaVelX := obj.Acceleration.X * deltaTime
	deltaVelY := obj.Acceleration.Y * deltaTime
	if math.IsNaN(deltaVelX) || math.IsNaN(deltaVelY) ||
		math.IsInf(deltaVelX, 0) || math.IsInf(deltaVelY, 0) {
		return
	}
	println(obj.Velocity.X)
	obj.Velocity.X += deltaVelX
	obj.Velocity.Y += deltaVelY
	obj.Velocity.X *= FrictionCoefficient

	obj.Position.X += obj.Velocity.X * deltaTime
	obj.Position.Y += obj.Velocity.Y * deltaTime

	obj.Acceleration = Vector2{}
}
