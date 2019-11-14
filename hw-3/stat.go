package hw3

import (
	"sort"
	"strings"
	"unicode"
)

// Top10 returns the 10 most frequently occurring words
func Top10(text string) []string {
	text = strings.ToLower(text)

	// get words
	f := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	words := strings.FieldsFunc(text, f) //strings.FieldsFunc(s, unicode.IsSpace)

	// get the count of each word
	counts := make(map[string]int, len(words))
	for _, w := range words {
		counts[w]++
	}

	// sorting by count
	type WordFreq struct {
		word string
		freq int
	}
	wordFreqList := make([]WordFreq, 0, len(counts))
	for k, v := range counts {
		wordFreqList = append(wordFreqList, WordFreq{k, v})
	}
	sort.Slice(wordFreqList, func(i, j int) bool {
		return wordFreqList[i].freq > wordFreqList[j].freq
	})

	// get top 10 words
	count := len(wordFreqList)
	if count > 10 {
		count = 10
	}
	topWordFreqList := wordFreqList[:count]
	top := make([]string, 0, count)
	for _, v := range topWordFreqList {
		top = append(top, v.word)
	}

	return top
}
