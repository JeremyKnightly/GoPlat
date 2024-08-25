package main

import (
	camera "GoPlat/Engine/camera"
	startup "GoPlat/engine/processes/startup"
	controls "GoPlat/gameComponents/controls"
	levels "GoPlat/gameComponents/levels"
	sound "GoPlat/gameComponents/sounds"
	sprites "GoPlat/gameComponents/sprites"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	SoundManager                        *sound.SoundManager
	Player                              *sprites.Player
	levels                              []*levels.Level
	controls                            []controls.Control
	currentLevel                        *levels.Level
	currentBGMIdx                       int
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
<<<<<<< HEAD
	game.SetGameProperties()
=======
	game.Player = startup.CreateDefaultPlayer()
	game.levels = startup.CreateLevels()
	game.controls = startup.GetControls()
	game.currentLevelIndex = 0
	game.tileSize = 16
	game.screenHeight = 240
	game.screenWidth = 320
	game.camera = camera.NewCamera(0, 0)
>>>>>>> 16f3c53bd428513a1c986c0fe9c23443b6469d9c

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
	game.currentBGMIdx = 0
	game.camera = camera.NewCamera(0, 0)
	game.Player = startup.CreateDefaultPlayer(game.levels[game.currentLevelIndex])
	game.SoundManager = startup.GetAllSounds()
}
