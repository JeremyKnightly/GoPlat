package sprites

import "time"

type Player struct {
	*BioSprite
	DashCooldowntime   time.Duration
	DashLastUsed       time.Time
	CanJump            bool
	IsPhysicsLocked    bool
	CanAnimationCancel bool
	IsAnimationLocked  bool
	IsAirborn          bool
	HasSecondJump      bool
	IsWallSliding      bool
	IsWallHanging      bool
}

func (p *Player) Kill() {
	println("dead")
	p.IsDead = true
	p.CurrentAnimationIndex = 8
	p.CanAnimationCancel = false
	p.IsWallHanging = false
	p.IsWallSliding = false
	p.IsPhysicsLocked = false
}

func (p *Player) Resurrect() {
	p.IsDead = false
	p.X = 50
	p.Y = 50
}
