package camera

import "math"

type Camera struct {
	X, Y float64
}

func NewCamera(x, y float64) *Camera {
	return &Camera{
		X: x,
		Y: y,
	}
}

func (c *Camera) FollowTarget(playerX, playerY, screenWidth, screenHeight float64) {
	c.X = -playerX + screenWidth/2
	c.Y = -playerY + screenHeight/2
}

func (c *Camera) Constrain(tilemapHeightPixels, tilemapWidthPixels, screenWidth, screenHeight float64) {
	c.X = math.Min(c.X, 0.0)
	c.Y = math.Min(c.Y, 0.0)

	c.X = math.Max(c.X, screenWidth-tilemapWidthPixels)
	c.Y = math.Max(c.Y, screenHeight-tilemapHeightPixels)
}
