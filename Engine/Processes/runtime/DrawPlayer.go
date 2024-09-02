package runtime

import (
	sprites "GoPlat/gameComponents/sprites"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func DrawPlayer(player *sprites.Player, screen *ebiten.Image) {
	screen.DrawImage(player.Frame.ImageToDraw, &player.Frame.ImageOptions)
	if player.Frame.HasEffect {
		screen.DrawImage(player.Frame.EffectImageToDraw, &player.Frame.EffectOptions)
	}
}

func DrawPlayerDeathCount(numDeaths int, screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Deaths: %v", numDeaths))
}

func DrawScore(p *sprites.Player, screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Score: %v", p.Status.Score.GetScore()))
}
