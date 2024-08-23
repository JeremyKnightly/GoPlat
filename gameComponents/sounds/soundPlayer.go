package sound

import (
	"fmt"
)

func (ss *SoundStation) GetSoundPlayerByName(trackName string) (*AudioPlayer, error) {
	playIdx := 0
	for idx, audioPlayer := range ss.players {
		if audioPlayer.TrackName == trackName {
			playIdx = idx
			return ss.players[playIdx], nil
		}
	}
	println("track not found")
	return &AudioPlayer{}, fmt.Errorf("sound '%s' not found", trackName)
}

func (ss *SoundStation) GetSoundPlayerByNum(trackNum int) (*AudioPlayer, error) {
	if trackNum >= len(ss.players) {
		return nil, fmt.Errorf("Index Out of Bounds")
	}

	return ss.players[trackNum], nil
}
