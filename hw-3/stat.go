package hw3

import (
	//"fmt"
	"strings"
	"unicode"
    "sort"
)

// Top10 returns the 10 most frequently occurring words
func Top10(text string) []string {
    text = strings.ToLower(text)
    
    f := func(c rune) bool {
		//return !unicode.IsLetter(c)
        return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
    words := strings.FieldsFunc(text, f) //strings.FieldsFunc(s, unicode.IsSpace)
    
    counts := make(map[string]int, len(words))
    for _, w := range words {
        counts[w]++
    }
    
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
