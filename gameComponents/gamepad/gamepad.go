package gamepad

import (
	"GoPlat/gameComponents/controls"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Gamepad struct {
	gamepadIDsBuf  []ebiten.GamepadID
	gamepadIDs     map[ebiten.GamepadID]struct{}
	axes           map[ebiten.GamepadID][]string
	pressedButtons map[ebiten.GamepadID][]string
}

func GetNewGamepad() *Gamepad {
	return &Gamepad{}
}

func (gp *Gamepad) CheckGamepadConnection() {
	if gp.gamepadIDs == nil {
		gp.gamepadIDs = map[ebiten.GamepadID]struct{}{}
	}

	gp.gamepadIDsBuf = inpututil.AppendJustConnectedGamepadIDs(gp.gamepadIDsBuf[:0])
	for _, id := range gp.gamepadIDsBuf {
		gp.gamepadIDs[id] = struct{}{}
	}
	for id := range gp.gamepadIDs {
		if inpututil.IsGamepadJustDisconnected(id) {
			delete(gp.gamepadIDs, id)
		}
	}
}

func (gp *Gamepad) UpdateInput() {
	gp.axes = map[ebiten.GamepadID][]string{}
	gp.pressedButtons = map[ebiten.GamepadID][]string{}
}

func (gp *Gamepad) AxisMatchesInput(control controls.Control) bool {
	gp.CheckGamepadConnection()
	gp.UpdateInput()

	for id := range gp.gamepadIDs {
		var axisDirection float64
		maxAxis := ebiten.GamepadAxisType(ebiten.GamepadAxisCount(id))
		for axisType := ebiten.GamepadAxisType(0); axisType < maxAxis; axisType++ {
			if axisType != int(control.InputAxis) {
				continue
			}
			axisDirection = ebiten.GamepadAxisValue(id, axisType)
			if control.Direction.Name == "RIGHT" && axisDirection > 0.1 {
				return true
			} else if control.Direction.Name == "LEFT" && axisDirection < -0.1 {
				return true
			}
		}
	}
	return false
}

func (gp *Gamepad) IsButtonPressed(button ebiten.StandardGamepadButton) bool {
	gp.CheckGamepadConnection()
	gp.UpdateInput()

	for id := range gp.gamepadIDs {
		if !ebiten.IsStandardGamepadLayoutAvailable(id) {
			continue
		}

		for b := ebiten.StandardGamepadButton(0); b <= ebiten.StandardGamepadButtonMax; b++ {
			if inpututil.IsStandardGamepadButtonJustPressed(id, button) {
				return true
			}
		}
	}
	return false
}
