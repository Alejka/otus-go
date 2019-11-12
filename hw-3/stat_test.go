package hw3

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTop10(t *testing.T) {
    var text string
    var top10words []string
    
    text = "slovo1 slovo2 slovo3 slovo1 slovo2 slovo1 slovo2 slovo1 fkdsfj fds dfksv slovo3 dljflskd sLovo1 sLovo1 ddsa sloVO1"
    //top10words = []string{"slovo1", "slovo2", "slovo3"}
    top10words = []string{"slovo1", "slovo2", "slovo3", "dfksv", "dljflskd", "fkdsfj", "ddsa", "fds"}
    require.Equal(t, top10words, Top10(text), "Analyze a string: "+text)
}
