package startup

import (
	controls "GoPlat/gameComponents/controls"

	"github.com/hajimehoshi/ebiten/v2"
)

func GetControls() []controls.Control {
	//gamepad support
	GPLeft1 := controls.Control{
		Input:     ebiten.StandardGamepadButtonCenterLeft,
		Direction: controls.LEFT,
		Type:      "Gamepad",
	}
	GPRight1 := controls.Control{
		Input:     ebiten.StandardGamepadButtonCenterRight,
		Direction: controls.RIGHT,
		Type:      "Gamepad",
	}
	GPLeft2 := controls.Control{
		Input:     ebiten.StandardGamepadButtonLeftLeft,
		Direction: controls.LEFT,
		Type:      "Gamepad",
	}
	GPRight2 := controls.Control{
		Input:     ebiten.StandardGamepadButtonLeftRight,
		Direction: controls.RIGHT,
		Type:      "Gamepad",
	}
	GPDashRight := controls.Control{
		Input:     ebiten.StandardGamepadButtonFrontTopRight,
		Direction: controls.DASHRIGHT,
		Type:      "Gamepad",
	}
	GPDashLeft := controls.Control{
		Input:     ebiten.StandardGamepadButtonFrontTopLeft,
		Direction: controls.DASHLEFT,
		Type:      "Gamepad",
	}
	GPJump := controls.Control{
		Input:     ebiten.StandardGamepadButtonRightRight,
		Direction: controls.JUMP,
		Type:      "Gamepad",
	}

	//keyboard support
	left := controls.Control{
		Key:       ebiten.KeyA,
		Direction: controls.LEFT,
		Type:      "Keyboard",
	}
	right := controls.Control{
		Key:       ebiten.KeyD,
		Direction: controls.RIGHT,
		Type:      "Keyboard",
	}
	dashRight := controls.Control{
		Keys: []ebiten.Key{
			ebiten.KeyShift,
			ebiten.KeyD,
		},
		Direction: controls.DASHRIGHT,
		Type:      "Keyboard",
	}
	dashLeft := controls.Control{
		Keys: []ebiten.Key{
			ebiten.KeyShift,
			ebiten.KeyA,
		},
		Direction: controls.DASHLEFT,
		Type:      "Keyboard",
	}
	space := controls.Control{
		Key:       ebiten.KeySpace,
		Direction: controls.JUMP,
		Type:      "Keyboard",
	}

	return []controls.Control{
		//Gamepad
		GPLeft1,
		GPLeft2,
		GPRight1,
		GPRight2,
		GPDashLeft,
		GPDashRight,
		GPJump,

		//keyboard
		left,
		right,
		dashLeft,
		dashRight,
		space,
	}
}
