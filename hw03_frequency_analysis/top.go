package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type entrySlice struct {
	key   string
	value int
}

func countWordsInText(text string) map[string]int {
	wordCounter := make(map[string]int)

	words := strings.Fields(text)

	for _, word := range words {
		wordCounter[word]++
	}

	return wordCounter
}

func Top10(srcString string) []string {
	wordCounter := countWordsInText(srcString)

	sortedSlice := make([]entrySlice, 0, len(wordCounter))

	for key, value := range wordCounter {
		sortedSlice = append(sortedSlice, entrySlice{key, value})
	}

	sort.Slice(sortedSlice, func(i, j int) bool {
		if sortedSlice[j].value == sortedSlice[i].value {
			return sortedSlice[j].key > sortedSlice[i].key
		}

		return sortedSlice[j].value < sortedSlice[i].value
	})

	var resSlice []string
	for i := 0; i < min(10, len(sortedSlice)); i++ {
		resSlice = append(resSlice, sortedSlice[i].key)
	}

	return resSlice
}
