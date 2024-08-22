package sound

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

func (ss *SoundStation) LoadSound(filepath, fileType, trackName string) error {
	if ss.audioContext == nil {
		println("nil context")
	}

	soundFile, err := os.Open(filepath)
	if err != nil {
		return err
	}
	//defer soundFile.Close()

	switch fileType {
	case "mp3":
		decoder, err := mp3.DecodeWithSampleRate(SampleRate, soundFile)
		if err != nil {
			return err
		}

		player, err := ss.audioContext.NewPlayer(decoder)
		if err != nil {
			return err
		}
		ss.players = append(ss.players, &AudioPlayer{
			player:    player,
			TrackName: trackName,
		})
	case "wav":
		decoder, err := wav.DecodeWithSampleRate(SampleRate, soundFile)
		if err != nil {
			return err
		}
		player, err := ss.audioContext.NewPlayer(decoder)
		if err != nil {
			return err
		}
		ss.players = append(ss.players, &AudioPlayer{
			player:    player,
			TrackName: trackName,
		})
	default:
		panic("audioType not found")
	}

	return nil
}
