package collision

import (
	controls "GoPlat/gameComponents/controls"
	levels "GoPlat/gameComponents/levels"
	"GoPlat/gameComponents/sprites"
)

func IsValidMove(lvl *levels.Level, player *sprites.Player, velocity controls.Vector2) bool {
	collisionData := ExtractCollisionData(lvl)

	playerRect := GetPlayerRect(player)
	playerRect.X += velocity.X
	playerRect.Y += velocity.Y

	for _, collision := range collisionData {
		collidingX, collidingY := CheckPlayerCollisionXY(playerRect, collision)
		if collidingX && collidingY {
			return false
		}

	}
	return true
}

/*
func IsValidMoveRect(lvl *levels.Level, pRect Rect) bool {
	collisionData := ExtractCollisionData(lvl)

	for _, collision := range collisionData {
		collidingX, collidingY := CheckPlayerCollisionXY(pRect, collision)
		if collidingX && collidingY {
			return false
		}

	}
	return true
}*/

func GetPlayerRect(p *sprites.Player) Rect {
	var playerRect Rect
	playerRect.X = p.Physics.Position.X + 8
	playerRect.Y = p.Physics.Position.Y + 8

	if p.IsIdle {
		playerRect.Width = p.IdleAnimation.MaxFrameWidth - 16
		playerRect.Height = p.IdleAnimation.MaxFrameHeight - 8
	} else {
		playerRect.Width = p.ActionAnimations[p.CurrentAnimationIndex].MaxFrameWidth - 16
		playerRect.Height = p.ActionAnimations[p.CurrentAnimationIndex].MaxFrameHeight - 8
	}

	return playerRect
}
