package score

import "time"

const DEFAULT_SCORE_VALUE = 100000
const PUP_POINT_VALUE = 4000
const ITEM_POINT_VALUE = 5000
const SCROLL_POINT_VALUE = 2500

type Score struct {
	GameStartTime                              time.Time
	ScoreValue                                 int
	PUpsCollected, ItemsCollected, ScrollsRead int
}

func (s *Score) SetGameStartTime() {
	s.GameStartTime = time.Now()
}

func (s *Score) evaluateScore() {
	timeScoreDepletion := int(time.Since(s.GameStartTime).Milliseconds() / 10)
	s.ScoreValue = DEFAULT_SCORE_VALUE - timeScoreDepletion
	s.ScoreValue += s.PUpsCollected * PUP_POINT_VALUE
	s.ScoreValue += s.ItemsCollected * ITEM_POINT_VALUE
	s.ScoreValue += s.ScrollsRead * SCROLL_POINT_VALUE
}

func (s *Score) GetScore() int {
	s.evaluateScore()
	return s.ScoreValue
}

func (s *Score) ResetScore() {
	s.ScoreValue = DEFAULT_SCORE_VALUE
	s.PUpsCollected = 0
	s.ItemsCollected = 0
	s.ScrollsRead = 0
}

func (s *Score) AddItemPoints() {
	s.ItemsCollected++
}

func (s *Score) AddPUpPoints() {
	s.PUpsCollected++
}

func (s *Score) AddScrollPoints() {
	s.ScrollsRead++
}

func (s *Score) SubItemPoints() {
	s.ItemsCollected--
}

func (s *Score) SubPUpPoints() {
	s.PUpsCollected--
}

func (s *Score) SubScrollPoints() {
	s.ScrollsRead--
}
