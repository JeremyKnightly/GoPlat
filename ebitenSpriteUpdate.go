package main

import (
	"GoPlat/engine/collision"
	"GoPlat/gameComponents/PlayerStatus/inventory"
	"GoPlat/gameComponents/levels"
	"GoPlat/gameComponents/sprites"
	"errors"
)

func UpdateInteractableSprites(npsm *sprites.NPSpriteManager, lvl *levels.Level, inventory *inventory.Inventory) {
	for _, interactable := range npsm.ExistingInteractables {
		spawner, err := GetSpawner(lvl, interactable.GetSpawnName())
		if err != nil {
			println("spawner not found: ", interactable.GetSpawnName())
		}

		spawnCollected, showSpawn, goNextSprite := spawner.GetInteractStatus(inventory)
		spawner.SetPropValue("CanInteract", !spawnCollected)
		if !showSpawn {
			interactable.SetSpriteFrameImage(npsm.GetImageFromNamedSprite("Blank"))
			continue
		}
		if goNextSprite {
			interactable.GoToNextSprite()
		}
		interactable.SetSpriteFrameImage(npsm.GetImageFromNamedSprite(interactable.GetCurrentSpriteName()))
	}
}

func GetSpawner(lvl *levels.Level, spawnName string) (*collision.Rect, error) {
	colliders := collision.ExtractCollisionData(lvl)
	for _, collider := range colliders {
		if !collider.HasProp("UniqueName") {
			continue
		}
		if collider.GetPropValue("UniqueName").(string) == spawnName {
			return &collider, nil
		}
	}

	return nil, errors.New("Prop Not Found")
}

// -------------------initial loadup---------------------//
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

			var nextSpriteName string
			if temp.HasNextSprite {
				nextSpriteName = collision.GetPropValue("NextSpriteName").(string)
			}
			temp.SetNextSpriteName(nextSpriteName)
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
			var nextSpriteName string
			if temp.HasNextSprite {
				nextSpriteName = collision.GetPropValue("NextSpriteName").(string)
			}
			temp.SetNextSpriteName(nextSpriteName)
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
			var nextSpriteName string
			if temp.HasNextSprite {
				nextSpriteName = collision.GetPropValue("NextSpriteName").(string)
			}
			temp.SetNextSpriteName(nextSpriteName)
			temp.SetPosition(collision.X, collision.Y)

			if collision.HasProp("UniqueName") {
				temp.UniqueName = collision.GetPropValue("UniqueName").(string)
			}

			npsm.ExistingInteractables = append(npsm.ExistingInteractables, temp)
		}

	}
}
