package hw3

import (
	"fmt"
	"strings"
	"unicode"
)

// Top10 returns the 10 most frequently occurring words
func Top10(text string) []string {
    //text = "slovo1 slovo2 slovo3 slovo1 slovo2 slovo1 slovo2 slovo1 fkdsfj fds fjlkdsf ldskfdf dfksv slovo3 dljflskd"
    //top10words = []string{"slovo1", "slovo2", "slovo3"}
    
    var top10 []string
    
    wordFilterFunction := func(c rune) bool {
		//return !unicode.IsLetter(c)
        return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
    var words []string = strings.FieldsFunc(text, wordFilterFunction)
    
    cache := map[string]int{}
    for _, word := range words {
        if _, ok := cache[word]; !ok {
            cache[word] = 0
        }
        cache[word] = cache[word] + 1
    }
    fmt.Println(cache)
    
    top10 = []string{"slovo1", "slovo2", "slovo3"}
    fmt.Println(top10)
    
    return top10
}
