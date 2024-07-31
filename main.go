package main

import (
	controls "GoPlat/Components/Controls"
	levels "GoPlat/Components/Levels"
	sprites "GoPlat/Components/Sprites"
	startup "GoPlat/Processes/Startup"
	runtime "GoPlat/Processes/runtime"
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Player       *sprites.Player
	levels []*levels.Level
	controls []controls.Control
}

func (g *Game) Update() error {
	shouldIdle := true


	//handle movement
	for idx,_ := range g.controls {
		controlActivated := true
		if len(g.controls[idx].Keys) >0 {
			for _,key := range g.controls[idx].Keys {
				if !ebiten.IsKeyPressed(key){
					controlActivated = false
				}
			}
		}else if ebiten.IsKeyPressed(g.controls[idx].Key){
			controlActivated = true
		} else {
			controlActivated = false
		}

		if controlActivated {
			fmt.Println("moving ",g.controls[idx].GetDirection())
		}
	}




	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.Player.IsMovingRight = true
		g.Player.X += 1
		g.Player.CurrentAnimationIndex = 0
		g.Player.IsIdle = false
		shouldIdle = false
	} else if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.Player.IsMovingRight = false
		g.Player.X -= 1
		g.Player.CurrentAnimationIndex = 0
		g.Player.IsIdle = false
		shouldIdle = false
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.Player.IsMovingRight = true
		g.Player.X += 1
		g.Player.CurrentAnimationIndex = 0
		g.Player.IsIdle = false
		shouldIdle = false
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.Player.IsMovingRight = false
		g.Player.X -= 1
		g.Player.CurrentAnimationIndex = 0
		g.Player.IsIdle = false
		shouldIdle = false
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.Player.CurrentAnimationIndex = 1
		g.Player.IsIdle = false
		shouldIdle = false
		g.Player.IsAnimationLocked = true
	}
	if shouldIdle {
		g.Player.IsIdle = true
	}

	return nil
}

//func (g *Game) Draw(screen *ebiten.Image) {
func (g *Game) Draw(screen *ebiten.Image) {
	runtime.DrawLevel(g.levels[0], screen)
	runtime.DrawPlayer(g.Player, screen)	
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("GoPlat!")

	var game Game
	game.Player = startup.CreateDefaultPlayer()
	game.levels = startup.CreateLevels()
	game.controls = startup.GetControls()

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
