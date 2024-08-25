package sprites

import (
	"GoPlat/gameComponents/levels"
)

type Player struct {
	*BioSprite
	CanJump                bool
	IsPhysicsLocked        bool
	CanAnimationCancel     bool
	IsAnimationLocked      bool
	IsAirborn              bool
	HasSecondJump          bool
	CanDash                bool
	IsWallSliding          bool
	IsWallHanging          bool
	CurrentCheckpointIndex int
	Deaths                 int
}

func (p *Player) Kill() {
	p.IsDead = true
	p.CurrentAnimationIndex = 8
	p.CanAnimationCancel = false
	p.IsWallHanging = false
	p.IsWallSliding = false
	p.IsPhysicsLocked = false
	p.CanDash = true
}

func (p *Player) Resurrect(lvl *levels.Level) {
	p.Deaths++
	p.IsDead = false
	checkpointX, checkpointY := lvl.GetCheckpointXY(p.CurrentCheckpointIndex)
	p.X = checkpointX
	p.Y = checkpointY
}

func (p *Player) SetNewCheckpoint(newIdx int) {
	p.CurrentCheckpointIndex = newIdx
}
