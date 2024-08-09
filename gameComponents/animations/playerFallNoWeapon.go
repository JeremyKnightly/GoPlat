package animations

import (
	"GoPlat/gameComponents/controls"
	"image"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func GeneratePlayerFallNoWeapon() *ActionAnimation {
	fullJumpPng, _, err := ebitenutil.NewImageFromFile("Assets/Images/KnightNoWeap/Jump.png")
	if err != nil {
		log.Fatal(err)
	}

	jumpPng12 := fullJumpPng.SubImage(
		image.Rect(720, 16, 752, 48),
	).(*ebiten.Image)

	jumpPng13 := fullJumpPng.SubImage(
		image.Rect(784, 16, 816, 48),
	).(*ebiten.Image)

	jumpPng14 := fullJumpPng.SubImage(
		image.Rect(848, 16, 880, 48),
	).(*ebiten.Image)

	jumpPng15 := fullJumpPng.SubImage(
		image.Rect(912, 16, 944, 48),
	).(*ebiten.Image)

	jumpPng16 := fullJumpPng.SubImage(
		image.Rect(976, 16, 1008, 48),
	).(*ebiten.Image)

	frames := []*ebiten.Image{
		jumpPng12,
		jumpPng13,
		jumpPng14,
		jumpPng15,
		jumpPng16,
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
	}

	cancelDirections := []controls.Direction{
		controls.JUMP,
		controls.DASHLEFT,
		controls.DASHRIGHT,
	}

	playerFall := &ActionAnimation{
		Animation: &Animation{
			Frames:            frames,
			NumberOfFrames:    uint16(len(frames)),
			CurrentFrameIndex: 0,
			frameDuration:     time.Millisecond * 60,
			MaxFrameWidth:     float64(frames[0].Bounds().Dx()),
			MaxFrameHeight: float64(frames[0].Bounds().Dy()),
		},
		AnimationComplete:       false,
		FrameVectors:            frameVectors,
		AllowCancelAfterFrame:   0,
		AllowCancelOnDirections: cancelDirections,
		HasEffect: false,
		LoopAnimation: true,
	}

	return playerFall
}
