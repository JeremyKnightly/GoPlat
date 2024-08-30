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
	levelThreeMap, err := levels.NewTilemapScene("Assets/Maps/LevelTwo.JSON")
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
	levelThree := &levels.Level{
		TilemapScene: levelThreeMap,
		TilemapImage: dungeonTiles,
	}

	levels := []*levels.Level{
		levelOne,
		levelTwo,
		levelThree,
	}

	return levels
}

func CreateStartScreen() *levels.Level {
	startScreenRaw, err := levels.NewTilemapScene("Assets/Maps/TitleScreen.JSON")
	if err != nil {
		log.Fatal(err)
	}
	dungeonTiles, _, err := ebitenutil.NewImageFromFile("Assets/Maps/Tilesets/Dungeon Tile Set.png")
	if err != nil {
		log.Fatal(err)
	}

	startScreen := &levels.Level{
		TilemapScene: startScreenRaw,
		TilemapImage: dungeonTiles,
	}

	return startScreen
}
