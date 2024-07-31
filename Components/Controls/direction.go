package controls

type Direction struct {
	Name           string
	DeltaX, DeltaY float64
}

var (
	LEFT      = Direction{Name: "LEFT", DeltaX: -2, DeltaY: 0}
	RIGHT     = Direction{Name: "RIGHT", DeltaX: 2, DeltaY: 0}
	UP        = Direction{Name: "UP", DeltaX: 0, DeltaY: -2}
	DOWN      = Direction{Name: "DOWN", DeltaX: 0, DeltaY: 2}
	DASHLEFT  = Direction{Name: "DASHLEFT", DeltaX: -3, DeltaY: 0}
	DASHRIGHT = Direction{Name: "DASHRIGHT", DeltaX: 3, DeltaY: 0}
	JUMP      = Direction{Name: "JUMP", DeltaX: 0, DeltaY: -4}
)