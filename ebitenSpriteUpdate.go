package main

import (
	"GoPlat/engine/collision"
	"GoPlat/gameComponents/sprites"
)

func (g *Game) LoadSpriteSpawns() {
	npsm := g.NPSpriteManager
	collisionData := collision.ExtractCollisionData(g.currentLevel)

	for _, collision := range collisionData {
		if !collision.HasSpecialProps() {
			continue
		}
		if !collision.HasProp("SpriteSpawn") || !collision.HasProp("CanInteract") {
			continue
		}

		switch collision.GetPropValue("SpawnType").(string) {
		case "Item":
			temp := &sprites.ItemSprite{
				HasNextSprite:   collision.GetPropValue("HasNextSprite").(bool),
				FirstSpriteName: collision.GetPropValue("SpriteName").(string),
			}

			temp.X = collision.X
			temp.Y = collision.Y

			if collision.HasProp("UniqueName") {
				temp.UniqueName = collision.GetPropValue("UniqueName").(string)
			}

			npsm.ExistingInteractables = append(npsm.ExistingInteractables, temp)

		case "Key":
			temp := &sprites.KeySprite{
				HasNextSprite:   collision.GetPropValue("HasNextSprite").(bool),
				FirstSpriteName: collision.GetPropValue("SpriteName").(string),
			}

			temp.X = collision.X
			temp.Y = collision.Y

			if collision.HasProp("UniqueName") {
				temp.UniqueName = collision.GetPropValue("UniqueName").(string)
			}

			npsm.ExistingInteractables = append(npsm.ExistingInteractables, temp)

		case "Power":
			temp := &sprites.PUpSprite{
				HasNextSprite:   collision.GetPropValue("HasNextSprite").(bool),
				FirstSpriteName: collision.GetPropValue("SpriteName").(string),
			}

			temp.X = collision.X
			temp.Y = collision.Y

			if collision.HasProp("UniqueName") {
				temp.UniqueName = collision.GetPropValue("UniqueName").(string)
			}

			npsm.ExistingInteractables = append(npsm.ExistingInteractables, temp)
		}

	}
}
