package startup

import (
	levels "GoPlat/components/levels"
	"log"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func CreateLevels() ([]*levels.Level) {
	//testLevel, err := levels.NewTilemapScene("Assets/Maps/test.json")
	//if err != nil {
	//	log.Fatal(err)
	//}
	levelOneMap, err := levels.NewTilemapScene("Assets/Maps/LevelOne.json")
	if err != nil {
		log.Fatal(err)
	}
	dungeonTiles, _, err := ebitenutil.NewImageFromFile("Assets/Maps/Tilesets/Dungeon Tile Set.png")
	if err != nil {
		log.Fatal(err)
	}

	levelOne := &levels.Level{
		TilemapScene: levelOneMap,
		TilemapImage: dungeonTiles,
	}

	levels := []*levels.Level{
		levelOne,
	}

	return levels
}
