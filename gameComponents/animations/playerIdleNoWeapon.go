package animations

import (
	"image"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func GeneratePlayerIdleNoWeapon() *Animation {
	fullIdlePng, _, err := ebitenutil.NewImageFromFile("Assets/Images/KnightNoWeap/Idle.png")
	if err != nil {
		log.Fatal(err)
	}

	idlePng1 := fullIdlePng.SubImage(
		image.Rect(8, 16, 48, 48),
	).(*ebiten.Image)

	idlePng2 := fullIdlePng.SubImage(
		image.Rect(72, 16, 112, 48),
	).(*ebiten.Image)

	idlePng3 := fullIdlePng.SubImage(
		image.Rect(136, 16, 176, 48),
	).(*ebiten.Image)

	idlePng4 := fullIdlePng.SubImage(
		image.Rect(200, 16, 240, 48),
	).(*ebiten.Image)

	idlePng5 := fullIdlePng.SubImage(
		image.Rect(264, 16, 304, 48),
	).(*ebiten.Image)

	idlePng6 := fullIdlePng.SubImage(
		image.Rect(328, 16, 368, 48),
	).(*ebiten.Image)

	idlePng7 := fullIdlePng.SubImage(
		image.Rect(392, 16, 432, 48),
	).(*ebiten.Image)

	idlePng8 := fullIdlePng.SubImage(
		image.Rect(456, 16, 496, 48),
	).(*ebiten.Image)

	idlePng9 := fullIdlePng.SubImage(
		image.Rect(520, 16, 560, 48),
	).(*ebiten.Image)

	idlePng10 := fullIdlePng.SubImage(
		image.Rect(584, 16, 624, 48),
	).(*ebiten.Image)

	idlePng11 := fullIdlePng.SubImage(
		image.Rect(648, 16, 688, 48),
	).(*ebiten.Image)

	idlePng12 := fullIdlePng.SubImage(
		image.Rect(712, 16, 752, 48),
	).(*ebiten.Image)

	idlePng13 := fullIdlePng.SubImage(
		image.Rect(776, 16, 816, 48),
	).(*ebiten.Image)

	idlePng14 := fullIdlePng.SubImage(
		image.Rect(840, 16, 880, 48),
	).(*ebiten.Image)

	frames := []*ebiten.Image{
		idlePng1,
		idlePng2,
		idlePng3,
		idlePng4,
		idlePng5,
		idlePng6,
		idlePng7,
		idlePng8,
		idlePng9,
		idlePng10,
		idlePng11,
		idlePng12,
		idlePng13,
		idlePng14,
	}

	playerIdle := &Animation{
		Frames:            frames,
		NumberOfFrames:    uint16(len(frames)),
		CurrentFrameIndex: 0,
		FrameDuration:     time.Millisecond * 200,
		MaxFrameWidth:     float64(frames[0].Bounds().Dx()),
		MaxFrameHeight:    float64(frames[0].Bounds().Dy()),
	}

	return playerIdle
}
