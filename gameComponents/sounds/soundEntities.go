package sound

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

const SampleRate = 44100

type SoundStation struct {
	StationName  string
	audioContext *audio.Context
	players      []*AudioPlayer
}

type AudioPlayer struct {
	player         *audio.Player
	TrackName      string
	RecentlyPlayed bool
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

func (ap *AudioPlayer) Play() error {
	if ap.player == nil {
		return fmt.Errorf("player is not initialized")
	}
	ap.player.Play()
	return nil
}

func (ap *AudioPlayer) Rewind() error {
	if ap.player == nil {
		return fmt.Errorf("player is not initialized")
	}
	ap.player.Rewind()
	return nil
}

func (ap *AudioPlayer) Pause() error {
	if ap.player == nil {
		return fmt.Errorf("player is not initialized")
	}
	ap.player.Pause()
	return nil
}

func (ap *AudioPlayer) SetVolume(volume float64) error {
	if ap.player == nil {
		return fmt.Errorf("player is not initialized")
	}
	ap.player.SetVolume(volume)
	return nil
}
