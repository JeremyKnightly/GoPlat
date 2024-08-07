package animations

import (
	controls "GoPlat/gameComponents/controls"
	"image"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func GeneratePlayerWalkNoWeapon() *ActionAnimation {
	fullWalkPng, _, err := ebitenutil.NewImageFromFile("Assets/Images/KnightNoWeap/Walk.png")
	if err != nil {
		log.Fatal(err)
	}

	walkPng1 := fullWalkPng.SubImage(
		image.Rect(16, 16, 48, 48),
	).(*ebiten.Image)

	walkPng2 := fullWalkPng.SubImage(
		image.Rect(112, 16, 80, 48),
	).(*ebiten.Image)

	walkPng3 := fullWalkPng.SubImage(
		image.Rect(176, 16, 144, 48),
	).(*ebiten.Image)

	walkPng4 := fullWalkPng.SubImage(
		image.Rect(208, 16, 240, 48),
	).(*ebiten.Image)

	walkPng5 := fullWalkPng.SubImage(
		image.Rect(304, 16, 272, 48),
	).(*ebiten.Image)

	walkPng6 := fullWalkPng.SubImage(
		image.Rect(368, 16, 336, 48),
	).(*ebiten.Image)

	walkPng7 := fullWalkPng.SubImage(
		image.Rect(432, 16, 400, 48),
	).(*ebiten.Image)

	walkPng8 := fullWalkPng.SubImage(
		image.Rect(496, 16, 464, 48),
	).(*ebiten.Image)


	frames := []*ebiten.Image{
		walkPng1,
		walkPng2,
		walkPng3,
		walkPng4,
		walkPng5,
		walkPng6,
		walkPng7,
		walkPng8,
	}

	frameVectors := []controls.Vector{
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
	}

	playerWalk := &ActionAnimation{
		Animation: &Animation{
			Frames:            frames,
			NumberOfFrames:    uint16(len(frames)),
			CurrentFrameIndex: 0,
			frameDuration:     time.Millisecond * 100,
			MaxFrameWidth:     float64(frames[0].Bounds().Dx()),
			MaxFrameHeight: float64(frames[0].Bounds().Dy()),
		},
		AnimationComplete: false,
		FrameVectors: frameVectors,
		HasEffect: false,
		LoopAnimation: true,
	}

	return playerWalk
}
