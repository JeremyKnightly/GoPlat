package movement

import (
	"GoPlat/gameComponents/sprites"
)

func DoSpecialAction(p *sprites.Player, action string) bool {
	if action == "JUMP" {
		p.IsPhysicsLocked = true
		if p.IsWallHanging { //trigger edgeclimb
			p.IsWallHanging = false
			return handleWallHang(p)
		}
		valid := handleJump(p)
		if valid {
			p.CanDash = true
			p.IsAnimationLocked = true
		}
		return valid
	} else if action == "DASHLEFT" || action == "DASHRIGHT" {
		p.IsPhysicsLocked = true
		return handleDash(p, action == "DASHRIGHT")
	}
	return false
}

func handleJump(p *sprites.Player) bool {
	//checks if cooldown has reset or if player can double jump
	if p.CanJump {
		p.CurrentAnimationIndex = 2
		p.ActionAnimations[p.CurrentAnimationIndex].Reset()
		p.HasSecondJump = true
		p.CanJump = false
		return true
	} else if p.HasSecondJump && p.IsAirborn {
		p.CurrentAnimationIndex = 3
		p.ActionAnimations[p.CurrentAnimationIndex].Reset()
		p.HasSecondJump = false
		return true
	}
	return false
}

func handleDash(p *sprites.Player, movingRight bool) bool {
	if p.CanDash {
		p.CanDash = false
		p.IsMovingRight = movingRight
		p.CurrentAnimationIndex = 1
		p.ActionAnimations[p.CurrentAnimationIndex].Reset()
		p.IsAnimationLocked = true
		return true
	}
	return false
}

func handleWallHang(p *sprites.Player) bool {
	p.CurrentAnimationIndex = 4
	p.IsAnimationLocked = true
	return false
}
