package controls

import "math"

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
	return vec.Add(posX, posY)
}

func (vec *Vector) BumpY() {
	if vec.DeltaY > 0 {
		vec.DeltaY -= .05 * vec.DeltaY
	} else if vec.DeltaY < 0 {
		vec.DeltaY += .05 * vec.DeltaY
	}
}

func (vec *Vector) BumpX(movingRight bool) {
	if math.Abs(vec.DeltaX) <= 0 {
		return
	}
	if !movingRight {
		vec.DeltaX += .05
		println("movingleft")
	} else {
		vec.DeltaX -= .05
	}
	println(vec.DeltaX)
}
