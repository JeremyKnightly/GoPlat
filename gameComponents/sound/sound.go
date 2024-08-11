package sound

import (
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

var (
	audioContext *audio.Context
	player       *audio.Player
)

func init() {
	//audioContext = audio.NewContext(44100)
	audioContext = audio.NewContext(11100)
}

/*
func loadSound(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Printf("Could not open sound file: %w", err)
	}s
	defer file.Close()

	decodedAudio, err := wav.DecodeWithSampleRate(audioContext, file)


}*/
