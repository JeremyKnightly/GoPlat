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

	err := ss.LoadSound("Assets/Audio/Footfalls/Footstep1.wav", "wav", "Footstep", .2)
	if err != nil {
		panic(err)
	}

	err1 := ss.LoadSound("Assets/Audio/Footfalls/sfx_Hurt.wav", "wav", "Hurt", .6)
	if err1 != nil {
		panic(err1)
	}

	err2 := ss.LoadSound("Assets/Audio/Footfalls/sfx_Death.wav", "wav", "Death", .8)
	if err1 != nil {
		panic(err2)
	}
}
