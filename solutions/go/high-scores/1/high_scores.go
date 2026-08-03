package highscores

import "slices"

type HighScores struct{
    scoresInside []int
}

// NewHighScores returns a new HighScores object.
func NewHighScores(scores []int) *HighScores {
	newNumbers := append([]int(nil), scores...)
	newHighScores := HighScores{
		scoresInside: newNumbers,
	}
    return &newHighScores
}

// Scores returns all the scores.
func (s *HighScores) Scores() []int {
	return s.scoresInside
}

// Latest returns the latest (last) score.
func (s *HighScores) Latest() int {
	return s.scoresInside[len(s.scoresInside) - 1]
}

// PersonalBest returns the best (highest) score.
func (s *HighScores) PersonalBest() int {
    return slices.Max(s.scoresInside)
}

// TopThree returns the top three scores.
func (s *HighScores) TopThree() []int {
	result := slices.Clone(s.scoresInside)
	slices.Sort(result)
    slices.Reverse(result)
	if len(result) >= 3 {
        return result[:3]
    } else {
        return result
    }
}
