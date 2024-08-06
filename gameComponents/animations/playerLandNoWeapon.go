package animations

import (
	"GoPlat/gameComponents/controls"
	"image"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func GeneratePlayerLandNoWeapon() *ActionAnimation {
	fullJumpPng, _, err := ebitenutil.NewImageFromFile("Assets/Images/KnightNoWeap/Jump.png")
	if err != nil {
		log.Fatal(err)
	}

	jumpPng17 := fullJumpPng.SubImage(
		image.Rect(1040, 16, 1072, 48),
	).(*ebiten.Image)

	jumpPng18 := fullJumpPng.SubImage(
		image.Rect(1104, 16, 1136, 48),
	).(*ebiten.Image)

	jumpPng19 := fullJumpPng.SubImage(
		image.Rect(1168, 16, 1200, 48),
	).(*ebiten.Image)

	frames := []*ebiten.Image{
		jumpPng17,
		jumpPng18,
		jumpPng19,
	}

	frameVectors := []controls.Vector{
		{0, 0},
		{0, 0},
		{0, 0},
	}

	cancelDirections := []controls.Direction{
		controls.JUMP,
		controls.DASHLEFT,
		controls.DASHRIGHT,
	}

	playerLand := &ActionAnimation{
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
		HasEffect: true,
		Effect: *GenerateEffectLand(),
	}

	return playerLand
}
