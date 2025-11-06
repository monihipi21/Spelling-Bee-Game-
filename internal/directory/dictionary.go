package dictionary

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

var instance *Dictionary
var once sync.Once

type Dictionary struct {
	words map[string]bool
	Words []string
}

// ładowanie danych
func loadWordsFromJSON(filePath string) (map[string]bool, []string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var data map[string]int
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return nil, nil, err
	}

	wordsMap := make(map[string]bool)
	wordsList := make([]string, 0, len(data))

	for w := range data {
		word := strings.ToUpper(w)
		wordsMap[word] = true
		wordsList = append(wordsList, word)
	}

	return wordsMap, wordsList, nil
}

// najwazniejsza czesc wzorca, gdzie jest realizowany
func GetInstance() (*Dictionary, error) {
	var err error
	once.Do(func() {
		var words map[string]bool
		var list []string
		words, list, err = loadWordsFromJSON("words_dictionary.json")
		if err != nil {
			fmt.Println("Error", err)
			return
		}
		instance = &Dictionary{
			words: words,
			Words: list,
		}
	})
	if err != nil {
		return nil, err
	}
	return instance, nil
}

// metody
func (d *Dictionary) IsValidWord(word string) bool {
	word = strings.ToUpper(word)
	_, exists := d.words[word]
	return exists
}

func (d *Dictionary) GetRandomWord() string {
	if len(d.Words) == 0 {
		return "RANDOM"
	}
	rand.Seed(time.Now().UnixNano())
	return d.Words[rand.Intn(len(d.Words))]
}
