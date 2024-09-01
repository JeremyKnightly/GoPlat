package gamepad

import (
	"log"
	"math"

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
		log.Printf("gamepad connected: id: %d, SDL ID: %s", id, ebiten.GamepadSDLID(id))
		gp.gamepadIDs[id] = struct{}{}
	}
	for id := range gp.gamepadIDs {
		if inpututil.IsGamepadJustDisconnected(id) {
			log.Printf("gamepad disconnected: id: %d", id)
			delete(gp.gamepadIDs, id)
		}
	}
}

func (gp *Gamepad) UpdateInput() {
	gp.axes = map[ebiten.GamepadID][]string{}
	gp.pressedButtons = map[ebiten.GamepadID][]string{}
}

func (gp *Gamepad) IsButtonPressed(button ebiten.StandardGamepadButton) bool {
	gp.CheckGamepadConnection()
	gp.UpdateInput()

	for id := range gp.gamepadIDs {
		var axisDirection float64
		maxAxis := ebiten.GamepadAxisType(ebiten.GamepadAxisCount(id))
		for a := ebiten.GamepadAxisType(0); a < maxAxis; a++ {
			axisDirection = ebiten.GamepadAxisValue(id, a)
			//g.axes[id] = append(g.axes[id], fmt.Sprintf("%d:%+0.2f", a, v))
		}
		if math.Abs(axisDirection) > 0 {
			if button == ebiten.StandardGamepadButtonCenterRight && axisDirection > 0 {
				println("moving right")
				return true
			}
		}
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
