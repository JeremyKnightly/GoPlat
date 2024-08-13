package controls

type Direction struct {
	Name           string
	ForceX, ForceY float64
}

var (
	LEFT = Direction{
		Name:   "LEFT",
		ForceX: -50,
		ForceY: 0,
	}
	RIGHT = Direction{
		Name:   "RIGHT",
		ForceX: 50,
		ForceY: 0,
	}
	//UP        = Direction{Name: "UP", X: 0, Y: -1}
	DASHLEFT = Direction{
		Name:   "DASHLEFT",
		ForceX: -250,
		ForceY: 0,
	}
	DASHRIGHT = Direction{
		Name:   "DASHRIGHT",
		ForceX: 250,
		ForceY: 0,
	}
	JUMP = Direction{
		Name:   "JUMP",
		ForceX: 0,
		ForceY: -350,
	}
	WALLSLIDE = Direction{
		Name:   "WALLSLIDE",
		ForceX: 0,
		ForceY: -500,
	}
)
