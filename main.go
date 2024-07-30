package main

import (
	"image/color"
	"log"

	"GoPlat/Components/animations"
	"GoPlat/Components/sprites"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	Player *sprites.Player
}

func (g *Game) Update() error {
	canIdle := true
	//handle movement
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.Player.IsMovingRight = true
		g.Player.X += 1
		g.Player.CurrentAnimationIndex = 0
		g.Player.IsIdle = false
		canIdle = false
	} else if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.Player.IsMovingRight = false
		g.Player.X -= 1
		g.Player.CurrentAnimationIndex = 0
		g.Player.IsIdle = false
		canIdle = false
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.Player.IsMovingRight = true
		g.Player.X += 1
		g.Player.CurrentAnimationIndex = 0
		g.Player.IsIdle = false
		canIdle = false
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.Player.IsMovingRight = false
		g.Player.X -= 1
		g.Player.CurrentAnimationIndex = 0
		g.Player.IsIdle = false
		canIdle = false
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.Player.CurrentAnimationIndex = 1
		g.Player.IsIdle = false
		canIdle = false
	}
	if canIdle {
		g.Player.IsIdle = true
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if len(g.Player.ActionAnimations) == 0 {
		return
	}
	screen.Fill(color.RGBA{150, 150, 255, 255})
	ebitenutil.DebugPrint(screen, "Hello! Welcome to GoPlat!")

	playerDrawOptions := ebiten.DrawImageOptions{}

	if g.Player.CurrentAnimationIndex == 1 {
		playerDrawOptions.GeoM.Translate(0, -16)
	}
	if !g.Player.IsMovingRight {
		playerDrawOptions.GeoM.Scale(-1, 1)

		var frameWidth float64
		if g.Player.IsIdle {
			frameWidth = g.Player.IdleAnimation.MaxFrameWidth
		} else {
			frameWidth = g.Player.ActionAnimations[g.Player.CurrentAnimationIndex].MaxFrameWidth
		}
		playerDrawOptions.GeoM.Translate(frameWidth, 0)
	}

	playerDrawOptions.GeoM.Translate(g.Player.X, g.Player.Y)
	var currentFrame *ebiten.Image
	if g.Player.IsIdle {
		currentFrame = g.Player.IdleAnimation.Animate()
	} else {
		currentFrame = g.Player.ActionAnimations[g.Player.CurrentAnimationIndex].Animate()
	}
	screen.DrawImage(currentFrame, &playerDrawOptions)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	playerWalk := animations.GeneratePlayerWalk()
	playerJump := animations.GeneratePlayerJump()
	playerIdle := animations.GeneratePlayerIdle()

	player := &sprites.Player{
		BioSprite: &sprites.BioSprite{
			Sprite: &sprites.Sprite{
				Image: playerWalk.Frames[0],
				X:     50,
				Y:     150,
			},
			ActionAnimations: []*animations.ActionAnimation{
				playerWalk,
				playerJump,
			},
			IsMovingRight:         true,
			IdleAnimation:         playerIdle,
			IsIdle:                true,
			CurrentAnimationIndex: 0,
		},
		HasSecondJump: true,
		IsWallSliding: false,
	}

	game := Game{
		Player: player,
	}

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
