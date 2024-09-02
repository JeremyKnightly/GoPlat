package startup

import (
	controls "GoPlat/gameComponents/controls"

	"github.com/hajimehoshi/ebiten/v2"
)

func GetControls() []controls.Control {
	//gamepad support
	GPLeft1 := controls.Control{
		InputAxis: ebiten.StandardGamepadAxisRightStickHorizontal,
		Direction: controls.LEFT,
		InputType: "AXIS",
		Type:      "Gamepad",
	}
	GPRight1 := controls.Control{
		InputAxis: ebiten.StandardGamepadAxisRightStickHorizontal,
		Direction: controls.RIGHT,
		InputType: "AXIS",
		Type:      "Gamepad",
	}
	GPLeft2 := controls.Control{
		InputAxis: ebiten.StandardGamepadAxisLeftStickHorizontal,
		Direction: controls.LEFT,
		InputType: "AXIS",
		Type:      "Gamepad",
	}
	GPRight2 := controls.Control{
		InputAxis: ebiten.StandardGamepadAxisLeftStickHorizontal,
		Direction: controls.RIGHT,
		InputType: "AXIS",
		Type:      "Gamepad",
	}
	GPDashRight := controls.Control{
		Input:     ebiten.StandardGamepadButtonFrontTopRight,
		Direction: controls.DASHRIGHT,
		InputType: "BTN",
		Type:      "Gamepad",
	}
	GPDashLeft := controls.Control{
		Input:     ebiten.StandardGamepadButtonFrontTopLeft,
		Direction: controls.DASHLEFT,
		InputType: "BTN",
		Type:      "Gamepad",
	}
	GPJump := controls.Control{
		Input:     ebiten.StandardGamepadButtonRightRight,
		Direction: controls.JUMP,
		InputType: "BTN",
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
