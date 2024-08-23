package main

func (g *Game) playSFX() {
	if g.Player.IsIdle {
		return
	}
	/*
		Dash,      //1
		Jump,      //2
		DblJump,   //3
		EdgeClimb, //4
		Hurt,      //5
		WallGrab,  //6
		WallSlide, //7
	*/
	switch g.Player.CurrentAnimationIndex {
	case 0:
		g.handleWalkSFX()
	case 5:
		g.handleHurtSFX()
	case 8:
		g.handleDeathSFX()
	}
}

func (g *Game) handleWalkSFX() {
	if g.Player.IsAirborn {
		return
	}
	audioPlayer, err := g.SoundManager.GetStation("SFX").GetSoundPlayer("Footstep")
	if err != nil {
		println("Could not getSoundPlayer", err)
	}
	frameIdx := g.Player.ActionAnimations[0].CurrentFrameIndex
	if (frameIdx%4 == 0) && !audioPlayer.RecentlyPlayed {
		err = audioPlayer.Play()
		audioPlayer.RecentlyPlayed = true
		if err != nil {
			println("Could not Play", err)
		}
	} else if (frameIdx%4 == 3) && audioPlayer.RecentlyPlayed {
		err = audioPlayer.Rewind()
		audioPlayer.RecentlyPlayed = false
		if err != nil {
			println("Could not Rewind", err)
		}
	}
}

func (g *Game) handleDeathSFX() {
	audioPlayer, err := g.SoundManager.GetStation("SFX").GetSoundPlayer("Hurt")
	if err != nil {
		println("Could not getSoundPlayer", err)
	}
	frameIdx := g.Player.ActionAnimations[8].CurrentFrameIndex
	if frameIdx == 1 && !audioPlayer.RecentlyPlayed {
		err = audioPlayer.Play()
		audioPlayer.RecentlyPlayed = true
		if err != nil {
			println("Could not Play", err)
		}
	} else if frameIdx >= 7 && audioPlayer.RecentlyPlayed {
		err = audioPlayer.Rewind()
		audioPlayer.RecentlyPlayed = false
		if err != nil {
			println("Could not Rewind", err)
		}
	}
}

func (g *Game) handleHurtSFX() {
	audioPlayer, err := g.SoundManager.GetStation("SFX").GetSoundPlayer("Death")
	if err != nil {
		println("Could not getSoundPlayer", err)
	}
	frameIdx := g.Player.ActionAnimations[5].CurrentFrameIndex
	if frameIdx == 1 && !audioPlayer.RecentlyPlayed {
		err = audioPlayer.Play()
		audioPlayer.RecentlyPlayed = true
		if err != nil {
			println("Could not Play", err)
		}
	} else if frameIdx >= 5 && audioPlayer.RecentlyPlayed {
		err = audioPlayer.Rewind()
		audioPlayer.RecentlyPlayed = false
		if err != nil {
			println("Could not Rewind", err)
		}
	}
}
