package collision

import (
	"GoPlat/gameComponents/levels"
	"GoPlat/gameComponents/sprites"
)

type Rect struct {
	X, Y, Width, Height float64
	Properties          []levels.Property
}

func (r *Rect) HasSpecialProps() bool {
	for _, property := range r.Properties {
		if property.Name == "Special" {
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

func (r *Rect) HandleProps(p *sprites.Player) {
	for _, property := range r.Properties {
		if property.Name == "KillPlayer" && property.Value == true {
			p.Kill()
		}
	}
}
