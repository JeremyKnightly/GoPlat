package controls

type Direction struct {
	Name       string
	VelX, VelY float64
	AccX, AccY float64
}

var (
	LEFT = Direction{
		Name: "LEFT",
		VelX: -2,
		VelY: 0,
		AccX: 0,
		AccY: 0,
	}
	RIGHT = Direction{
		Name: "RIGHT",
		VelX: 2,
		VelY: 0,
		AccX: 0,
		AccY: 0,
	}
	//UP        = Direction{Name: "UP", X: 0, Y: -1}
	DASHLEFT = Direction{
		Name: "DASHLEFT",
		VelX: 0,
		VelY: 0,
		AccX: 0,
		AccY: 0,
	}
	DASHRIGHT = Direction{
		Name: "DASHRIGHT",
		VelX: 0,
		VelY: 0,
		AccX: 0,
		AccY: 0,
	}
	JUMP = Direction{
		Name: "JUMP",
		VelX: 0,
		VelY: 0,
		AccX: 0,
		AccY: 0,
	}
	WALLSLIDE = Direction{
		Name: "WALLSLIDE",
		VelX: 0,
		VelY: 4,
		AccX: 0,
		AccY: 0,
	}
)
