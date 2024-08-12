package animations

import (
	"image"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func GenerateEffectDash() *Effect {
	fullPng, _, err := ebitenutil.NewImageFromFile("Assets/Images/Effects/Enemy_Attack.png")
	if err != nil {
		log.Fatal(err)
	}
	dashBlank := fullPng.SubImage(
		image.Rect(24, 16, 24, 48),
	).(*ebiten.Image)

	dash1 := fullPng.SubImage(
		image.Rect(0, 16, 32, 48),
	).(*ebiten.Image)

	dash2 := fullPng.SubImage(
		image.Rect(64, 16, 112, 48),
	).(*ebiten.Image)

	dash3 := fullPng.SubImage(
		image.Rect(128, 16, 176, 48),
	).(*ebiten.Image)

	dash4 := fullPng.SubImage(
		image.Rect(192, 16, 240, 48),
	).(*ebiten.Image)

	dash5 := fullPng.SubImage(
		image.Rect(256, 16, 304, 48),
	).(*ebiten.Image)

	dash6 := fullPng.SubImage(
		image.Rect(320, 16, 368, 48),
	).(*ebiten.Image)

	frames := []*ebiten.Image{
		dashBlank,
		dash1,
		dash2,
		dash3,
		dashBlank,
		dashBlank,
		dashBlank,
		dashBlank,
		dash4,
		dash5,
		dash6,
	}

	effect :=
		&Effect{
			Animation: &Animation{
				Frames:            frames,
				NumberOfFrames:    uint16(len(frames)),
				CurrentFrameIndex: 0,
				FrameDuration:     time.Millisecond * 80,
				MaxFrameWidth:     float64(frames[0].Bounds().Dx()),
				MaxFrameHeight:    float64(frames[0].Bounds().Dy()),
			},
			Offset: -16,
		}

	return effect
}
