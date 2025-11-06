package game

// definicja strategii
type ScoringStrategy interface {
	Calculate(word string) int
}

// odpowiada podstawowej punktacji
type BasicStrategy struct{}

func (b *BasicStrategy) Calculate(word string) int {
	if len(word) <= 4 {
		return 1
	}
	return len(word)
}

// zmienia sposob liczenia punktow
type BonusStrategy struct{}

func (b *BonusStrategy) Calculate(word string) int {
	points := len(word)
	if len(word) >= 6 {
		points += 3
	}
	return points
}
