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

func (vec *Vector) BumpY() {
	if vec.DeltaY > 0 {
		vec.DeltaY -= .03
	} else if vec.DeltaY < 0 {
		vec.DeltaY += .03
	}
}

func (vec *Vector) BumpX(movingRight bool) {
	if !movingRight {
		vec.DeltaX += .05
	} else {
		vec.DeltaX -= .05
	}
}

func (vec *Vector) ScaleByTPS(TicksThisFrame, TicksPerFrame float64) Vector {
	rtnVec := GetBlankVector()
	if TicksPerFrame == 0 {
		return rtnVec
	}
	multiplier := TicksThisFrame / TicksPerFrame
	rtnVec.Scale(multiplier)

	return rtnVec
}

func (vec *Vector) Scale(multiplier float64) {
	vec.DeltaX = vec.DeltaX * multiplier
	vec.DeltaY = vec.DeltaY * multiplier
}
