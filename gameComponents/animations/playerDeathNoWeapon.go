package animations

import (
	"GoPlat/gameComponents/controls"
	"image"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func GeneratePlayerDeathNoWeapon() *ActionAnimation {
	fullPng, _, err := ebitenutil.NewImageFromFile("Assets/Images/KnightNoWeap/Death.png")
	if err != nil {
		log.Fatal(err)
	}
	png1 := fullPng.SubImage(
		image.Rect(16, 16, 48, 48),
	).(*ebiten.Image)

	png2 := fullPng.SubImage(
		image.Rect(80, 16, 112, 48),
	).(*ebiten.Image)

	png3 := fullPng.SubImage(
		image.Rect(144, 16, 176, 48),
	).(*ebiten.Image)

	png4 := fullPng.SubImage(
		image.Rect(208, 16, 240, 48),
	).(*ebiten.Image)

	png5 := fullPng.SubImage(
		image.Rect(272, 16, 304, 48),
	).(*ebiten.Image)

	png6 := fullPng.SubImage(
		image.Rect(336, 16, 368, 48),
	).(*ebiten.Image)

	png7 := fullPng.SubImage(
		image.Rect(400, 16, 432, 48),
	).(*ebiten.Image)

	png8 := fullPng.SubImage(
		image.Rect(464, 16, 496, 48),
	).(*ebiten.Image)

	png9 := fullPng.SubImage(
		image.Rect(464, 16, 496, 48),
	).(*ebiten.Image)

	frames := []*ebiten.Image{
		png1,
		png2,
		png3,
		png4,
		png5,
		png6,
		png7,
		png8,
		png9,
	}

	frameVectors := []controls.Vector{
		{
			DeltaX: 0,
			DeltaY: 0,
		},
		{
			DeltaX: 0,
			DeltaY: 0,
		},
		{
			DeltaX: 0,
			DeltaY: 0,
		},
		{
			DeltaX: 0,
			DeltaY: 0,
		},
		{
			DeltaX: 0,
			DeltaY: 0,
		},
		{
			DeltaX: 0,
			DeltaY: 0,
		},
		{
			DeltaX: 0,
			DeltaY: 0,
		},
		{
			DeltaX: 0,
			DeltaY: 0,
		},
		{
			DeltaX: 0,
			DeltaY: 0,
		},
	}

	cancelDirections := []controls.Direction{}

	death := &ActionAnimation{
		Animation: &Animation{
			Frames:            frames,
			NumberOfFrames:    uint16(len(frames)),
			CurrentFrameIndex: 0,
			FrameDuration:     time.Millisecond * 250,
			MaxFrameWidth:     float64(frames[0].Bounds().Dx()),
			MaxFrameHeight:    float64(frames[0].Bounds().Dy()),
		},
		AnimationComplete:       false,
		FrameVectors:            frameVectors,
		AllowCancelAfterFrame:   7,
		AllowCancelOnDirections: cancelDirections,
		HasEffect:               false,
		LoopAnimation:           false,
	}

	return death
}
