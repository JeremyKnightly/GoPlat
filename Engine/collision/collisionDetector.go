package collision

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
