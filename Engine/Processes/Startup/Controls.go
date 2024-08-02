package startup

import (
	controls "GoPlat/components/controls"

	"github.com/hajimehoshi/ebiten/v2"
)

func GetControls() []controls.Control {
	w := controls.Control{
		Key: ebiten.KeyW,
		Direction: controls.UP,
	}
	up := controls.Control{
		Key: ebiten.KeyArrowUp,
		Direction: controls.UP,
	}
	a := controls.Control{
		Key: ebiten.KeyA,
		Direction: controls.LEFT,
	}
	left := controls.Control{
		Key: ebiten.KeyArrowLeft,
		Direction: controls.LEFT,
	}
	d := controls.Control{
		Key: ebiten.KeyD,
		Direction: controls.RIGHT,
	}
	right := controls.Control{
		Key: ebiten.KeyArrowRight,
		Direction: controls.RIGHT,
	}
	s := controls.Control{
		Key: ebiten.KeyS,
		Direction: controls.DOWN,
	}
	down := controls.Control{
		Key: ebiten.KeyArrowDown,
		Direction: controls.DOWN,
	}
	shiftD := controls.Control {
		Keys: []ebiten.Key{
			ebiten.KeyShift,
			ebiten.KeyD,
		},
		Direction: controls.DASHRIGHT,
	}
	shiftR := controls.Control {
		Keys: []ebiten.Key{
			ebiten.KeyShift,
			ebiten.KeyArrowRight,
		},
		Direction: controls.DASHRIGHT,
	}
	shiftA := controls.Control {
		Keys: []ebiten.Key{
			ebiten.KeyShift,
			ebiten.KeyA,
		},
		Direction: controls.DASHLEFT,
	}
	shiftL := controls.Control {
		Keys: []ebiten.Key{
			ebiten.KeyShift,
			ebiten.KeyArrowLeft,
		},
		Direction: controls.DASHLEFT,
	}
	space := controls.Control {
		Key: ebiten.KeySpace,
		Direction: controls.JUMP,
	}

	return []controls.Control{
		w,
		up,
		a,
		left,
		d,
		right,
		s,
		down,
		shiftD,
		shiftR,
		shiftA,
		shiftL,
		space,
	}
}