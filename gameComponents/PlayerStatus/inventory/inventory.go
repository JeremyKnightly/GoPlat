package inventory

type Inventory struct {
	Items []*Item
}

func (i *Inventory) ClearInventory() {
	for _, item := range i.Items {
		item.SetQuantity(0)
	}
}

func (i *Inventory) MergeWInventory(inputInventory *Inventory) {
	for _, item := range inputInventory.Items {
		if item.GetQuantity() == 0 { // skip if empty
			continue
		}

		inventoryItem, hasItem := i.GetItem(item.GetName())
		if hasItem { // if item exists, increment by amount in input
			inventoryItem.AddQuantity(item.GetQuantity())
			item.SetQuantity(0)
		} else { // else, copy item and add to inventory
			i.AddNewItem(*item)
		}
	}
}

func (i *Inventory) AddNewItem(item Item) {
	i.Items = append(i.Items, &item)
}

func (i *Inventory) AddItem(itemName string) bool {
	item, hasItem := i.GetItem(itemName)
	if !hasItem {
		return false
	}

	item.AddItem()
	return true
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
		if item.GetName() == itemName {
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
		if item.GetName() == itemName {
			return item, true
		}
	}

	return nil, false
}
