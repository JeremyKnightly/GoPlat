package main

import (
	runtime "GoPlat/engine/processes/runtime"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Draw(screen *ebiten.Image) {
	runtime.DrawLevelFirstDraw(g.levels[g.currentLevel], screen, g.camera)
	runtime.DrawPlayer(g.Player, screen, g.levels[g.currentLevel], g.camera)
	runtime.DrawLevelSecondDraw(g.levels[g.currentLevel], screen, g.camera)
}
