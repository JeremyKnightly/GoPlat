package collision

import (
	controls "GoPlat/gameComponents/controls"
	"GoPlat/gameComponents/levels"
	"GoPlat/gameComponents/sprites"
	"math"
)

func ResolveCollisions(lvl *levels.Level, player *sprites.Player) {
	collisionData := ExtractCollisionData(lvl)
	pRect := GetPlayerRect(player)

	for _, collision := range collisionData {
		collidingX, collidingY := CheckPlayerCollisionXY(pRect, collision)
		if collidingX && collidingY {
			correctorVector := resolveCollision(pRect, collision)
			player.Physics.Velocity.X = 0
			player.Physics.Velocity.Y = 0
			player.Physics.Position.X += correctorVector.X
			player.Physics.Position.Y += correctorVector.Y
			pRect = GetPlayerRect(player)
		}
	}
}

func resolveCollision(pRect Rect, collision Rect) controls.Vector2 {
	overlapX := ((pRect.Width + collision.Width) - math.Max(pRect.Width, collision.Width)) / 2
	overlapY := ((pRect.Height + collision.Height) - math.Max(pRect.Height, collision.Height)) / 2

	var correctorVector controls.Vector2
	//shortest overlap is quickest way to resolve collision
	if overlapX < overlapY {
		if pRect.X < collision.X {
			correctorVector.X -= overlapX
		} else {
			correctorVector.X += overlapX
		}
	} else {
		if pRect.Y < collision.Y {
			correctorVector.Y -= overlapY
		} else {
			correctorVector.Y += overlapY
		}
	}
	return correctorVector
}
