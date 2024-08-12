package animations

import (
	"GoPlat/gameComponents/controls"
	"image"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func GeneratePlayerDoubleJumpNoWeapon() *ActionAnimation {
	fullDblJumpPng, _, err := ebitenutil.NewImageFromFile("Assets/Images/KnightNoWeap/Double-Jump.png")
	if err != nil {
		log.Fatal(err)
	}
	dblJumpPng1 := fullDblJumpPng.SubImage(
		image.Rect(16, 24, 48, 48),
	).(*ebiten.Image)

	dblJumpPng2 := fullDblJumpPng.SubImage(
		image.Rect(80, 24, 112, 48),
	).(*ebiten.Image)

	dblJumpPng3 := fullDblJumpPng.SubImage(
		image.Rect(144, 24, 176, 48),
	).(*ebiten.Image)

	dblJumpPng4 := fullDblJumpPng.SubImage(
		image.Rect(208, 24, 240, 48),
	).(*ebiten.Image)

	dblJumpPng5 := fullDblJumpPng.SubImage(
		image.Rect(272, 24, 304, 48),
	).(*ebiten.Image)

	dblJumpPng6 := fullDblJumpPng.SubImage(
		image.Rect(336, 24, 368, 48),
	).(*ebiten.Image)

	dblJumpPng7 := fullDblJumpPng.SubImage(
		image.Rect(400, 24, 432, 48),
	).(*ebiten.Image)

	dblJumpPng8 := fullDblJumpPng.SubImage(
		image.Rect(464, 24, 496, 48),
	).(*ebiten.Image)

	frames := []*ebiten.Image{
		dblJumpPng1,
		dblJumpPng2,
		dblJumpPng3,
		dblJumpPng4,
		dblJumpPng5,
		dblJumpPng6,
		dblJumpPng7,
		dblJumpPng8,
	}

	frameVectors := []controls.Vector{
		{
			DeltaX: .1,
			DeltaY: -.6,
		},
		{
			DeltaX: .1,
			DeltaY: -.6,
		},
		{
			DeltaX: .1,
			DeltaY: -.6,
		},
		{
			DeltaX: .1,
			DeltaY: -.6,
		},
		{
			DeltaX: .1,
			DeltaY: -.6,
		},
		{
			DeltaX: .1,
			DeltaY: -.6,
		},
		{
			DeltaX: .1,
			DeltaY: -.6,
		},
		{
			DeltaX: .1,
			DeltaY: -.1,
		},
	}

	cancelDirections := []controls.Direction{
		controls.DASHLEFT,
		controls.DASHRIGHT,
	}

	playerDblJump := &ActionAnimation{
		Animation: &Animation{
			Frames:            frames,
			NumberOfFrames:    uint16(len(frames)),
			CurrentFrameIndex: 0,
			FrameDuration:     time.Millisecond * 60,
			MaxFrameWidth:     float64(frames[0].Bounds().Dx()),
			MaxFrameHeight:    float64(frames[0].Bounds().Dy()),
		},
		AnimationComplete:       false,
		FrameVectors:            frameVectors,
		AllowCancelAfterFrame:   3,
		AllowCancelOnDirections: cancelDirections,
		HasEffect:               true,
		Effect:                  *GenerateEffectDblJump(),
	}

	return playerDblJump
}
