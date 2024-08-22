package sound

import "fmt"

func (ss *SoundStation) PlaySound(trackName string) error {
	playIdx := 0
	soundFound := false
	for idx, audioPlayer := range ss.players {
		if audioPlayer.TrackName == trackName {
			playIdx = idx
			soundFound = true
			break
		}
	}
	if !soundFound {
		return fmt.Errorf("sound '%s' not found", trackName)
	}
	ss.players[playIdx].player.Play()
	return nil
}
