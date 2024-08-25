package controls

type Direction struct {
	Name           string
	DeltaX, DeltaY float64
}

var (
	LEFT      = Direction{Name: "LEFT", DeltaX: -1.35, DeltaY: 0}
	RIGHT     = Direction{Name: "RIGHT", DeltaX: 1.35, DeltaY: 0}
	DASHLEFT  = Direction{Name: "DASHLEFT", DeltaX: 0, DeltaY: 0}
	DASHRIGHT = Direction{Name: "DASHRIGHT", DeltaX: 0, DeltaY: 0}
	JUMP      = Direction{Name: "JUMP", DeltaX: 0, DeltaY: 0}
	WALLSLIDE = Direction{Name: "WALLSLIDE", DeltaX: 0, DeltaY: 0}
)
