package collision

import (
	controls "GoPlat/gameComponents/controls"
	levels "GoPlat/gameComponents/levels"
	"GoPlat/gameComponents/sprites"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func MenuStateChange(menu *levels.Level) (bool, int) {
	collisionData := ExtractCollisionData(menu)
	x, y := ebiten.CursorPosition()
	cursorRect := NewBlankRect(float64(x), float64(y))

	for _, collision := range collisionData {
		if !collision.HasProp("Clickable") {
			continue
		}
		x, y := getXYOverlap(cursorRect, collision)
		if x > 0 && y > 0 && collision.HasProp("InitiateState") {
			return true, int(collision.GetPropValue("InitiateState").(float64))
		}
	}

	return false, 0
}

func GetWallHangCoords(lvl *levels.Level, pRect Rect, wallPlayerLeft bool) (float64, float64) {
	//I know that the player is near a wall
	//I need to check if there is space above for the player
	collisionData := ExtractCollisionData(lvl)

	pRectUp := pRect
	pRectDwn := pRect
	pRectUp.Y -= 1.2
	pRectDwn.Y += 1.2

	for _, collision := range collisionData {
		colliding := IsCollidingNoSpecial(pRect, collision)
		collidingUp := IsCollidingNoSpecial(pRectUp, collision)
		collidingDwn := IsCollidingNoSpecial(pRectDwn, collision)
		// if all collide, skip
		if colliding && collidingUp && collidingDwn {
			continue
			// if none collide, skip
		} else if !(colliding || collidingUp || collidingDwn) {
			continue
			// if one collides, return it
		} else {
			yIsClose := math.Abs(math.Abs(collision.Y)-math.Abs(pRect.Y)-pRect.Height) <= 14
			var xIsClose bool
			if wallPlayerLeft {
				xIsClose = math.Abs(pRect.X-collision.Width-collision.X) <= 20
			} else {
				xIsClose = math.Abs(collision.X-pRect.Width-pRect.X) <= 20
			}
			if yIsClose && xIsClose {
				return collision.X - pRect.Width/2 - 6, collision.Y
			}
		}
	}
	return 0, 0
}

func EnsureValidMove(lvl *levels.Level, player *sprites.Player, vector controls.Vector) {
	collisionData := ExtractCollisionData(lvl)

	yVecPos := vector.DeltaY >= 0
	xVecPos := vector.DeltaX >= 0
	playerRect := GetPlayerRect(player)
	playerRect.X += vector.DeltaX
	playerRect.Y += vector.DeltaY
	player.X += vector.DeltaX
	player.Y += vector.DeltaY

	for _, collision := range collisionData {
		doCollision(&playerRect, &collision, player, xVecPos, yVecPos)
	}
}

func getXOverlap(pRect, coll Rect) float64 {
	minPoint := math.Min(pRect.X, coll.X)
	maxPoint := math.Max(pRect.X+pRect.Width, coll.X+coll.Width)
	totalWidth := (pRect.Width + coll.Width)
	actualWidth := (maxPoint - minPoint)
	return (totalWidth - actualWidth)
}

func getYOverlap(pRect, coll Rect) float64 {
	minPoint := math.Min(pRect.Y, coll.Y)
	maxPoint := math.Max(pRect.Y+pRect.Height, coll.Y+coll.Height)
	return ((pRect.Height + coll.Height) - (maxPoint - minPoint))
}

func getXYOverlap(pRect, coll Rect) (float64, float64) {
	return getXOverlap(pRect, coll), getYOverlap(pRect, coll)
}

func doCollision(pRect, coll *Rect, player *sprites.Player, xVecPos, yVecPos bool) {
	overlapX, overlapY := getXYOverlap(*pRect, *coll)

	if overlapX <= 0 || overlapY <= 0 {
		return //all non-collisions end here
	}

	if coll.HasSpecialProps() {
		coll.HandleProps(player)
		return
	}

	xPercentOverlap := overlapX / pRect.Width
	yPercentOverlap := overlapY / pRect.Height
	isXOverlap := xPercentOverlap < yPercentOverlap

	if isXOverlap {
		if player.X < coll.X && xVecPos {
			player.X -= overlapX + .25
			pRect.X -= overlapX + .25
		} else {
			player.X += overlapX + .25
			pRect.X += overlapX + .25
		}
	} else {
		if player.Y < coll.Y && yVecPos {
			player.Y -= overlapY + .25
			pRect.Y -= overlapY + .25
		} else {
			player.Y += overlapY + .25
			pRect.Y += overlapY + .25
		}
	}
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

func IsCollidingNoSpecial(pRect Rect, coll Rect) bool {
	if coll.HasSpecialProps() {
		return false
	}
	xOverlap, yOverlap := getXYOverlap(pRect, coll)
	return xOverlap > 0 && yOverlap > 0
}

func CheckXCollisionPlayerLeft(pRect Rect, coll Rect) bool {
	return pRect.X > coll.X && pRect.X < coll.X+coll.Width
}

func CheckXCollisionPlayerRight(pRect Rect, coll Rect) bool {
	return pRect.X+pRect.Width > coll.X && pRect.X+pRect.Width < coll.X+coll.Width
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
