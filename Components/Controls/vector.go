package controls

type Vector struct {
	DeltaX, DeltaY float64
}

func AddVector(posX float64, posY float64, vector Vector) (float64,float64){
	posX += vector.DeltaX
	posY += vector.DeltaY

	return posX, posY
}