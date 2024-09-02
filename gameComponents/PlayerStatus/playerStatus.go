package playerstatus

import (
	score "GoPlat/gameComponents/PlayerStatus/Score"
	"GoPlat/gameComponents/PlayerStatus/inventory"
	"GoPlat/gameComponents/PlayerStatus/powers"
)

type Status struct {
	Score     *score.Score
	Powers    *powers.Powers
	Inventory *inventory.Inventory
}

func (s *Status) AddKey(keyName string) {
	//add key to inventory

}
