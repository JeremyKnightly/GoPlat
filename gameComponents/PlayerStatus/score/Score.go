package score

const DEFAULT_SCORE_VALUE = 100000
const PUP_POINT_VALUE = 4000
const ITEM_POINT_VALUE = 5000
const SCROLL_POINT_VALUE = 2500

type Score struct {
	Score                                      int
	PUpsCollected, ItemsCollected, ScrollsRead int
}

func (s *Score) ResetScore() {
	s.Score = DEFAULT_SCORE_VALUE
	s.PUpsCollected = 0
	s.ItemsCollected = 0
	s.ScrollsRead = 0
}

func (s *Score) AddItemPoints() {
	s.Score += ITEM_POINT_VALUE
	s.ItemsCollected++
}

func (s *Score) AddPUpPoints() {
	s.Score += PUP_POINT_VALUE
	s.PUpsCollected++
}

func (s *Score) AddScrollPoints() {
	s.Score += SCROLL_POINT_VALUE
	s.ScrollsRead++
}
