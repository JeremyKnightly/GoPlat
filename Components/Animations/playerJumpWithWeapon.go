package animations

import (
	controls "GoPlat/components/controls"
	"image"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func GeneratePlayerJumpWithWeapon() *ActionAnimation {
	fullJumpPng, _, err := ebitenutil.NewImageFromFile("Assets/Images/KnightWithWeap/Jump.png")
	if err != nil {
		log.Fatal(err)
	}

	jumpPngSquat := fullJumpPng.SubImage(
		image.Rect(80, 8, 112, 48),
	).(*ebiten.Image)

	jumpPngLeap := fullJumpPng.SubImage(
		image.Rect(128, 8, 176, 48),
	).(*ebiten.Image)

	jumpPngPinnacle := fullJumpPng.SubImage(
		image.Rect(520, 8, 560, 48),
	).(*ebiten.Image)

	jumpPngDown1 := fullJumpPng.SubImage(
		image.Rect(584, 8, 624, 48),
	).(*ebiten.Image)

	jumpPngDown2 := fullJumpPng.SubImage(
		image.Rect(656, 8, 688, 48),
	).(*ebiten.Image)

	jumpPngDown3 := fullJumpPng.SubImage(
		image.Rect(720, 8, 752, 48),
	).(*ebiten.Image)

	jumpPngDown4 := fullJumpPng.SubImage(
		image.Rect(784, 8, 816, 48),
	).(*ebiten.Image)

	jumpPngDown5 := fullJumpPng.SubImage(
		image.Rect(848, 8, 880, 48),
	).(*ebiten.Image)

	jumpPngDown6 := fullJumpPng.SubImage(
		image.Rect(904, 8, 944, 48),
	).(*ebiten.Image)

	jumpPngDown7 := fullJumpPng.SubImage(
		image.Rect(976, 8, 1008, 48),
	).(*ebiten.Image)

	jumpPngLand := fullJumpPng.SubImage(
		image.Rect(1040, 8, 1072, 48),
	).(*ebiten.Image)

	jumpPngFlourish := fullJumpPng.SubImage(
		image.Rect(1088, 8, 1136, 48),
	).(*ebiten.Image)

	jumpPngReset := fullJumpPng.SubImage(
		image.Rect(1160, 8, 1200, 48),
	).(*ebiten.Image)

	frames := []*ebiten.Image{
		jumpPngSquat,
		jumpPngLeap,
		jumpPngLeap,
		jumpPngLeap,
		jumpPngLeap,
		jumpPngLeap,
		jumpPngLeap,
		jumpPngPinnacle,
		jumpPngDown1,
		jumpPngDown2,
		jumpPngDown3,
		jumpPngDown4,
		jumpPngDown5,
		jumpPngDown6,
		jumpPngDown7,
		jumpPngLand,
		jumpPngFlourish,
		jumpPngReset,
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
		{.3, .25},
		{.3, .4},
		{.3, .4},
		{.3, .4},
		{.3, .4},
		{.3, .4},
		{.3, .4},
		{0, 0},
		{0, 0},
		{0, 0},
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
		AllowCancelAfterFrame:   6,
		AllowCancelOnDirections: cancelDirections,
		HasEffect: false,
	}

	return playerJump
}
