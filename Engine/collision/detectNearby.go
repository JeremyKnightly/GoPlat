package collision

import (
	"GoPlat/gameComponents/levels"
	"GoPlat/gameComponents/sprites"
)

func DetectWall(player *sprites.Player, lvl *levels.Level) (bool, bool) {
	collisionData := ExtractCollisionData(lvl)
	playerRect := getPlayerRect(player)

	for _, collision := range collisionData {
		wallNearby, isLeft := checkWallNearby(playerRect, collision)
		if wallNearby {
			return wallNearby, isLeft
		}
	}

	return false, false
}

func DetectGround(player *sprites.Player, lvl *levels.Level) bool {
	collisionData := ExtractCollisionData(lvl)
	playerRect := getPlayerRect(player)

	for _, collision := range collisionData {
		groundNearby := checkGroundNearby(playerRect, collision)
		if groundNearby {
			return true
		}
	}

	return false
}

func checkGroundNearby(pRect Rect, coll Rect) bool {
	// no X collision means there is no collision on the same vertical plane
	if !checkXCollisionPlayerLeft(pRect, coll) && !checkXCollisionPlayerRight(pRect, coll) {
		return false
	}

	pRect.Height += 2
	return checkYCollisionPlayerBottom(pRect, coll)
}

func checkWallNearby(pRect Rect, coll Rect) (bool, bool) {
	// no Y collision means there is no collision on the same horizontal plane
	if !checkYCollisionPlayerTop(pRect, coll) && !checkYCollisionPlayerBottom(pRect, coll) {
		return false, false
	}

	//increase width and shift left for detection range
	pRect.Width += 4
	pRect.X -= 2
	if checkXCollisionPlayerRight(pRect, coll) {
		return true, false
	} else if checkXCollisionPlayerLeft(pRect, coll) {
		return true, true
	}

	return false, false
}
