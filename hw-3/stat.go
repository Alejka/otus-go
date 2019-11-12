package hw3

import (
	"fmt"
	"strings"
	"unicode"
    "sort"
)

// Top10 returns the 10 most frequently occurring words
func Top10(text string) []string {
    //text = "slovo1 slovo2 slovo3 slovo1 slovo2 slovo1 slovo2 slovo1 fkdsfj fds dfksv slovo3 dljflskd sLovo1 sLovo1 ddsa sloVO1"
    //top10words = []string{"slovo1", "slovo2", "slovo3"}
    
    text = strings.ToLower(text)
    
    f := func(c rune) bool {
		//return !unicode.IsLetter(c)
        return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
    words := strings.FieldsFunc(text, f) //strings.FieldsFunc(s, unicode.IsSpace)
    fmt.Println("words:")
    fmt.Println(words)
    fmt.Println(len(words))
    fmt.Println()
    
    counts := make(map[string]int, len(words))
    for _, w := range words {
        counts[w]++
    }
    fmt.Println("counts:")
    fmt.Println(counts)
    fmt.Println(len(counts))
    fmt.Println()
    
    type WordFreq struct {
        word string
        freq int
    }
    wordFreqList := make([]WordFreq, 0, len(counts))
    for k, v := range counts {
        wordFreqList = append(wordFreqList, WordFreq{k, v})
    }
    fmt.Println("wordFreqList:")
    fmt.Println(wordFreqList)
    fmt.Println(len(wordFreqList))
    fmt.Println()
    
    sort.Slice(wordFreqList, func(i, j int) bool {
		return wordFreqList[i].freq > wordFreqList[j].freq
	})
    fmt.Println("wordFreqList sorted:")
    fmt.Println(wordFreqList)
    fmt.Println(len(wordFreqList))
    fmt.Println()
    
    top10map := wordFreqList[:len(wordFreqList)]
    fmt.Println("top10map:")
    fmt.Println(top10map)
    fmt.Println(len(top10map))
    fmt.Println()
    
    top10 := make([]string, 0, 10)
    for _, v := range top10map {
        top10 = append(top10, v.word)
    }
    fmt.Println("top10:")
    fmt.Println(top10)
    fmt.Println()
    
    // todo 1) if count of words is more than 10?
    // todo 2) sort.Slice() returns random result
    // todo 3) add more tests
    
    return top10
}
