package main

import (
	runtime "GoPlat/engine/processes/runtime"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Draw(screen *ebiten.Image) {
	if g.gameState == 0 { //start screen
		g.DrawStartScreen(screen)
	} else if g.gameState == 1 { //game
		g.DrawLevel(screen)
	} else if g.gameState == 2 { //credits

	} else if g.gameState == 3 { //High Scores

	}

}

func (g *Game) DrawLevel(screen *ebiten.Image) {
	runtime.DrawLevelFirstDraw(g.currentLevel, screen, g.camera)
	runtime.DrawPlayer(g.Player, screen)
	runtime.DrawLevelSecondDraw(g.currentLevel, screen, g.camera)
	runtime.DrawPlayerDeathCount(g.Player.Deaths, screen)
}

func (g *Game) DrawStartScreen(screen *ebiten.Image) {
	runtime.DrawLevelFirstDraw(g.currentLevel, screen, g.camera)
	runtime.DrawPlayer(g.Player, screen)
	//runtime.DrawLevelSecondDraw(g.currentLevel, screen, g.camera)
}
