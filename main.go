package main

import (
	"image"
	"log"

	"GoPlat/Components/animations"
	"GoPlat/Components/sprites"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	Player       *sprites.Player
	tilemapJSON  *TilemapJSON
	tilemapImage *ebiten.Image
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

	mapDrawOptions := ebiten.DrawImageOptions{}
	//loop over layers
	for _, layer := range g.tilemapJSON.Layers {
		for index, id := range layer.Data {
			x := index % layer.Width
			y := index / layer.Width

			x *= 16
			y *= 16

			srcX := (id - 1) % 15
			srcY := (id - 1) / 15

			srcX *= 16
			srcY *= 16

			mapDrawOptions.GeoM.Translate(float64(x), float64(y))

			screen.DrawImage(g.tilemapImage.SubImage(image.Rect(srcX, srcY, srcX+16, srcY+16)).(*ebiten.Image),
				&mapDrawOptions,
			)

			mapDrawOptions.GeoM.Reset()
		}
	}

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
	ebiten.SetWindowTitle("GoPlat!")

	tilemapJSON, err := newTilemapJSON("Assets/Maps/test.json")
	if err != nil {
		log.Fatal(err)
	}
	tilemapImage, _, err := ebitenutil.NewImageFromFile("Assets/Maps/Tilesets/Dungeon Tile Set.png")

	if err != nil {
		log.Fatal(err)
	}
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
		Player:       player,
		tilemapJSON:  tilemapJSON,
		tilemapImage: tilemapImage,
	}

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
