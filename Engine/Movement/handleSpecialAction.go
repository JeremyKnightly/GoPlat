package movement

import (
	"GoPlat/components/sprites"
	"time"
)

func HandleSpecialAction(p *sprites.Player, action string) bool {
	if action == "JUMP" {
		return handleJump(p)
	} else if action == "DASHLEFT" {
		return handleDashLeft(p)
	} else if action == "DASHRIGHT" {
		return handleDashRight(p)
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
		p.CurrentAnimationIndex = 1
		p.JumpLastUsed = time.Now()
		return true
	} else if p.HasSecondJump && p.IsAirborn{
		if p.IsAnimationLocked {
			p.ActionAnimations[p.CurrentAnimationIndex].ResetAnimation = true
		}
		p.CurrentAnimationIndex = 1
		p.HasSecondJump = false
		return true
	} 
	return false
}

func handleDashRight(p *sprites.Player) bool {
	if time.Since(p.DashLastUsed) >= p.DashCooldowntime {
		p.CurrentAnimationIndex = 2
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
		p.CurrentAnimationIndex = 2
		p.DashLastUsed = time.Now()
		if p.IsAnimationLocked {
			p.ActionAnimations[p.CurrentAnimationIndex].ResetAnimation = true
		}
		return true
	} 
	return false
}