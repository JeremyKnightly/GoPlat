package sprites

import (
	"GoPlat/gameComponents/levels"
	"time"
)

type Player struct {
	*BioSprite
	DashCooldowntime       time.Duration
	DashLastUsed           time.Time
	CanJump                bool
	IsPhysicsLocked        bool
	CanAnimationCancel     bool
	IsAnimationLocked      bool
	IsAirborn              bool
	HasSecondJump          bool
	IsWallSliding          bool
	IsWallHanging          bool
	CurrentCheckpointIndex int
}

func (p *Player) Kill() {
	p.IsDead = true
	p.CurrentAnimationIndex = 8
	p.CanAnimationCancel = false
	p.IsWallHanging = false
	p.IsWallSliding = false
	p.IsPhysicsLocked = false
}

func (p *Player) Resurrect(lvl *levels.Level) {
	p.IsDead = false
	checkpointX, checkpointY := lvl.GetCheckpointXY(p.CurrentCheckpointIndex)
	p.X = checkpointX
	p.Y = checkpointY
}

func (p *Player) SetNewCheckpoint(newIdx int) {
	p.CurrentCheckpointIndex = newIdx
}
