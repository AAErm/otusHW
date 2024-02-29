package hw03frequencyanalysis

import (
	"container/heap"
)

var taskWithAsteriskIsCompleted = false

type wordCount struct {
	word  string
	count int
}

type wordHeap []wordCount

func (h wordHeap) Len() int {
	return len(h)
}

func (h wordHeap) Less(i, j int) bool {
	if h[i].count == h[j].count {
		return h[i].word > h[j].word
	}
	return h[i].count < h[j].count
}

func (h wordHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *wordHeap) Push(x interface{}) {
	*h = append(*h, x.(wordCount))
}

func (h *wordHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Top10 returns the 10 most frequently occurring words.
func Top10(text string) []string {
	parseWords := splitWithOrWithoutAsterisk(taskWithAsteriskIsCompleted, text)
	counters := make(map[string]int)

	for _, word := range parseWords {
		if taskWithAsteriskIsCompleted && word == "-" {
			continue
		}
		counters[word]++
	}

	wh := &wordHeap{}
	heap.Init(wh)
	for word, count := range counters {
		heap.Push(wh, wordCount{word, count})
		if wh.Len() > maxWord {
			heap.Pop(wh)
		}
	}

	result := make([]string, 0, maxWord)
	for wh.Len() > 0 {
		word := heap.Pop(wh).(wordCount).word
		result = append([]string{word}, result...)
	}

	return result
}
