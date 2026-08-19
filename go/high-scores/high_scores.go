package highscores

type HighScores struct {
	scores []int
}

// NewHighScores returns a new HighScores object.
func NewHighScores(scores []int) *HighScores {
	return &HighScores{scores: scores}
}

// Scores returns all the scores.
func (s *HighScores) Scores() []int {
	return s.scores
}

// Latest returns the latest (last) score.
func (s *HighScores) Latest() int {
	return s.scores[len(s.scores)-1]
}

// PersonalBest returns the best (highest) score.
func (s *HighScores) PersonalBest() int {
	max := 0
	for key, value := range s.scores {
		if key == 0 {
			max = value
		}

		if value > max {
			max = value
		}
	}

	return max
}

// TopThree returns the top three scores.
func (s *HighScores) TopThree() []int {
	panic("Please implement the TopThree function")
}
