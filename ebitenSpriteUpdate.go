package main

import (
	"GoPlat/engine/collision"
	"GoPlat/gameComponents/sprites"
)

func (g *Game) UpdateSpriteSpawns() {
	npsm := g.NPSpriteManager
	collisionData := collision.ExtractCollisionData(g.currentLevel)

	for _, collision := range collisionData {
		if !collision.HasSpecialProps() {
			continue
		}
		if !collision.HasProp("SpriteSpawn") {
			continue
		}

		switch collision.GetPropValue("SpawnType").(string) {
		case "Item":
			temp := &sprites.ItemSprite{
				Sprite: &sprites.Sprite{
					Frame: sprites.GetNewFrame(),
					X:     0,
					Y:     0,
				},
				HasNextSprite:     collision.GetPropValue("HasNextSprite").(bool),
				FirstSpriteName:   collision.GetPropValue("SpriteName").(string),
				CurrentSpriteName: collision.GetPropValue("SpriteName").(string),
			}

			temp.SetPosition(collision.X, collision.Y)

			if collision.HasProp("UniqueName") {
				temp.UniqueName = collision.GetPropValue("UniqueName").(string)
			}

			npsm.ExistingInteractables = append(npsm.ExistingInteractables, temp)

		case "Key":
			temp := &sprites.KeySprite{
				Sprite: &sprites.Sprite{
					Frame: sprites.GetNewFrame(),
					X:     0,
					Y:     0,
				},
				HasNextSprite:     collision.GetPropValue("HasNextSprite").(bool),
				FirstSpriteName:   collision.GetPropValue("SpriteName").(string),
				CurrentSpriteName: collision.GetPropValue("SpriteName").(string),
			}
			temp.SetPosition(collision.X, collision.Y)

			if collision.HasProp("UniqueName") {
				temp.UniqueName = collision.GetPropValue("UniqueName").(string)
			}

			npsm.ExistingInteractables = append(npsm.ExistingInteractables, temp)

		case "Power":
			temp := &sprites.PUpSprite{
				Sprite: &sprites.Sprite{
					Frame: sprites.GetNewFrame(),
					X:     0,
					Y:     0,
				},
				HasNextSprite:     collision.GetPropValue("HasNextSprite").(bool),
				FirstSpriteName:   collision.GetPropValue("SpriteName").(string),
				CurrentSpriteName: collision.GetPropValue("SpriteName").(string),
			}
			temp.SetPosition(collision.X, collision.Y)

			if collision.HasProp("UniqueName") {
				temp.UniqueName = collision.GetPropValue("UniqueName").(string)
			}

			npsm.ExistingInteractables = append(npsm.ExistingInteractables, temp)
		}

	}
}
