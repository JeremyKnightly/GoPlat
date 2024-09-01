package controls

import "github.com/hajimehoshi/ebiten/v2"

type Control struct {
	Input     ebiten.StandardGamepadButton
	Key       ebiten.Key
	Keys      []ebiten.Key
	Direction Direction
	Type      string
}

func (c *Control) GetDirection() string {
	return c.Direction.Name
}

func (c *Control) GetType() string {
	return c.Type
}
