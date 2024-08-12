package collision

import (
	"GoPlat/gameComponents/levels"
)

func ExtractCollisionData(level *levels.Level) []Rect {
	var collisionRects []Rect

	for _, layer := range level.ObjectLayers {
		for _, rect := range layer.Objects {
			//properties := []levels.Property{}
			/*for _, prop := range rect.Properties{
				properties = append(properties, prop)
			}*/
			newRect := Rect{
				X:          rect.X,
				Y:          rect.Y,
				Width:      rect.Width,
				Height:     rect.Height,
				Properties: rect.Properties,
			}
			collisionRects = append(collisionRects, newRect)
		}
	}

	return collisionRects
}
