package model

type Difficulty int

const (
	Easy Difficulty = iota
	Medium
	Hard
)

var difficultyStrings = map[Difficulty]string{
	Easy:   "Easy",
	Medium: "Medium",
	Hard:   "Hard",
}

func (d Difficulty) String() string {
	return difficultyStrings[d]
}

func ParseDifficulty(difficultyStr string) Difficulty {
	for diff, diffStr := range difficultyStrings {
		if diffStr == difficultyStr {
			return diff
		}
	}
	return Easy
}
