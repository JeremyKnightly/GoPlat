package main

func (g *Game) playSFX() {
	if g.Player.IsIdle {
		return
	}
	/*
		Dash,      //1
		DblJump,   //3
		EdgeClimb, //4
		WallGrab,  //6
		WallSlide, //7
	*/
	switch g.Player.CurrentAnimationIndex {
	case 0:
		g.handleWalkSFX()
	case 2:
		g.handleJumpSFX()
	case 5:
		g.handleHurtSFX()
	case 7:
		g.handleWallSlideSFX()
	case 8:
		g.handleDeathSFX()
	}
}

func (g *Game) loopBGM(trackNum int) {
	bgmPlayer, err := g.SoundManager.GetStation("BGM").GetSoundPlayerByNum(trackNum)
	if err != nil {
		panic(err)
	}
	bgmPlayer.ResetOnTrackComplete()
}

func (g *Game) handleWalkSFX() {
	if g.Player.IsAirborn {
		return
	}
	audioPlayer, err := g.SoundManager.GetStation("SFX").GetSoundPlayerByName("Footstep1")
	if err != nil {
		println("Could not getSoundPlayer", err)
	}
	audioPlayer2, err1 := g.SoundManager.GetStation("SFX").GetSoundPlayerByName("Footstep2")
	if err1 != nil {
		println("Could not getSoundPlayer", err1)
	}
	frameIdx := g.Player.ActionAnimations[0].CurrentFrameIndex
	if (frameIdx == 0) && !audioPlayer.RecentlyPlayed {
		err = audioPlayer.Play()
		audioPlayer.RecentlyPlayed = true
		if err != nil {
			println("Could not Play", err)
		}
	} else if (frameIdx == 6) && audioPlayer.RecentlyPlayed {
		err = audioPlayer.Rewind()
		audioPlayer.RecentlyPlayed = false
		if err != nil {
			println("Could not Rewind", err)
		}
	} else if (frameIdx == 4) && !audioPlayer2.RecentlyPlayed {
		err = audioPlayer2.Play()
		audioPlayer2.RecentlyPlayed = true
		if err != nil {
			println("Could not Play", err)
		}
	} else if (frameIdx == 2) && audioPlayer2.RecentlyPlayed {
		err = audioPlayer2.Rewind()
		audioPlayer2.RecentlyPlayed = false
		if err != nil {
			println("Could not Rewind", err)
		}
	}
}

func (g *Game) handleDeathSFX() {
	audioPlayer, err := g.SoundManager.GetStation("SFX").GetSoundPlayerByName("Death")
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
	audioPlayer, err := g.SoundManager.GetStation("SFX").GetSoundPlayerByName("Hurt")
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

func (g *Game) handleWallSlideSFX() {
	audioPlayer, err := g.SoundManager.GetStation("SFX").GetSoundPlayerByName("WallSlide")
	if err != nil {
		println("Could not getSoundPlayer", err)
	}
	frameIdx := g.Player.ActionAnimations[7].CurrentFrameIndex
	if frameIdx == 1 && !audioPlayer.RecentlyPlayed {
		err = audioPlayer.Play()
		audioPlayer.RecentlyPlayed = true
		if err != nil {
			println("Could not Play", err)
		}
	} else if frameIdx >= 6 && audioPlayer.RecentlyPlayed {
		err = audioPlayer.Rewind()
		audioPlayer.RecentlyPlayed = false
		if err != nil {
			println("Could not Rewind", err)
		}
	}
}

func (g *Game) handleJumpSFX() {
	audioPlayer, err := g.SoundManager.GetStation("SFX").GetSoundPlayerByName("Jump")
	if err != nil {
		println("Could not getSoundPlayer", err)
	}
	frameIdx := g.Player.ActionAnimations[2].CurrentFrameIndex
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
