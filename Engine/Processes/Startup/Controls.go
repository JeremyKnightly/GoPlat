package startup

import (
	controls "GoPlat/gameComponents/controls"

	"github.com/hajimehoshi/ebiten/v2"
)

func GetControls() []controls.Control {
	a := controls.Control{
		Key: ebiten.KeyA,
		Direction: controls.LEFT,
	}
	d := controls.Control{
		Key: ebiten.KeyD,
		Direction: controls.RIGHT,
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
		a,
		d,
		dashRight,
		dashLeft,
		space,
	}
}