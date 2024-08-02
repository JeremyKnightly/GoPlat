package main

import (
	controls "GoPlat/components/controls"
	levels "GoPlat/components/levels"
	sprites "GoPlat/components/sprites"
	movement "GoPlat/engine/movement"
	runtime "GoPlat/engine/processes/runtime"
	startup "GoPlat/engine/processes/startup"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Player   *sprites.Player
	levels   []*levels.Level
	controls []controls.Control
}

func (g *Game) Update() error {
	if g.Player.IsAnimationLocked {
		return nil
	}

	//handle movement
	directions := movement.GetControlsPressed(g.controls)
	playerVector, specialAction := movement.GetMovementVector(directions)

	g.Player.IsMovingRight = movement.IsMovingRight(playerVector)
	g.Player.IsIdle = movement.IsIdle(playerVector, specialAction)
	//Idle Detection

	if len(specialAction.Name) > 0 {
		if specialAction.Name == "JUMP" && time.Now().Sub(g.Player.JumpLastUsed) >= g.Player.JumpCooldownTime {
			g.Player.CurrentAnimationIndex = 1
			g.Player.JumpLastUsed = time.Now()
		} else if time.Now().Sub(g.Player.DashLastUsed) >= g.Player.DashCooldowntime &&
			specialAction.Name == "DASHLEFT" {
			g.Player.CurrentAnimationIndex = 2
			g.Player.DashLastUsed = time.Now()
		} else if time.Now().Sub(g.Player.DashLastUsed) >= g.Player.DashCooldowntime &&
			specialAction.Name == "DASHRIGHT" {
			g.Player.CurrentAnimationIndex = 2
			g.Player.DashLastUsed = time.Now()
		} else {
			return nil
		}
		g.Player.IsAnimationLocked = true
	} else {
		g.Player.CurrentAnimationIndex = 0
	}
	newX, newY := controls.AddVector(g.Player.X, g.Player.Y, playerVector)
	g.Player.X = newX
	g.Player.Y = newY

	return nil
}

// func (g *Game) Draw(screen *ebiten.Image) {
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
