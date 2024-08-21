package startup

import (
	levels "GoPlat/gameComponents/levels"
	"log"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func CreateLevels() []*levels.Level {
	levelOneMap, err := levels.NewTilemapScene("Assets/Maps/The_Descent_LevelOne.json")
	if err != nil {
		log.Fatal(err)
	}
	levelTwoMap, err := levels.NewTilemapScene("Assets/Maps/The_Descent_LevelTwo.JSON")
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
	levelTwo := &levels.Level{
		TilemapScene: levelTwoMap,
		TilemapImage: dungeonTiles,
	}

	levels := []*levels.Level{
		levelOne,
		levelTwo,
	}

	return levels
}
