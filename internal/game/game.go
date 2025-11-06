package game

import (
	"fmt"
	"math/rand"
	dictionary "spellingbee_game/internal/directory"
	"strings"
	"time"
)

type Game struct {
	Letters      []string
	CenterLetter string
	Score        int
	Strategy     ScoringStrategy
}

func NewGame() *Game {
	rand.Seed(time.Now().UnixNano())

	dict, err := dictionary.GetInstance()
	if err != nil {
		fmt.Println("Error loading dictionary:", err)
		return randomFallbackGame()
	}

	word := dict.GetRandomWord()
	word = strings.ToUpper(word)

	unique := map[rune]bool{}
	for _, ch := range word {
		if ch >= 'A' && ch <= 'Z' {
			unique[ch] = true
		}
	}

	alphabet := []rune("ABCDEFGHIJKLMNOPQRTUVWXYZ")

	for len(unique) < 8 {
		r := alphabet[rand.Intn(len(alphabet))]
		unique[r] = true
	}

	letters := make([]string, 0, 8)
	for ch := range unique {
		letters = append(letters, string(ch))
	}

	if len(letters) > 8 {
		rand.Shuffle(len(letters), func(i, j int) { letters[i], letters[j] = letters[j], letters[i] })
		letters = letters[:8]
	}

	center := letters[3]

	return &Game{
		Letters:      letters,
		CenterLetter: center,
		Score:        0,
		Strategy:     &BasicStrategy{},
	}
}

func (g *Game) ValidateWord(word string) (bool, string) {
	word = strings.ToUpper(word)

	if !strings.Contains(word, g.CenterLetter) {
		return false, fmt.Sprintf("Word must contain the letter '%s'", g.CenterLetter)
	}

	for _, ch := range word {
		if !contains(g.Letters, string(ch)) {
			return false, fmt.Sprintf("Invalid letter '%s'", string(ch))
		}
	}

	if len(word) < 4 {
		return false, "Word must be at least 4 letters long"
	}

	dict, err := dictionary.GetInstance()
	if err != nil {
		return false, "Dictionary could not be loaded"
	}
	if !dict.IsValidWord(word) {
		return false, "Word not found in dictionary"
	}

	return true, "Valid word"
}

func (g *Game) GetTotalScore() int {
	return g.Score
}

func (g *Game) AddScore(word string) int {
	if g.Strategy == nil {
		g.Strategy = &BasicStrategy{}
	}
	basePoints := g.Strategy.Calculate(word)

	isPangram := true
	upper := strings.ToUpper(word)
	for _, l := range g.Letters {
		if !strings.Contains(upper, l) {
			isPangram = false
			break
		}
	}

	points := basePoints
	if isPangram {
		points += 7
	}

	g.Score += points
	return points
}

func contains(list []string, ch string) bool {
	for _, c := range list {
		if c == ch {
			return true
		}
	}
	return false
}

func randomFallbackGame() *Game {
	rand.Seed(time.Now().UnixNano())

	alphabet := []rune("ABCDEFGHIJKLMNOPQRTUVWXYZ")
	used := make(map[rune]bool)
	letters := make([]string, 0, 7)

	for len(letters) < 7 {
		r := alphabet[rand.Intn(len(alphabet))]
		if !used[r] {
			used[r] = true
			letters = append(letters, string(r))
		}
	}

	center := letters[rand.Intn(len(letters))]

	return &Game{
		Letters:      letters,
		CenterLetter: center,
		Score:        0,
		Strategy:     &BasicStrategy{},
	}
}
