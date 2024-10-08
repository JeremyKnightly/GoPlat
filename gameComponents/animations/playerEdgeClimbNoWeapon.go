package animations

import (
	"GoPlat/gameComponents/controls"
	"image"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func GeneratePlayerEdgeClimbNoWeapon() *ActionAnimation {
	fullPng, _, err := ebitenutil.NewImageFromFile("Assets/Images/KnightNoWeap/Edge-Climb.png")
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

	frames := []*ebiten.Image{
		png1,
		png2,
		png3,
		png4,
		png5,
		png6,
	}

	frameVectors := []controls.Vector{
		{
			DeltaX: 0,
			DeltaY: -4,
		},
		{
			DeltaX: 0,
			DeltaY: -3,
		},
		{
			DeltaX: 1.2,
			DeltaY: -2.4,
		},
		{
			DeltaX: 2.0,
			DeltaY: -2.4,
		},
		{
			DeltaX: 1.45,
			DeltaY: -2.1,
		},
		{
			DeltaX: 1.05,
			DeltaY: -.4,
		},
	}

	cancelDirections := []controls.Direction{
		controls.DASHLEFT,
		controls.DASHRIGHT,
	}

	edgeClimb := &ActionAnimation{
		Animation: &Animation{
			Frames:            frames,
			NumberOfFrames:    uint16(len(frames)),
			CurrentFrameIndex: 0,
			FrameDuration:     time.Millisecond * 80,
			MaxFrameWidth:     float64(frames[0].Bounds().Dx()),
			MaxFrameHeight:    float64(frames[0].Bounds().Dy()),
		},
		AnimationComplete:       false,
		FrameVectors:            frameVectors,
		AllowCancelAfterFrame:   2,
		AllowCancelOnDirections: cancelDirections,
		HasEffect:               false,
	}

	return edgeClimb
}
