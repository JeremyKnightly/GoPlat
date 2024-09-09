package status

import (
	score "GoPlat/gameComponents/PlayerStatus/Score"
	"GoPlat/gameComponents/PlayerStatus/inventory"
	"GoPlat/gameComponents/PlayerStatus/powers"
)

type Status struct {
	Score         *score.Score
	Powers        *powers.Powers
	Inventory     *inventory.Inventory
	TempInventory *inventory.Inventory
}

func (s *Status) AddKey(keyName string) {
	//add key to inventory

}

func CreateNewStatus() *Status {
	return &Status{
		Score:         &score.Score{},
		Inventory:     inventory.NewInventory(),
		TempInventory: inventory.NewInventory(),
		Powers:        &powers.Powers{},
	}
}
