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
			correctorVector, horizontalCollision := resolveCollision(pRect, collision)
			if horizontalCollision {
				player.Physics.Velocity.X = 0
			} else {
				player.Physics.Velocity.Y = 0
			}
			player.Physics.Position.X += correctorVector.X
			player.Physics.Position.Y += correctorVector.Y
			pRect = GetPlayerRect(player)
		}
	}
}

func resolveCollision(pRect Rect, collision Rect) (controls.Vector2, bool) {
	maxPoint := math.Max(pRect.X+pRect.Width, collision.X+collision.Width)
	minPoint := math.Min(pRect.X, collision.X)
	overlapX := (pRect.Width + collision.Width) - math.Abs(maxPoint-minPoint)

	maxPoint = math.Max(pRect.Y+pRect.Height, collision.Y+collision.Height)
	minPoint = math.Min(pRect.Y, collision.Y)
	overlapY := ((pRect.Height + collision.Height) - math.Abs(maxPoint-minPoint))

	var correctorVector controls.Vector2
	var horizontalCollision bool
	//shortest overlap is quickest way to resolve collision
	if overlapX < overlapY {
		horizontalCollision = true
		if pRect.X < collision.X {
			correctorVector.X -= overlapX
		} else {
			correctorVector.X += overlapX + .01 //for rounding error moving left
		}
	} else {
		horizontalCollision = false
		if pRect.Y < collision.Y {
			correctorVector.Y -= overlapY
		} else {
			correctorVector.Y += overlapY
		}
	}
	return correctorVector, horizontalCollision
}
