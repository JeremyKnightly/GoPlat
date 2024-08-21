package movement

import (
	"GoPlat/gameComponents/sprites"
	"time"
)

func HandleSpecialAction(p *sprites.Player, action string) bool {
	if action == "JUMP" {
		p.IsPhysicsLocked = true
		if p.IsWallHanging { //trigger edgeclimb
			p.IsWallHanging = false
			return handleWallHang(p)
		}
		valid := handleJump(p)
		if valid {
			p.CanDash = true
		}
		return valid
	} else if action == "DASHLEFT" {
		p.IsPhysicsLocked = true
		return handleDashLeft(p)
	} else if action == "DASHRIGHT" {
		p.IsPhysicsLocked = true
		return handleDashRight(p)
	}
	return false
}

func handleJump(p *sprites.Player) bool {
	//checks if cooldown has reset or if player can double jump
	if p.CanJump {
		if p.IsAnimationLocked {
			p.ActionAnimations[p.CurrentAnimationIndex].ResetAnimation = true
		}
		p.HasSecondJump = true
		p.CanJump = false
		p.CurrentAnimationIndex = 2
		return true
	} else if p.HasSecondJump && p.IsAirborn {
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
	if time.Since(p.DashLastUsed) >= p.DashCooldowntime && p.CanDash {
		p.CanDash = false
		p.IsMovingRight = true
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
	if time.Since(p.DashLastUsed) >= p.DashCooldowntime && p.CanDash {
		p.CanDash = false
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

func handleWallHang(p *sprites.Player) bool {
	p.CurrentAnimationIndex = 4
	p.IsAnimationLocked = true
	return false
}
