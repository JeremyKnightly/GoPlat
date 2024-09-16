package inventory

import score "GoPlat/gameComponents/PlayerStatus/Score"

type Inventory struct {
	Items []*Item
}

func NewInventory() *Inventory {
	return &Inventory{}
}

func (i *Inventory) ClearInventory() {
	for _, item := range i.Items {
		item.SetQuantity(0)
		if !item.IsCollectSaved() {
			item.UnCollect()
		}
	}
}

func (i *Inventory) MergeWInventory(inputInventory *Inventory, score *score.Score) {
	for _, item := range inputInventory.Items {
		if item.GetQuantity() == 0 { // skip if empty
			continue
		}

		inventoryItem, hasItem := i.GetItem(item.GetItemName())
		if hasItem { // if item exists, increment by amount in input
			inventoryItem.AddQuantity(item.GetQuantity())
			score.AddPoints(item.GetQuantity(), item.GetItemSuperType())
			item.SaveCollect()
			item.SetQuantity(0)
		} else { // else, copy item and add to inventory
			i.AddNewItem(*item)
			item.SaveCollect()
		}
	}
}

// Always use this when picking up item for spawner tracking
func (i *Inventory) AddNewItem(item Item) {
	item.Collect()
	i.Items = append(i.Items, &item)
}

func (i *Inventory) RemoveItem(itemName string) bool {
	item, hasItem := i.GetItem(itemName)
	if !hasItem || (hasItem && item.GetQuantity() == 0) {
		return false
	}

	item.RemoveItem()
	return true
}

func (i *Inventory) HasItem(itemName string) bool {
	for _, item := range i.Items {
		if item.GetItemName() == itemName {
			return true
		}
	}
	return false
}

func (i *Inventory) GetItem(itemName string) (*Item, bool) {
	if !i.HasItem(itemName) {
		return nil, false
	}

	for _, item := range i.Items {
		if item.GetItemName() == itemName {
			return item, true
		}
	}

	return nil, false
}

func (i *Inventory) IsSpawnCollected(spawnName string) bool {
	for _, item := range i.Items {
		if item.GetSpawnName() == spawnName {
			return item.IsCollected()
		}
	}
	return false
}
