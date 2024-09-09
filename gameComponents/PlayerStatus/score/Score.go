package score

import "time"

const DEFAULT_SCORE_VALUE = 100000
const PUP_POINT_VALUE = 4000
const ITEM_POINT_VALUE = 5000
const KEY_POINT_VALUE = 8000
const SCROLL_POINT_VALUE = 2500
const DEATH_POINT_LOSS = 1000

type Score struct {
	GameStartTime                                      time.Time
	ScoreValue, KeysCollected                          int
	Deaths, PUpsCollected, ItemsCollected, ScrollsRead int
}

func (s *Score) SetGameStartTime() {
	s.GameStartTime = time.Now()
}

func (s *Score) evaluateScore() {
	timeScoreDepletion := int(time.Since(s.GameStartTime).Milliseconds() / 100)
	s.ScoreValue = DEFAULT_SCORE_VALUE - timeScoreDepletion
	s.ScoreValue += s.PUpsCollected * PUP_POINT_VALUE
	s.ScoreValue += s.ItemsCollected * ITEM_POINT_VALUE
	s.ScoreValue += s.KeysCollected * KEY_POINT_VALUE
	s.ScoreValue += s.ScrollsRead * SCROLL_POINT_VALUE
	s.ScoreValue -= s.Deaths * DEATH_POINT_LOSS
}

func (s *Score) GetScore() int {
	s.evaluateScore()
	return s.ScoreValue
}

func (s *Score) ResetScore() {
	s.ScoreValue = DEFAULT_SCORE_VALUE
	s.PUpsCollected = 0
	s.ItemsCollected = 0
	s.KeysCollected = 0
	s.ScrollsRead = 0
}

func (s *Score) AddPoints(numItems int, itemType string) {
	switch itemType {
	case "item":
		s.addItemPoints(numItems)
	case "key":
		s.addKeyPoints(numItems)
	case "pup":
		s.addPUpPoints(numItems)
	case "scroll":
		s.addScrollPoints(numItems)
	}
}

func (s *Score) SubPoints(numItems int, itemType string) {
	switch itemType {
	case "item":
		s.subItemPoints(numItems)
	case "key":
		s.subKeyPoints(numItems)
	case "pup":
		s.subPUpPoints(numItems)
	case "scroll":
		s.subScrollPoints(numItems)
	}
}

func (s *Score) addKeyPoints(numKeys int) {
	s.KeysCollected += numKeys
}

func (s *Score) subKeyPoints(numKeys int) {
	s.KeysCollected -= numKeys
}

func (s *Score) addItemPoints(numItems int) {
	s.ItemsCollected += numItems
}

func (s *Score) subItemPoints(numItems int) {
	s.ItemsCollected -= numItems
}

func (s *Score) addPUpPoints(numPUps int) {
	s.PUpsCollected += numPUps
}

func (s *Score) subPUpPoints(numPUps int) {
	s.PUpsCollected -= numPUps
}

func (s *Score) addScrollPoints(numScrolls int) {
	s.ScrollsRead += numScrolls
}

func (s *Score) subScrollPoints(numScrolls int) {
	s.ScrollsRead -= numScrolls
}

func (s *Score) AddDeath() {
	s.Deaths++
}

func (s *Score) GetDeathCount() int {
	return s.Deaths
}
