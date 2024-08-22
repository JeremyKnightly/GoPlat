package main

import (
	runtime "GoPlat/engine/processes/runtime"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Draw(screen *ebiten.Image) {
	runtime.DrawLevelFirstDraw(g.currentLevel, screen, g.camera)
	runtime.DrawPlayer(g.Player, screen)
	runtime.DrawLevelSecondDraw(g.currentLevel, screen, g.camera)
	runtime.DrawPlayerDeathCount(g.Player.Deaths, screen)
}
