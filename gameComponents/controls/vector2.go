package controls

import "fmt"

const gravity = 9.81 / 6

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
	obj.Acceleration.Y += gravity

	obj.Velocity.X += obj.Acceleration.X * deltaTime
	obj.Velocity.Y += obj.Acceleration.Y * deltaTime

	obj.Position.X += obj.Velocity.X * deltaTime
	obj.Position.Y += obj.Velocity.Y * deltaTime

	fmt.Printf("velocity:%v,%v\n", obj.Velocity.X, obj.Velocity.Y)
	fmt.Printf("position:%v,%v\n", obj.Position.X, obj.Position.Y)
	fmt.Printf("acceleration:%v,%v\n", obj.Acceleration.X, obj.Acceleration.Y)
	obj.Acceleration = Vector2{}
}
