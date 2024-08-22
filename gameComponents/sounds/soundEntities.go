package sound

import (
	"github.com/hajimehoshi/ebiten/v2/audio"
)

const SampleRate = 44100

type SoundStation struct {
	StationName  string
	audioContext *audio.Context
	players      []*AudioPlayer
}

type AudioPlayer struct {
	player    *audio.Player
	TrackName string
}

type SoundManager struct {
	stations []*SoundStation
}

func NewSoundStation(context *audio.Context) *SoundStation {
	return &SoundStation{
		audioContext: context,
		players:      []*AudioPlayer{},
	}
}

func CreateContext() *audio.Context {
	audioContext := audio.NewContext(SampleRate)
	return audioContext
}

func (sm *SoundManager) GetStation(stationName string) *SoundStation {
	for _, station := range sm.stations {
		if station.StationName == stationName {
			return station
		}
	}
	return nil
}

func NewSoundManager() *SoundManager {
	return &SoundManager{}
}
