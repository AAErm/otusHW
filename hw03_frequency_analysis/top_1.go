package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

// Top10Old returns the 10 most frequently occurring words.
func Top10_first(text string) []string {
	parseWords := splitWithOrWithoutAsterisk(taskWithAsteriskIsCompleted, text)
	counters := make(map[string]int)

	for _, word := range parseWords {
		if taskWithAsteriskIsCompleted && word == "-" {
			continue
		}
		counters[word]++
	}

	words := make([]wordCount, 0, len(counters))
	for word, count := range counters {
		words = append(words, wordCount{word, count})
	}

	sort.Slice(words, func(i, j int) bool {
		if words[i].count == words[j].count {
			return words[i].word < words[j].word
		}
		return words[i].count > words[j].count
	})

	result := make([]string, 0, maxWord)
	for i, wordCount := range words {
		if i >= maxWord {
			break
		}
		result = append(result, wordCount.word)
	}

	return result
}

// p{L} - any letters of any language.
// p{N} - any kind of numeric characters in any language.
var wordRegex = regexp.MustCompile(`[\p{L}\p{N}-]+`)

func splitWithOrWithoutAsterisk(sign bool, text string) []string {
	if !sign {
		return strings.Fields(text)
	}

	text = strings.ToLower(text)
	words := wordRegex.FindAllString(text, -1)
	return words
}
