package controls

const gravity = 9.81

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

	obj.Acceleration = Vector2{}
}
