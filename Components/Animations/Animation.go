package animations

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Animation struct {
	Frames            []*ebiten.Image
	NumberOfFrames    uint16
	CurrentFrameIndex uint16
	frameDuration     time.Duration
	lastUpdate        time.Time
	MaxFrameWidth     float64
}

type ActionAnimation struct {
	*Animation
	AnimationComplete bool
	WillAwaitInput    bool
	stopAnimation     bool
}

func (a *Animation) Animate() *ebiten.Image {
	now := time.Now()

	if now.Sub(a.lastUpdate) >= a.frameDuration {
		a.CurrentFrameIndex++

		if a.CurrentFrameIndex >= a.NumberOfFrames {
			a.CurrentFrameIndex = 0
		}
		a.lastUpdate = now
	}

	return a.Frames[a.CurrentFrameIndex]
}

func (a *ActionAnimation) Animate() *ebiten.Image {
	now := time.Now()

	if now.Sub(a.lastUpdate) >= a.frameDuration {
		if !a.stopAnimation {
			a.CurrentFrameIndex++
		}

		if a.CurrentFrameIndex >= a.NumberOfFrames {
			a.AnimationComplete = true
			if a.WillAwaitInput {
				a.CurrentFrameIndex = a.NumberOfFrames - 1
				a.stopAnimation = true
			} else {
				a.CurrentFrameIndex = 0
			}
		}
		a.lastUpdate = now
	}

	return a.Frames[a.CurrentFrameIndex]
}
