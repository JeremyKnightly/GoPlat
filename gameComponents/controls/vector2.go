package controls

import (
	"math"
)

const Gravity = 9.81 * 5
const FrictionCoefficient = .9

type Vector2 struct {
	X, Y float64
}

type PhysicsObj struct {
	NetForce     Vector2
	Position     Vector2
	Velocity     Vector2
	Acceleration Vector2
	Mass         float64
}

func (obj *PhysicsObj) UpdatePhysics(deltaTime float64) {
	if deltaTime <= 0 {
		deltaTime = 1 / 60
	}
	obj.NetForce.Y += Gravity
	//println(obj.NetForce.X, "  ,  ", obj.NetForce.Y, "\n")
	obj.Acceleration.X = obj.NetForce.X / obj.Mass
	obj.Acceleration.Y = obj.NetForce.Y / obj.Mass
	//fmt.Printf("Acceleration: X: %v, Y: %v", obj.Acceleration.X, obj.Acceleration.Y)

	deltaVelX := obj.Acceleration.X * deltaTime
	deltaVelY := obj.Acceleration.Y * deltaTime
	if math.IsNaN(deltaVelX) || math.IsNaN(deltaVelY) ||
		math.IsInf(deltaVelX, 0) || math.IsInf(deltaVelY, 0) {
		return
	}

	obj.Velocity.X = deltaVelX
	obj.Velocity.Y = deltaVelY

	obj.Position.X += obj.Velocity.X * deltaTime
	obj.Position.Y += obj.Velocity.Y * deltaTime

	//obj.NetForce.X = 0
	//obj.NetForce.Y = 0
	//obj.Acceleration = Vector2{}
}
