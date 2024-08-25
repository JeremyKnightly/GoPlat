package controls

type Direction struct {
	Name           string
<<<<<<< HEAD
	DeltaX, DeltaY float64
}

var (
	LEFT      = Direction{Name: "LEFT", DeltaX: -1.35, DeltaY: 0}
	RIGHT     = Direction{Name: "RIGHT", DeltaX: 1.35, DeltaY: 0}
	DASHLEFT  = Direction{Name: "DASHLEFT", DeltaX: 0, DeltaY: 0}
	DASHRIGHT = Direction{Name: "DASHRIGHT", DeltaX: 0, DeltaY: 0}
	JUMP      = Direction{Name: "JUMP", DeltaX: 0, DeltaY: 0}
	WALLSLIDE = Direction{Name: "WALLSLIDE", DeltaX: 0, DeltaY: 0}
=======
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
		ForceY: -50,
	}
>>>>>>> 16f3c53bd428513a1c986c0fe9c23443b6469d9c
)
