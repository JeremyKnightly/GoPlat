package animations

import (
	"image"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func GenerateEffectJump() *Effect {
	fullJumpPng, _, err := ebitenutil.NewImageFromFile("Assets/Images/Effects/Jump_effect.png")
	if err != nil {
		log.Fatal(err)
	}

	png1 := fullJumpPng.SubImage(
		image.Rect(16, 16, 48, 48),
	).(*ebiten.Image)

	png2 := fullJumpPng.SubImage(
		image.Rect(80, 16, 112, 48),
	).(*ebiten.Image)

	png3 := fullJumpPng.SubImage(
		image.Rect(144, 16, 176, 48),
	).(*ebiten.Image)

	blank := fullJumpPng.SubImage(
		image.Rect(16, 16, 20, 20),
	).(*ebiten.Image)

	frames := []*ebiten.Image{
		png1,
		png2,
		png3,
		blank,
		blank,
		blank,
		blank,
		blank,
	}

	effect := &Effect{
				Animation: &Animation{
			Frames:            frames,
			NumberOfFrames:    uint16(len(frames)),
			CurrentFrameIndex: 0,
			frameDuration:     time.Millisecond * 60,
			MaxFrameWidth:     float64(frames[0].Bounds().Dx()),
			MaxFrameHeight: float64(frames[0].Bounds().Dy()),
		},
		Offset: 0,
	}


	return effect
}
