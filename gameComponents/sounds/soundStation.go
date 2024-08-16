package sound

import (
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

type SoundStation struct {
	audioContext *audio.Context
	audioPlayers []*audio.Player
}

/*
func (s *SoundStation) loadSound(filepath, fileType string) (*audio.Player, error) {
	soundFile, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer soundFile.Close()

	switch fileType {
	case "mp3":
		decoder, err := mp3.newDecoder(soundFile)
		if err != nil {
			return nil, err
		}

		player, err := audio.NewPlayer(s.audioContext, decoder)
		return player, nil
	case "wav":
		decoder, err := wav.newDecoder(soundFile)
		if err != nil {
			return nil, err
		}

		player, err := audio.NewPlayer(s.audioContext, decoder)
		return player, nil
	default:
		panic("audioType not found")
	}
}*/
