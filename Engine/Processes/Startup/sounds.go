package startup

import (
	sound "GoPlat/gameComponents/sounds"
)

func GetAllSounds() *sound.SoundManager {
	return sound.GetAllSounds()
}
