package collision

import (
	controls "GoPlat/components/controls"
	levels "GoPlat/components/levels"
	"GoPlat/components/sprites"
)

func IsValidMove(lvl *levels.Level, player *sprites.Player, vector controls.Vector) bool {
	collisionData := ExtractCollisionData(lvl)
	
	playerRect := getPlayerRect(player)
	playerRect.X += vector.DeltaX
	playerRect.Y += vector.DeltaY

	for _,collision := range collisionData {
		colliding := checkPlayerCollision(playerRect, collision)
		if colliding {
			return false
		}
	}
	 return true
}

func getPlayerRect(p *sprites.Player) Rect {
	var playerRect Rect
	playerRect.X = p.X
	playerRect.Y = p.Y
	
	if p.IsIdle {
		playerRect.Width = p.IdleAnimation.MaxFrameWidth
		playerRect.Height = p.IdleAnimation.MaxFrameHeight
	} else {
		playerRect.Width = p.ActionAnimations[p.CurrentAnimationIndex].MaxFrameWidth
		playerRect.Height = p.ActionAnimations[p.CurrentAnimationIndex].MaxFrameHeight
	}

	return playerRect
}


func checkPlayerCollision(pRect Rect, coll Rect) bool {
	xCollision := (pRect.X > coll.X && pRect.X < coll.X + coll.Width) ||
	(pRect.X + pRect.Width > coll.X && pRect.X + pRect.Width < coll.X + coll.Width)

	yCollision := (pRect.Y > coll.Y && pRect.Y < coll.Y + coll.Height) ||
	(pRect.Y + pRect.Height > coll.Y && pRect.Y + pRect.Height < coll.Y + coll.Height)


	return xCollision && yCollision
}