package sound

func GetAllSounds() *SoundManager {
	manager := NewSoundManager()
	context := CreateContext()

	station := NewSoundStation(context)
	station.GetBGM()
	manager.stations = append(manager.stations, station)

	station = NewSoundStation(context)
	station.GetSFX()
	manager.stations = append(manager.stations, station)

	return manager
}

func (ss *SoundStation) StopOtherStationSounds(stopIdx int) {
	for idx, player := range ss.players {
		if stopIdx == idx {
			continue
		} else {
			player.Rewind()
			player.Pause()
		}

		//player.Rewind()
	}
}

func (ss *SoundStation) GetBGM() {
	ss.StationName = "BGM"

	err := ss.LoadSound("Assets/Audio/Final/BGM/Track 1.wav", "wav", "Track 1", .1)
	if err != nil {
		panic(err)
	}
	err = ss.LoadSound("Assets/Audio/Final/BGM/Track 2.wav", "wav", "Track 2", .1)
	if err != nil {
		panic(err)
	}
	err = ss.LoadSound("Assets/Audio/Final/BGM/Track 3.wav", "wav", "Track 3", .1)
	if err != nil {
		panic(err)
	}
	err = ss.LoadSound("Assets/Audio/Final/BGM/Track 4.wav", "wav", "Track 4", .1)
	if err != nil {
		panic(err)
	}
	err = ss.LoadSound("Assets/Audio/Final/BGM/Track 5.wav", "wav", "Track 5", .1)
	if err != nil {
		panic(err)
	}
}

func (ss *SoundStation) GetSFX() {
	ss.StationName = "SFX"

	err := ss.LoadSound("Assets/Audio/Final/SFX/Footstep1.wav", "wav", "Footstep1", .27)
	if err != nil {
		panic(err)
	}

	err = ss.LoadSound("Assets/Audio/Final/SFX/Footstep2.wav", "wav", "Footstep2", .27)
	if err != nil {
		panic(err)
	}

	err = ss.LoadSound("Assets/Audio/Final/SFX/Hurt.wav", "wav", "Hurt", .5)
	if err != nil {
		panic(err)
	}

	err = ss.LoadSound("Assets/Audio/Final/SFX/Death.wav", "wav", "Death", .6)
	if err != nil {
		panic(err)
	}

	err = ss.LoadSound("Assets/Audio/Final/SFX/WallSlide.wav", "wav", "WallSlide", .4)
	if err != nil {
		panic(err)
	}

	err = ss.LoadSound("Assets/Audio/Final/SFX/Jump.wav", "wav", "Jump", .35)
	if err != nil {
		panic(err)
	}
}
