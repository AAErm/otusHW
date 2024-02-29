package hw03frequencyanalysis

var maxWord = 10

// Top10_2 returns the 10 most frequently occurring words.
func Top10_2(text string) []string {
	parseWords := splitWithOrWithoutAsterisk(taskWithAsteriskIsCompleted, text)
	counters := make(map[string]int)
	for _, word := range parseWords {
		if taskWithAsteriskIsCompleted && word == "-" {
			continue
		}
		counters[word]++
	}
	var (
		topWords []string
		topCount []int
		position int
	)
	for word, qty := range counters {
		if word == "" {
			continue
		}
		position, topCount = addElem(qty, topCount)
		topWords = addWord(topWords, word, position)
		if len(topWords) > maxWord {
			topWords = topWords[:maxWord]
			topCount = topCount[:maxWord]
		}
	}

	return topWords
}

func addElem(number int, arr []int) (int, []int) {
	pos := len(arr)
	for i := len(arr) - 1; i >= 0; i-- {
		if number < arr[i] {
			break
		}
		pos = i
	}
	if pos == len(arr) {
		return len(arr), append(arr, number)
	}
	arr = append(arr[:pos+1], arr[pos:]...)
	arr[pos] = number
	return pos, arr
}

func addWord(words []string, word string, position int) []string {
	if position == len(words) {
		return append(words, word)
	}
	words = append(words[:position+1], words[position:]...)
	words[position] = word

	return words
}
