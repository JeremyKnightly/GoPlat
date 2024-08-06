package animations

import (
	"GoPlat/gameComponents/controls"
	"image"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func GeneratePlayerJumpNoWeapon() *ActionAnimation {
	fullJumpPng, _, err := ebitenutil.NewImageFromFile("Assets/Images/KnightNoWeap/Jump.png")
	if err != nil {
		log.Fatal(err)
	}
	jumpPng1 := fullJumpPng.SubImage(
		image.Rect(16, 16, 48, 48),
	).(*ebiten.Image)

	jumpPng2 := fullJumpPng.SubImage(
		image.Rect(80, 16, 112, 48),
	).(*ebiten.Image)

	jumpPng3 := fullJumpPng.SubImage(
		image.Rect(144, 16, 176, 48),
	).(*ebiten.Image)

	jumpPng4 := fullJumpPng.SubImage(
		image.Rect(208, 16, 240, 48),
	).(*ebiten.Image)

	jumpPng5 := fullJumpPng.SubImage(
		image.Rect(272, 16, 304, 48),
	).(*ebiten.Image)

	jumpPng6 := fullJumpPng.SubImage(
		image.Rect(336, 16, 368, 48),
	).(*ebiten.Image)

	jumpPng7 := fullJumpPng.SubImage(
		image.Rect(400, 16, 432, 48),
	).(*ebiten.Image)

	jumpPng8 := fullJumpPng.SubImage(
		image.Rect(464, 16, 496, 48),
	).(*ebiten.Image)

	jumpPng9 := fullJumpPng.SubImage(
		image.Rect(528, 16, 560, 48),
	).(*ebiten.Image)

	frames := []*ebiten.Image{
		jumpPng1,
		jumpPng2,
		jumpPng3,
		jumpPng4,
		jumpPng5,
		jumpPng6,
		jumpPng7,
		jumpPng8,
		jumpPng9,
	}

	frameVectors := []controls.Vector{
		{0, 0},
		{0, 0},
		{.3, -.4},
		{.3, -.4},
		{.3, -.4},
		{.3, -.4},
		{.3, -.4},
		{.3, -.4},
		{.3, -.25},
	}

	cancelDirections := []controls.Direction{
		controls.JUMP,
		controls.DASHLEFT,
		controls.DASHRIGHT,
	}

	playerJump := &ActionAnimation{
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
		AllowCancelAfterFrame:   5,
		AllowCancelOnDirections: cancelDirections,
		HasEffect: true,
		Effect: *GenerateEffectJump(),
	}

	return playerJump
}
