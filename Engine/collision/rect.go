package collision

import (
	"GoPlat/gameComponents/PlayerStatus/inventory"
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

func (r *Rect) SetPropValue(propName string, value interface{}) {
	for _, property := range r.Properties {
		if property.Name == propName {
			property.Value = value
			return
		}
	}
	println("value not found")
}

func (r *Rect) HandleProps(p *sprites.Player) {
	if r.HasProp("KillPlayer") {
		p.Kill()
	}
	isCheckpoint := r.HasProp("Checkpoint")
	if isCheckpoint {
		p.SetNewCheckpoint(int(r.GetPropValue("CheckPointIndex").(float64)))
	}

	if r.HasProp("CanInteract") && r.GetPropValue("CanInteract").(bool) {
		item := *r.GetSpawnable()
		r.SetPropValue("CanInteract", false)
		p.Status.TempInventory.AddNewItem(item)
		//need func to repopulate sprites inventory clear
	}

}

func (r *Rect) GetSpawnable() *inventory.Item {
	item := inventory.NewItem()
	item.SetSpawnName(r.GetPropValue("UniqueName").(string))
	item.Collect()
	item.SetItemSuperType(r.GetPropValue("SpawnType").(string))
	item.SetItemSubType(r.GetPropValue("SpawnSubType").(string))
	item.SetItemText(r.GetPropValue("ItemText").(string))
	item.SetItemName(r.GetPropValue("ItemName").(string))

	return item
}

// returns if spawn has been collected and if sprite should show
// if sprite should show but has been collected, last is if sprite should advance an image
func (r *Rect) GetInteractStatus(inventory *inventory.Inventory) (bool, bool, bool) {
	spawnName := r.GetPropValue("UniqueName").(string)
	collected := inventory.IsSpawnCollected(spawnName)
	showSpawn := !r.GetPropValue("DespawnOnPickup").(bool)

	return collected, showSpawn, collected && showSpawn
}
