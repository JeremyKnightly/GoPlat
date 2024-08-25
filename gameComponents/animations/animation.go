package animations

import (
	controls "GoPlat/gameComponents/controls"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Animation struct {
	Frames                        []*ebiten.Image
	NumberOfFrames                uint16
	CurrentFrameIndex             uint16
	FrameDuration                 time.Duration
	TicksPerFrame                 float64
	TicksThisFrame                float64
	lastUpdate                    time.Time
	MaxFrameWidth, MaxFrameHeight float64
}

type ActionAnimation struct {
	*Animation
	AnimationComplete       bool
	WillAwaitInput          bool
	StopAnimation           bool
	FrameVectors            []controls.Vector
	AllowCancelAfterFrame   uint16
	LoopAnimation           bool
	AllowCancelOnDirections []controls.Direction
	ResetAnimation          bool
	HasEffect               bool
	Effect                  Effect
}

type Effect struct {
	*Animation
	OffsetX         float64
	OffsetOneWay    bool
	OffsetRightOnly bool
}

func (a *Animation) Animate() *ebiten.Image {
	now := time.Now()
	a.TicksThisFrame++

	if now.Sub(a.lastUpdate) >= a.FrameDuration {
		a.CurrentFrameIndex++
		a.TicksThisFrame = 1

		if a.CurrentFrameIndex >= a.NumberOfFrames-1 {
			a.CurrentFrameIndex = 0
		}
		a.lastUpdate = now
	}

	return a.Frames[a.CurrentFrameIndex]
}

func (a *ActionAnimation) AnimateAction() (*ebiten.Image, controls.Vector, bool, bool) {
	now := time.Now()
	a.TicksThisFrame++

	changedIndex := false
	if now.Sub(a.lastUpdate) >= a.FrameDuration {
		if !a.StopAnimation {
			a.CurrentFrameIndex++
			a.TicksThisFrame = 1
			changedIndex = true
		}

		if a.CurrentFrameIndex > a.NumberOfFrames-1 {
			a.endAnimationCycle()
		}
		if a.ResetAnimation {
			a.Reset()
		}
		a.lastUpdate = now
	}
	allowCancel := a.setAnimationCancel()
	returnVector := a.FrameVectors[a.CurrentFrameIndex].ScaleByTPS(a.TicksThisFrame, a.TicksPerFrame)

	return a.Frames[a.CurrentFrameIndex], returnVector, allowCancel, changedIndex
}

func (a *ActionAnimation) setAnimationCancel() bool {
	return a.CurrentFrameIndex >= a.AllowCancelAfterFrame
}

func (a *ActionAnimation) Reset() {
	a.CurrentFrameIndex = 0
	a.ResetAnimation = false
}

func (a *ActionAnimation) endAnimationCycle() {
	if a.WillAwaitInput {
		a.CurrentFrameIndex = a.NumberOfFrames - 1
		a.StopAnimation = true
	} else {
		a.AnimationComplete = true
		a.ResetAnimation = true
	}
}
