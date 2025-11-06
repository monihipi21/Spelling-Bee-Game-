package game

import (
	"fmt"
	"log"
	"os"
)

type GameInterface interface {
	ValidateWord(word string) (bool, string)
	AddScore(word string) int
	GetTotalScore() int
}

// sturktura decoratora
type LoggingDecorator struct {
	Wrapped GameInterface
	Logger  *log.Logger
}

// fabryka decoratora
func NewLoggingDecorator(wrapped GameInterface) *LoggingDecorator {
	file, err := os.OpenFile("game.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error creating log file:", err)
	}

	logger := log.New(file, "[LOG] ", log.LstdFlags)
	return &LoggingDecorator{Wrapped: wrapped, Logger: logger}
}

// metody decoratora
func (l *LoggingDecorator) ValidateWord(word string) (bool, string) {
	valid, msg := l.Wrapped.ValidateWord(word)
	l.Logger.Printf("Checked word: %s → %v (%s)", word, valid, msg)
	return valid, msg
}

func (l *LoggingDecorator) AddScore(word string) int {
	score := l.Wrapped.AddScore(word)
	l.Logger.Printf("Scored %d for word: %s", score, word)
	return score
}

func (l *LoggingDecorator) GetTotalScore() int {
	return l.Wrapped.GetTotalScore()
}
