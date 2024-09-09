package inventory

type Item struct {
	uniqueItemName string
	itemSuperType  string
	itemSubType    string
	messageText    string
	quantity       int
}

func NewItem() *Item {
	return &Item{
		uniqueItemName: "genericItem",
		itemSuperType:  "item",
		messageText:    "An Item",
		quantity:       0,
	}
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

func (i *Item) GetMessageText() string {
	return i.messageText
}

func (i *Item) SetMessageText(text string) {
	i.messageText = text
}

func (i *Item) GetName() string {
	return i.uniqueItemName
}

func (i *Item) SetName(name string) {
	i.uniqueItemName = name
}
