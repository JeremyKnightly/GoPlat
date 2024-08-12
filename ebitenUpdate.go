package main

import (
	"GoPlat/engine/collision"
	movement "GoPlat/engine/movement"
)

func (g *Game) Update() error {
	g.setPlayerPositionWithInput()
	g.setCameraPosition()
	g.currentLevel = g.levels[g.currentLevelIndex]

	return nil
}

func (g *Game) setPlayerPositionWithInput() {
	inputVector := movement.HandleMovementCalculations(g.Player, g.controls, g.currentLevel)

	validVector := collision.IsValidMove(g.currentLevel, g.Player, inputVector)
	if validVector {
		g.Player.X += inputVector.DeltaX
		g.Player.Y += inputVector.DeltaY
	}
}

func (g *Game) setCameraPosition() {
	g.camera.FollowTarget(g.Player.X+g.tileSize/2, g.Player.Y+g.tileSize/2, g.screenWidth, g.screenHeight)
	g.camera.Constrain(g.currentLevel.Layers[0].Height,
		g.currentLevel.Layers[0].Width,
		g.screenWidth,
		g.screenHeight,
	)
}
