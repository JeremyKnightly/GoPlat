package collision

import (
	"GoPlat/components/levels"
	"GoPlat/components/sprites"
	"math"
)

func GetPlayerCorners(p *sprites.Player, tileSize float64, layer levels.TilemapLayer) (float64, float64, float64, float64) {
	left := math.Ceil(p.X / tileSize)
	top := math.Ceil(p.Y / tileSize)
	right := 0.0
	bottom := 0.0

	if p.IsIdle {
		right = math.Floor((p.X + p.IdleAnimation.MaxFrameWidth) / tileSize)
		bottom = math.Floor((p.Y + p.IdleAnimation.MaxFrameHeight) / tileSize)
	} else {
		right = math.Floor((p.X + p.ActionAnimations[p.CurrentAnimationIndex].MaxFrameWidth) / tileSize)
		bottom = math.Floor((p.Y + p.ActionAnimations[p.CurrentAnimationIndex].MaxFrameHeight) / tileSize)
	}

	return left, top, right, bottom
}

func CheckCollision(layer levels.TilemapLayer, left float64, top float64, right float64, bottom float64) []bool {
	collisionMap := []bool{}

	//TopLeft Collision
	collisionMap = append(collisionMap, layer.Data[uint((top - 1)*layer.Width+left - 1)] > 0)
	//TopRight Collision
	collisionMap = append(collisionMap, layer.Data[uint((top - 1)*layer.Width+right)] > 0)
	//BottomLeft Collision
	collisionMap = append(collisionMap, layer.Data[uint((bottom)*layer.Width+left - 1)] > 0)
	//BottomRight Collision
	collisionMap = append(collisionMap, layer.Data[uint((bottom)*layer.Width+right)] > 0)

	return collisionMap
}

func GetPlayerCorners2(p *sprites.Player, layer levels.TilemapLayer) (float64, float64, float64, float64) {
	left := p.X
	top := p.Y
	right := 0.0
	bottom := 0.0

	if p.IsIdle {
		right = (p.X + p.IdleAnimation.MaxFrameWidth)
		bottom = (p.Y + p.IdleAnimation.MaxFrameHeight)
	} else {
		right = (p.X + p.ActionAnimations[p.CurrentAnimationIndex].MaxFrameWidth) 
		bottom = (p.Y + p.ActionAnimations[p.CurrentAnimationIndex].MaxFrameHeight)
	}

	return left, top, right, bottom
}

func CheckCollision2(layer levels.TilemapLayer, left float64, top float64, right float64, bottom float64, tileSize float64) []bool {
	collisionMap := []bool{}

	//TopLeft Collision
	collisionMap = append(collisionMap, layer.Data[uint(math.Floor(((top - 1)*layer.Width+left - 1)/tileSize))] > 0)
	//TopRight Collision
	collisionMap = append(collisionMap, layer.Data[uint(math.Ceil(((top - 1)*layer.Width+right)/tileSize))] > 0)
	//BottomLeft Collision
	collisionMap = append(collisionMap, layer.Data[uint(math.Floor(((bottom)*layer.Width+left - 1)/tileSize))] > 0)
	//BottomRight Collision
	collisionMap = append(collisionMap, layer.Data[uint(math.Ceil(((bottom)*layer.Width+right)/tileSize))] > 0)

	return collisionMap
}