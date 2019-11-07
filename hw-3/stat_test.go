package hw3

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTop10(t *testing.T) {
    var text string
    var top10words []string
    
    text = "slovo1 slovo2 slovo3 slovo1 slovo2 slovo1 slovo2 slovo1 fkdsfj fds fjlkdsf ldskfdf dfksv slovo3 dljflskd"
    top10words = []string{"slovo1", "slovo2", "slovo3"}
    require.Equal(t, top10words, Top10(text), "Analyze a string: "+text)
}
