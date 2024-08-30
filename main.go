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
	SoundManager                                                                 *sound.SoundManager
	Player                                                                       *sprites.Player
	levels                                                                       []*levels.Level
	startScreen                                                                  *levels.Level
	controls                                                                     []controls.Control
	currentLevel                                                                 *levels.Level
	currentBGMIdx                                                                int
	currentLevelIndex                                                            int
	camera                                                                       *camera.Camera
	gameState                                                                    int
	defaultWindowWidth, defaultWindowHeight, startScreenWidth, startScreenHeight int
	tileSize, defaultScreenWidth, defaultScreenHeight                            float64
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	var screenWidthIn float64
	var screenHeightIn float64
	if g.gameState == 0 {
		screenWidthIn = float64(g.startScreenWidth)
		screenHeightIn = float64(g.startScreenHeight)
	} else if g.gameState == 1 {
		screenWidthIn = g.defaultScreenWidth
		screenHeightIn = g.defaultScreenHeight
	}

	return int(screenWidthIn), int(screenHeightIn)
}

func main() {
	var game Game
	game.SetGameProperties()

	ebiten.SetWindowSize(game.defaultWindowWidth, game.defaultWindowHeight)
	ebiten.SetWindowTitle("The Descent")

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}

func (game *Game) SetGameProperties() {
	game.levels = startup.CreateLevels()
	game.startScreen = startup.CreateStartScreen()
	game.controls = startup.GetControls()
	game.currentLevelIndex = 0
	game.tileSize = 16
	game.defaultScreenWidth = 380
	game.defaultScreenHeight = 285
	game.startScreenWidth = 800
	game.startScreenHeight = 630
	game.currentBGMIdx = 0
	game.defaultWindowWidth = 1200
	game.defaultWindowHeight = 950
	game.gameState = 0
	game.camera = camera.NewCamera(0, 0)
	game.Player = startup.CreateDefaultPlayer(game.levels[game.currentLevelIndex])
	game.SoundManager = startup.GetAllSounds()
}
