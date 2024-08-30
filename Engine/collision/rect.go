package collision

import (
	"GoPlat/gameComponents/levels"
	"GoPlat/gameComponents/sprites"
)

type Rect struct {
	X, Y, Width, Height float64
	Properties          []levels.Property
}

func NewBlankRect(x, y float64) Rect {
	return Rect{
		X:      x,
		Y:      y,
		Width:  10,
		Height: 10,
	}
}

func (r *Rect) HasSpecialProps() bool {
	for _, property := range r.Properties {
		if property.Name == "Special" || property.Name == "Checkpoint" {
			return true
		}
	}

	return false
}

func (r *Rect) HasProp(propName string) bool {
	for _, property := range r.Properties {
		if property.Name == propName {
			return true
		}
	}

	return false
}

func (r *Rect) GetPropValue(propName string) interface{} {
	for _, property := range r.Properties {
		if property.Name == propName {
			return property.Value
		}
	}

	return false
}

func (r *Rect) HandleProps(p *sprites.Player) {
	isCheckpoint := false
	checkpointIdx := 0
	for _, property := range r.Properties {
		if property.Name == "KillPlayer" && property.Value == true {
			p.Kill()
		}
		if property.Name == "Checkpoint" {
			isCheckpoint = true
		}
		if property.Name == "CheckPointIndex" {
			checkpointIdx = int(property.Value.(float64))
		}
	}

	if isCheckpoint {
		p.SetNewCheckpoint(checkpointIdx)
	}
}
