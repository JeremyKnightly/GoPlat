package levels

import (
	"encoding/json"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type Level struct {
	*TilemapScene
	TilemapImage *ebiten.Image
}

type TilemapLayer struct {
	Data   []int `json:"data"`
	Width  float64   `json: "width"`
	Height float64   `json:"height"`
}

type TilemapScene struct {
	Layers []TilemapLayer `json:"layers"`
}

func NewTilemapScene(filepath string) (*TilemapScene, error) {
	contents, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var tilemap TilemapScene
	err = json.Unmarshal(contents, &tilemap)
	if err != nil {
		return nil, err
	}

	return &tilemap, nil
}
