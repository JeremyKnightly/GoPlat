package controls

type Vector struct {
	DeltaX, DeltaY float64
}

func GetBlankVector() Vector {
	return Vector{
		DeltaX: 0,
		DeltaY: 0,
	}
}

func (vec *Vector) Add(posX float64, posY float64) Vector {
	newVec := Vector{
		posX + vec.DeltaX,
		posY + vec.DeltaY,
	}

	return newVec
}

func (vec *Vector) Subtract(posX float64, posY float64) Vector {
	newVec := Vector{
		posX - vec.DeltaX,
		posY - vec.DeltaY,
	}

	return newVec
}

func (vec *Vector) InvertX(posX float64, posY float64) Vector {
	newVec := Vector{
		posX - vec.DeltaX,
		posY + vec.DeltaY,
	}

	return newVec
}

func (vec *Vector) InvertY(posX float64, posY float64) Vector {
	newVec := Vector{
		posX + vec.DeltaX,
		posY - vec.DeltaY,
	}

	return newVec
}

func (vec *Vector) PlayerMove(posX float64, posY float64, isMovingRight bool) Vector {
	if isMovingRight {
		return vec.Add(posX, posY)

	} else {
		return vec.InvertX(posX, posY)
	}
}
