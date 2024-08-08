package collision

import (
	"GoPlat/gameComponents/levels"
	"GoPlat/gameComponents/sprites"
)

func DetectWall(player *sprites.Player, lvl *levels.Level) (bool, bool) {
	collisionData := ExtractCollisionData(lvl)
	playerRect := GetPlayerRect(player)

	for _, collision := range collisionData {
		wallNearby, isLeft := CheckWallNearby(playerRect, collision)
		if wallNearby {
			return wallNearby, isLeft
		}
	}

	return false, false
}

func DetectGround(player *sprites.Player, lvl *levels.Level) bool {
	collisionData := ExtractCollisionData(lvl)
	playerRect := GetPlayerRect(player)

	for _, collision := range collisionData {
		groundNearby := CheckGroundNearby(playerRect, collision)
		if groundNearby {
			return true
		}
	}

	return false
}

func CheckGroundNearby(pRect Rect, coll Rect) bool {
	// no X collision means there is no collision on the same vertical plane
	if !CheckXCollisionPlayerLeft(pRect, coll) && !CheckXCollisionPlayerRight(pRect, coll) {
		return false
	}

	pRect.Height += 2
	return CheckYCollisionPlayerBottom(pRect, coll)
}

func CheckWallNearby(pRect Rect, coll Rect) (bool, bool) {
	// no Y collision means there is no collision on the same horizontal plane
	if !CheckYCollisionPlayerTop(pRect, coll) && !CheckYCollisionPlayerBottom(pRect, coll) {
		return false, false
	}

	//increase width and shift left for detection range
	pRect.Width += 4
	pRect.X -= 2
	if CheckXCollisionPlayerRight(pRect, coll) {
		return true, false
	} else if CheckXCollisionPlayerLeft(pRect, coll) {
		return true, true
	}

	return false, false
}
