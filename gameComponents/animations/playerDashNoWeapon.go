package animations

import (
	controls "GoPlat/gameComponents/controls"
	"image"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func GeneratePlayerDashNoWeapon() *ActionAnimation {
	fullRunPng, _, err := ebitenutil.NewImageFromFile("Assets/Images/KnightNoWeap/Run.png")
	if err != nil {
		log.Fatal(err)
	}
	runPng1 := fullRunPng.SubImage(
		image.Rect(16, 16, 48, 48),
	).(*ebiten.Image)

	runPng2 := fullRunPng.SubImage(
		image.Rect(80, 16, 112, 48),
	).(*ebiten.Image)

	runPng3 := fullRunPng.SubImage(
		image.Rect(144, 16, 176, 48),
	).(*ebiten.Image)

	runPng4 := fullRunPng.SubImage(
		image.Rect(208, 16, 240, 48),
	).(*ebiten.Image)

	runPng5 := fullRunPng.SubImage(
		image.Rect(272, 16, 304, 48),
	).(*ebiten.Image)

	runPng6 := fullRunPng.SubImage(
		image.Rect(336, 16, 368, 48),
	).(*ebiten.Image)

	runPng7 := fullRunPng.SubImage(
		image.Rect(400, 16, 432, 48),
	).(*ebiten.Image)

	runPng8 := fullRunPng.SubImage(
		image.Rect(464, 16, 496, 48),
	).(*ebiten.Image)

	frames := []*ebiten.Image{
		runPng1,
		runPng2,
		runPng2,
		runPng2,
		runPng2,
		runPng3,
		runPng4,
		runPng5,
		runPng6,
		runPng7,
		runPng8,
	}

	frameVectors := []controls.Vector{
		{
			DeltaX: 1.5,
			DeltaY: 0,
		},
		{
			DeltaX: 1.75,
			DeltaY: 0,
		},
		{
			DeltaX: 1.75,
			DeltaY: 0,
		},
		{
			DeltaX: 1.75,
			DeltaY: 0,
		},
		{
			DeltaX: 1.75,
			DeltaY: 0,
		},
		{
			DeltaX: 1.75,
			DeltaY: 0,
		},
		{
			DeltaX: 1.5,
			DeltaY: 0,
		},
		{
			DeltaX: 1.25,
			DeltaY: 0,
		},
		{
			DeltaX: 1,
			DeltaY: 0,
		},
		{
			DeltaX: .75,
			DeltaY: 0,
		}, {
			DeltaX: .65,
			DeltaY: 0,
		},
	}

	cancelDirections := []controls.Direction{
		controls.JUMP,
		controls.DASHLEFT,
		controls.DASHRIGHT,
	}

	playerRun := &ActionAnimation{
		Animation: &Animation{
			Frames:            frames,
			NumberOfFrames:    uint16(len(frames)),
			CurrentFrameIndex: 0,
			FrameDuration:     time.Millisecond * 50,
			MaxFrameWidth:     float64(frames[0].Bounds().Dx()),
			MaxFrameHeight:    float64(frames[0].Bounds().Dy()),
		},
		AnimationComplete:       false,
		FrameVectors:            frameVectors,
		AllowCancelAfterFrame:   2,
		AllowCancelOnDirections: cancelDirections,
		HasEffect:               true,
		Effect:                  *GenerateEffectDash(),
	}

	return playerRun
}
