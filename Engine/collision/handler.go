package collision

import (
	controls "GoPlat/gameComponents/controls"
	levels "GoPlat/gameComponents/levels"
	"GoPlat/gameComponents/sprites"
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


func checkPlayerCollision(pRect Rect, coll Rect) bool {
	//Because of the way rects are drawn, X < X + Width. If either the player or the coll
	//has an x or x + width that fits between the opposite two, it is a collision.
	// example sides: P C P C== collision. however, it will not detect PCCP collisions 
	xCollision := false
	yCollision := false

	//player left in bounds
	if (pRect.X > coll.X && pRect.X < coll.X + coll.Width) {
		xCollision = true

		//player right in bounds
	} else if (pRect.X + pRect.Width > coll.X && pRect.X + pRect.Width < coll.X + coll.Width){
		xCollision = true

		//coll left in bounds
	} else if (coll.X > pRect.X  && coll.X < pRect.X + pRect.Width) {
		xCollision = true

		//coll right in bounds
	} else if (coll.X + coll.Width > pRect.X && coll.X + coll.Width < pRect.X + pRect.Width){
		xCollision = true
	} 

	//player left in bounds
	if (pRect.Y > coll.Y && pRect.Y < coll.Y + coll.Height) {
		yCollision = true

		//player right in bounds
	} else if (pRect.Y + pRect.Height > coll.Y && pRect.Y + pRect.Height < coll.Y + coll.Height){
		yCollision = true

		//coll left in bounds
	} else if (coll.Y > pRect.Y  && coll.Y < pRect.Y + pRect.Height) {
		yCollision = true

		//coll right in bounds
	} else if (coll.Y + coll.Height > pRect.Y && coll.Y + coll.Height < pRect.Y + pRect.Height){
		yCollision = true
	} 

	return xCollision && yCollision
}