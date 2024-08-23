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

func (ss *SoundStation) GetBGM() {
	ss.StationName = "BGM"

	/*err := ss.LoadSound("Assets/Audio/Footfalls/grass 1.wav", "wav", "Footstep")
	if err != nil {
		panic(err)
	}*/
}

func (ss *SoundStation) GetSFX() {
	ss.StationName = "SFX"

	err := ss.LoadSound("Assets/Audio/Final/SFX/Footstep1.wav", "wav", "Footstep1", .1)
	if err != nil {
		panic(err)
	}

	err = ss.LoadSound("Assets/Audio/Final/SFX/Footstep2.wav", "wav", "Footstep2", .1)
	if err != nil {
		panic(err)
	}

	err = ss.LoadSound("Assets/Audio/Final/SFX/Hurt.wav", "wav", "Hurt", .6)
	if err != nil {
		panic(err)
	}

	err = ss.LoadSound("Assets/Audio/Final/SFX/Death.wav", "wav", "Death", .8)
	if err != nil {
		panic(err)
	}

	err = ss.LoadSound("Assets/Audio/Final/SFX/WallSlide.wav", "wav", "WallSlide", .8)
	if err != nil {
		panic(err)
	}

	err = ss.LoadSound("Assets/Audio/Final/SFX/Jump.wav", "wav", "Jump", .3)
	if err != nil {
		panic(err)
	}
}
