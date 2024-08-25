package movement

import (
	"GoPlat/gameComponents/controls"
	"GoPlat/gameComponents/sprites"
)

func HandleSpecialAction(p *sprites.Player, action string) bool {
	if action == "JUMP" {
<<<<<<< HEAD
		p.IsPhysicsLocked = true
=======
>>>>>>> 16f3c53bd428513a1c986c0fe9c23443b6469d9c
		if p.IsWallHanging { //trigger edgeclimb
			p.IsWallHanging = false
			return handleWallHang(p)
		}
		valid := handleJump(p)
		if valid {
			p.CanDash = true
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
		println("Jump")
		if p.IsAnimationLocked {
			p.ActionAnimations[p.CurrentAnimationIndex].ResetAnimation = true
		}
		p.HasSecondJump = true
		p.CanJump = false
		p.CurrentAnimationIndex = 2
		p.Physics.NetForce.Y += controls.JUMP.ForceY
		p.Physics.NetForce.X += controls.JUMP.ForceX
		return true
	} else if p.HasSecondJump && p.IsAirborn {
		println("dblJump")
		if p.IsAnimationLocked {
			p.ActionAnimations[p.CurrentAnimationIndex].ResetAnimation = true
		}
		p.CurrentAnimationIndex = 3
		p.HasSecondJump = false
		p.Physics.NetForce.Y += controls.JUMP.ForceY
		p.Physics.NetForce.X += controls.JUMP.ForceX
		return true
	}
	return false
}

func handleDash(p *sprites.Player, movingRight bool) bool {
	if p.CanDash {
		p.CanDash = false
		p.IsMovingRight = movingRight
		p.CurrentAnimationIndex = 1
<<<<<<< HEAD
=======
		p.DashLastUsed = time.Now()
		if p.IsAnimationLocked {
			p.ActionAnimations[p.CurrentAnimationIndex].ResetAnimation = true
		}
		p.Physics.NetForce.Y += controls.DASHLEFT.ForceY
		p.Physics.NetForce.X += controls.DASHLEFT.ForceX
		return true
	}
	return false
}

func handleDashLeft(p *sprites.Player) bool {
	if time.Since(p.DashLastUsed) >= p.DashCooldowntime {
		p.IsMovingRight = false
		p.CurrentAnimationIndex = 1
		p.DashLastUsed = time.Now()
>>>>>>> 16f3c53bd428513a1c986c0fe9c23443b6469d9c
		if p.IsAnimationLocked {
			p.ActionAnimations[p.CurrentAnimationIndex].ResetAnimation = true
		}
		p.Physics.NetForce.Y += controls.DASHRIGHT.ForceY
		p.Physics.NetForce.X += controls.DASHRIGHT.ForceX
		return true
	}
	return false
}

func handleWallHang(p *sprites.Player) bool {
	p.CurrentAnimationIndex = 4
	p.IsAnimationLocked = true
	return false
}
