package movement

import (
	"GoPlat/gameComponents/sprites"
	"time"
)

func HandleSpecialAction(p *sprites.Player, action string) bool {
	if action == "JUMP" {
		p.IsAirborn = false
		p.IsPhysicsLocked = true
		return handleJump(p)
	} else if action == "DASHLEFT" {
		p.IsAirborn = false
		p.IsPhysicsLocked = true
		return handleDashLeft(p)
	} else if action == "DASHRIGHT" {
		p.IsAirborn = false
		p.IsPhysicsLocked = true
		return handleDashRight(p)
	} else if action == "EDGECLIMB" {
		p.IsAirborn = false
		p.IsPhysicsLocked = true
	}
	return false
}

func handleJump(p *sprites.Player) bool {
	//checks if cooldown has reset or if player can double jump
	if time.Since(p.JumpLastUsed) >= p.JumpCooldownTime {
		if p.IsAnimationLocked {
			p.ActionAnimations[p.CurrentAnimationIndex].ResetAnimation = true
		}
		p.HasSecondJump = true
		p.CurrentAnimationIndex = 2
		p.JumpLastUsed = time.Now()
		return true
	} else if p.HasSecondJump && p.IsAirborn{
		if p.IsAnimationLocked {
			p.ActionAnimations[p.CurrentAnimationIndex].ResetAnimation = true
		}
		p.CurrentAnimationIndex = 3
		p.HasSecondJump = false
		return true
	} 
	return false
}

func handleDashRight(p *sprites.Player) bool {
	if time.Since(p.DashLastUsed) >= p.DashCooldowntime {
		p.CurrentAnimationIndex = 1
		p.DashLastUsed = time.Now()
		if p.IsAnimationLocked {
			p.ActionAnimations[p.CurrentAnimationIndex].ResetAnimation = true
		}
		return true
	}
	return false
}

func handleDashLeft(p *sprites.Player) bool {
	if time.Since(p.DashLastUsed) >= p.DashCooldowntime {
		p.IsMovingRight = false
		p.CurrentAnimationIndex = 1
		p.DashLastUsed = time.Now()
		if p.IsAnimationLocked {
			p.ActionAnimations[p.CurrentAnimationIndex].ResetAnimation = true
		}
		return true
	} 
	return false
}