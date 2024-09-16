package inventory

type Item struct {
	spawnName         string
	spawnCollected    bool
	spawnCollectSaved bool
	itemSuperType     string
	itemSubType       string
	itemText          string
	itemName          string
	quantity          int
}

func NewItem() *Item {
	return &Item{
		spawnName:         "genericItem",
		spawnCollected:    false,
		spawnCollectSaved: false,
		itemName:          "",
		itemSuperType:     "item",
		itemSubType:       "subItem",
		itemText:          "An Item",
		quantity:          0,
	}
}

func (i *Item) Collect() {
	i.spawnCollected = true
}

func (i *Item) UnCollect() {
	i.spawnCollected = false
}

func (i *Item) IsCollected() bool {
	return i.spawnCollected
}

func (i *Item) IsCollectSaved() bool {
	return i.spawnCollectSaved
}

func (i *Item) SaveCollect() {
	i.spawnCollectSaved = true
}

func (i *Item) UnSaveCollect() {
	i.spawnCollectSaved = false
}

func (i *Item) AddItem() {
	i.quantity++
}

func (i *Item) RemoveItem() {
	i.quantity--
}

func (i *Item) AddQuantity(qty int) {
	i.quantity += qty
}

func (i *Item) RemoveQuantity(qty int) {
	i.quantity -= qty
}

func (i *Item) GetQuantity() int {
	return i.quantity
}

func (i *Item) SetQuantity(qty int) {
	i.quantity = qty
}

func (i *Item) GetItemSuperType() string {
	return i.itemSuperType
}

func (i *Item) SetItemSuperType(superType string) {
	i.itemSuperType = superType
}

func (i *Item) GetItemSubType() string {
	return i.itemSubType
}

func (i *Item) SetItemSubType(subType string) {
	i.itemSubType = subType
}

func (i *Item) GetItemText() string {
	return i.itemText
}

func (i *Item) SetItemText(text string) {
	i.itemText = text
}

func (i *Item) GetSpawnName() string {
	return i.spawnName
}

func (i *Item) SetSpawnName(name string) {
	i.spawnName = name
}

func (i *Item) GetItemName() string {
	return i.itemName
}

func (i *Item) SetItemName(name string) {
	i.itemName = name
}
