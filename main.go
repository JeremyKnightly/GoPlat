package main

import (
	controls "GoPlat/components/controls"
	levels "GoPlat/components/levels"
	sprites "GoPlat/components/sprites"
	movement "GoPlat/engine/movement"
	runtime "GoPlat/engine/processes/runtime"
	startup "GoPlat/engine/processes/startup"
	"fmt"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Player   *sprites.Player
	levels   []*levels.Level
	controls []controls.Control
	maxTilesX int
	maxTilesY int
	tileSize float64
}

func (g *Game) Update() error {
	movement.HandleMovementCalculations(g.Player, g.controls)
	_ = g.collisionTest()
	return nil
}

// func (g *Game) Draw(screen *ebiten.Image) {
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

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}


func (g *Game)collisionTest () bool {

	left, top, right, bottom := getPlayerCorners(g.Player, g.tileSize, g.levels[0].TilemapScene.Layers[1])

	collisionMap := checkCollision(g.levels[0].TilemapScene.Layers[1], left, top, right, bottom)

	fmt.Printf("\n\n%v %v\n%v %v",collisionMap[0],collisionMap[1],collisionMap[2],collisionMap[3])


	return true
}


func getPlayerCorners(p *sprites.Player, tileSize float64, layer levels.TilemapLayer) (float64, float64, float64, float64){
	left := math.Ceil(p.X/tileSize)
	top := math.Ceil(p.Y/tileSize)
	right := 0.0
	bottom := 0.0

	if p.IsIdle{
		right = math.Floor((p.X + p.IdleAnimation.MaxFrameWidth) / tileSize)
		bottom = math.Floor((p.Y + p.IdleAnimation.MaxFrameHeight) / tileSize)
	} else {
		right = math.Floor((p.X + p.ActionAnimations[p.CurrentAnimationIndex].MaxFrameWidth) / tileSize)
		bottom = math.Floor((p.Y + p.ActionAnimations[p.CurrentAnimationIndex].MaxFrameHeight) / tileSize)
	}

	return left, top, right, bottom
}

func checkCollision(layer levels.TilemapLayer, left float64, top float64, right float64, bottom float64) []bool {
	collisionMap := []bool{}

	//TopLeft Collision
	collisionMap = append(collisionMap, layer.Data[uint((top)*layer.Width + left)]>0)
	//TopRight Collision
	collisionMap = append(collisionMap, layer.Data[uint((top)*layer.Width + right)]>0)
	//BottomLeft Collision
	collisionMap = append(collisionMap, layer.Data[uint((bottom)*layer.Width + left)]>0)	
	//BottomRight Collision
	collisionMap = append(collisionMap, layer.Data[uint((bottom)*layer.Width + right)]>0)

	return collisionMap
}