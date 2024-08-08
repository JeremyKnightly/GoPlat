package main

import (
	"GoPlat/engine/collision"
	movement "GoPlat/engine/movement"
	runtime "GoPlat/engine/processes/runtime"
	startup "GoPlat/engine/processes/startup"
	controls "GoPlat/gameComponents/controls"
	levels "GoPlat/gameComponents/levels"
	sprites "GoPlat/gameComponents/sprites"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Player   *sprites.Player
	levels   []*levels.Level
	controls []controls.Control
	currentLevel int
}

func (g *Game) Update() error {
	newVector := movement.HandleMovementCalculations(g.Player, g.controls, g.levels[g.currentLevel])
	
	validVector := collision.IsValidMove(g.levels[g.currentLevel], g.Player, newVector)
	if validVector {
		g.Player.X += newVector.DeltaX
		g.Player.Y += newVector.DeltaY
	}


	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	runtime.DrawLevel(g.levels[g.currentLevel], screen)
	runtime.DrawPlayer(g.Player, screen, g.levels[g.currentLevel])
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
	game.currentLevel = 0

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}