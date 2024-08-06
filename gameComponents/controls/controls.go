package controls

import "github.com/hajimehoshi/ebiten/v2"

type Control struct {
	Key ebiten.Key
	Keys []ebiten.Key
	Direction Direction
}

func (c *Control) GetDirection() string {
	return c.Direction.Name
}