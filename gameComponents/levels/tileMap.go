package levels

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

// full map struct for separating tilemaps and Obj layers
type Map struct {
	Height int               `json:"height"`
	Layers []json.RawMessage `json:"layers"`
	Type   string            `json:"type"`
	Width  int               `json:"width"`
}

type Level struct {
	*TilemapScene
	TilemapImage *ebiten.Image
}

type TilemapLayer struct {
	Data       []int   `json:"data"`
	Width      float64 `json: "width"`
	Height     float64 `json:"height"`
	Visible    bool    `json:"visible"`
	FirstDraw  bool
	Properties []Property `json:"properties"`
}

type ObjectLayer struct {
	Objects []Object `json:"objects"`
}

type Object struct {
	X          float64    `json:"x"`
	Y          float64    `json:"y"`
	Width      float64    `json:"width"`
	Height     float64    `json:"height"`
	Properties []Property `json:"properties"`
}

type Property struct {
	Name  string      `json:"name"`
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

type TilemapScene struct {
	Layers       []TilemapLayer `json:"layers"`
	ObjectLayers []ObjectLayer  `json:"layers"`
	Checkpoints  []*Checkpoint
}

func NewTilemapScene(filepath string) (*TilemapScene, error) {
	var tileMapReturn TilemapScene

	jsonMapData, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var fullMap Map
	err = json.Unmarshal([]byte(jsonMapData), &fullMap)
	if err != nil {
		return nil, err
	}

	for _, layerData := range fullMap.Layers {
		var layerType map[string]interface{}
		err := json.Unmarshal(layerData, &layerType)
		if err != nil {
			return nil, err
		}

		layerTypeName := layerType["type"].(string)
		//layerProperties := layerType["properties"].(PropertySlice)
		switch layerTypeName {
		default:
			fmt.Printf("Layer type %s unknown\n", layerTypeName)
		case "tilelayer":
			var tempLayer TilemapLayer
			err = json.Unmarshal(layerData, &tempLayer)
			if err != nil {
				return nil, err
			}
			for _, prop := range tempLayer.Properties {
				if prop.Name == "FirstDraw" {
					switch value := prop.Value.(type) {
					case bool:
						tempLayer.FirstDraw = value
					default:
					}
				}
			}
			tileMapReturn.Layers = append(tileMapReturn.Layers, tempLayer)
		case "objectgroup":
			var tempLayer ObjectLayer
			err = json.Unmarshal(layerData, &tempLayer)
			if err != nil {
				return nil, err
			}

			tempLayer.extractCheckpoints(&tileMapReturn)

			tileMapReturn.ObjectLayers = append(tileMapReturn.ObjectLayers, tempLayer)

		}
	}
	return &tileMapReturn, nil
}

func (ol *ObjectLayer) extractCheckpoints(tilemapRtn *TilemapScene) {
	for _, obj := range ol.Objects {
		checkpointFound := false
		isEnd := false
		isStart := false
		var checkpointIdx int
		for _, prop := range obj.Properties {
			if prop.Name == "Checkpoint" {
				checkpointFound = true
			}

			if prop.Name == "StartPoint" {
				isStart = true
			} else if prop.Name == "EndPoint" {
				isEnd = true
			} else if prop.Name == "CheckPointIndex" {
				checkpointIdx = int(prop.Value.(float64))
			}
		}

		if checkpointFound {
			//if found, append to checkpoints, then delete from obj and restart check
			tilemapRtn.Checkpoints = append(tilemapRtn.Checkpoints,
				&Checkpoint{
					X:          obj.X,
					Y:          obj.Y,
					Index:      checkpointIdx,
					StartPoint: isStart,
					EndPoint:   isEnd,
				},
			)
		}
	}
}
