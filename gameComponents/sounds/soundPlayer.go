package sound

import (
	"fmt"
)

func (ss *SoundStation) GetSoundPlayer(trackName string) (*AudioPlayer, error) {
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
