package startup

import (
	controls "GoPlat/gameComponents/controls"

	"github.com/hajimehoshi/ebiten/v2"
)

func GetControls() []controls.Control {
	w := controls.Control{
		Key: ebiten.KeyW,
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
	dashRight := controls.Control{
		Key: ebiten.KeyE,
		Direction: controls.DASHRIGHT,
	}
	dashLeft := controls.Control{
		Key: ebiten.KeyQ,
		Direction: controls.DASHLEFT,
	}
	space := controls.Control {
		Key: ebiten.KeySpace,
		Direction: controls.JUMP,
	}

	return []controls.Control{
		w,
		a,
		left,
		d,
		right,
		s,
		down,
		dashRight,
		dashLeft,
		space,
	}
}