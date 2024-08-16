package main

import (
	camera "GoPlat/Engine/camera"
	startup "GoPlat/engine/processes/startup"
	controls "GoPlat/gameComponents/controls"
	levels "GoPlat/gameComponents/levels"
	sprites "GoPlat/gameComponents/sprites"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Player                              *sprites.Player
	levels                              []*levels.Level
	controls                            []controls.Control
	currentLevel                        *levels.Level
	currentLevelIndex                   int
	camera                              *camera.Camera
	tileSize, screenWidth, screenHeight float64
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int(g.screenWidth), int(g.screenHeight)
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("GoPlat!")

	var game Game
	game.SetGameProperties()

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}

func (game *Game) SetGameProperties() {
	game.levels = startup.CreateLevels()
	game.controls = startup.GetControls()
	game.currentLevelIndex = 0
	game.tileSize = 16
	game.screenHeight = 210
	game.screenWidth = 280
	game.camera = camera.NewCamera(0, 0)
	game.Player = startup.CreateDefaultPlayer(game.levels[game.currentLevelIndex])
}
