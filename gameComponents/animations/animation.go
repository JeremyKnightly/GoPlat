package animations

import (
	controls "GoPlat/gameComponents/controls"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Animation struct {
	Frames            []*ebiten.Image
	NumberOfFrames    uint16
	CurrentFrameIndex uint16
	frameDuration     time.Duration
	lastUpdate        time.Time
	MaxFrameWidth, MaxFrameHeight     float64
}

type ActionAnimation struct {
	*Animation
	AnimationComplete bool
	WillAwaitInput    bool
	stopAnimation     bool
	FrameVectors []controls.Vector
	AllowCancelAfterFrame uint16
	LoopAnimation bool
	AllowCancelOnDirections []controls.Direction
	ResetAnimation bool
	HasEffect bool
	Effect Effect
}

type Effect struct {
	*Animation
	Offset float64
}

func (a *Animation) Animate() *ebiten.Image {
	now := time.Now()

	if now.Sub(a.lastUpdate) >= a.frameDuration {
		a.CurrentFrameIndex++

		if a.CurrentFrameIndex >= a.NumberOfFrames - 1 {
			a.CurrentFrameIndex = 0
		}
		a.lastUpdate = now
	}

	return a.Frames[a.CurrentFrameIndex]
}

func (a *ActionAnimation) AnimateAction() (*ebiten.Image,controls.Vector, bool) {
	now := time.Now()

	if now.Sub(a.lastUpdate) >= a.frameDuration {
		if !a.stopAnimation {
			a.CurrentFrameIndex++
		}

		if a.CurrentFrameIndex >= a.NumberOfFrames - 1 {
			a.AnimationComplete = true
			if a.WillAwaitInput {
				a.CurrentFrameIndex = a.NumberOfFrames - 1
				a.stopAnimation = true
			} else {
				a.CurrentFrameIndex = 0
			}
		}
		if a.ResetAnimation {
			a.CurrentFrameIndex = 0
			a.ResetAnimation = false
		}
		a.lastUpdate = now
	}	
	allowCancel := false
	if a.CurrentFrameIndex > a.AllowCancelAfterFrame{
		allowCancel = true
	}

	return a.Frames[a.CurrentFrameIndex], a.FrameVectors[a.CurrentFrameIndex], allowCancel
}
