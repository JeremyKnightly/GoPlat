package main

import (
	"GoPlat/engine/camera"
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
	camera *camera.Camera
	tileSize, screenWidth, screenHeight float64
}

func (g *Game) Update() error {
	newVector := movement.HandleMovementCalculations(g.Player, g.controls, g.levels[g.currentLevel])
	
	validVector := collision.IsValidMove(g.levels[g.currentLevel], g.Player, newVector)
	if validVector {
		g.Player.X += newVector.DeltaX
		g.Player.Y += newVector.DeltaY
	}

	g.camera.FollowTarget(g.Player.X + g.tileSize/2, g.Player.Y + g.tileSize/2, g.screenWidth, g.screenHeight)
	g.camera.Constrain(g.levels[g.currentLevel].Layers[0].Height,
		g.levels[g.currentLevel].Layers[0].Width,
		g.screenWidth, 
		g.screenHeight,
	)
	
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	runtime.DrawLevel(g.levels[g.currentLevel], screen, g.camera)
	runtime.DrawPlayer(g.Player, screen, g.levels[g.currentLevel], g.camera)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int(g.screenWidth), int(g.screenHeight)
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("GoPlat!")

	var game Game
	game.Player = startup.CreateDefaultPlayer()
	game.levels = startup.CreateLevels()
	game.controls = startup.GetControls()
	game.currentLevel = 0
	game.tileSize = 16
	game.screenHeight = 240
	game.screenWidth = 320
	game.camera = camera.NewCamera(0, 0)

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}