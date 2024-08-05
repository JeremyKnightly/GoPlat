package main

import (
	controls "GoPlat/components/controls"
	levels "GoPlat/components/levels"
	sprites "GoPlat/components/sprites"
	"GoPlat/engine/collision"
	movement "GoPlat/engine/movement"
	runtime "GoPlat/engine/processes/runtime"
	startup "GoPlat/engine/processes/startup"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Player   *sprites.Player
	levels   []*levels.Level
	controls []controls.Control
	currentLevel int
	maxTilesX int
	maxTilesY int
	tileSize float64
}

func (g *Game) Update() error {
	newVector := movement.HandleMovementCalculations(g.Player, g.controls)
	
	validVector := collision.IsValidMove(g.levels[g.currentLevel], g.Player, newVector)
	if validVector {
		g.Player.X += newVector.DeltaX
		g.Player.Y += newVector.DeltaY
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	runtime.DrawLevel(g.levels[0], screen)
	runtime.DrawPlayer(g.Player, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 340, 270
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("GoPlat!")

	var game Game
	game.Player = startup.CreateDefaultPlayer()
	game.levels = startup.CreateLevels()
	game.controls = startup.GetControls()
	game.maxTilesX = 20
	game.maxTilesY = 15
	game.tileSize = 16
	game.currentLevel = 0

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}




