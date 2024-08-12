package collision

import (
	controls "GoPlat/gameComponents/controls"
	levels "GoPlat/gameComponents/levels"
	"GoPlat/gameComponents/sprites"
)

func GetXYCollisionBools(lvl *levels.Level, player sprites.Player, pVec controls.Vector) (bool, bool) {
	return true, true
}

func IsValidMoveRect(lvl *levels.Level, pRect Rect) bool {
	collisionData := ExtractCollisionData(lvl)

	for _, collision := range collisionData {
		collidingX, collidingY := CheckPlayerCollisionXY(pRect, collision)
		if collidingX && collidingY {
			return false
		}

	}
	return true
}

func IsValidMove(lvl *levels.Level, player *sprites.Player, vector controls.Vector) bool {
	collisionData := ExtractCollisionData(lvl)

	playerRect := GetPlayerRect(player)
	playerRect.X += vector.DeltaX
	playerRect.Y += vector.DeltaY

	for _, collision := range collisionData {
		collidingX, collidingY := CheckPlayerCollisionXY(playerRect, collision)
		if collidingX && collidingY {
			return false
		}

	}
	return true
}

func GetPlayerRect(p *sprites.Player) Rect {
	var playerRect Rect
	playerRect.X = p.X + 8
	playerRect.Y = p.Y + 8

	if p.IsIdle {
		playerRect.Width = p.IdleAnimation.MaxFrameWidth - 16
		playerRect.Height = p.IdleAnimation.MaxFrameHeight - 8
	} else {
		playerRect.Width = p.ActionAnimations[p.CurrentAnimationIndex].MaxFrameWidth - 16
		playerRect.Height = p.ActionAnimations[p.CurrentAnimationIndex].MaxFrameHeight - 8
	}

	return playerRect
}

func CheckPlayerCollisionXY(pRect Rect, coll Rect) (bool, bool) {
	xCollision := false
	yCollision := false

	if CheckXCollisionPlayerLeft(pRect, coll) {
		xCollision = true
	} else if CheckXCollisionPlayerRight(pRect, coll) {
		xCollision = true
	}

	if CheckYCollisionPlayerTop(pRect, coll) {
		yCollision = true
	} else if CheckYCollisionPlayerBottom(pRect, coll) {
		yCollision = true
	}

	return xCollision, yCollision
}

func CheckXCollisionPlayerLeft(pRect Rect, coll Rect) bool {
	//player left, coll right
	if pRect.X > coll.X && pRect.X < coll.X+coll.Width {
		return true

	} else if coll.X+coll.Width > pRect.X && coll.X+coll.Width < pRect.X+pRect.Width {
		return true
	}
	return false
}

func CheckXCollisionPlayerRight(pRect Rect, coll Rect) bool {
	//player right, coll left
	if pRect.X+pRect.Width > coll.X && pRect.X+pRect.Width < coll.X+coll.Width {
		return true
	} else if coll.X > pRect.X && coll.X < pRect.X+pRect.Width {
		return true
	}
	return false
}

func CheckYCollisionPlayerTop(pRect Rect, coll Rect) bool {
	//player top, coll bottom
	if pRect.Y > coll.Y && pRect.Y < coll.Y+coll.Height {
		return true
	} else if coll.Y+coll.Height > pRect.Y && coll.Y+coll.Height < pRect.Y+pRect.Height {
		return true
	}
	return false
}

func CheckYCollisionPlayerBottom(pRect Rect, coll Rect) bool {
	//player bottom, coll top
	if pRect.Y+pRect.Height > coll.Y && pRect.Y+pRect.Height < coll.Y+coll.Height {
		return true

		//coll top in bounds
	} else if coll.Y > pRect.Y && coll.Y < pRect.Y+pRect.Height {
		return true
	}
	return false
}
