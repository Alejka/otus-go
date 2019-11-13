package hw3

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTop10(t *testing.T) {
    var text string
    var top []string
    var expected []string
    var actual []string
    
    text = "slovo1 slovo2 slovo3 slovo1 slovo2 slovo1 slovo2 slovo1 fkdsfj fds dfksv slovo3 dljflskd sLovo1 sLovo1 ddsa sloVO1"
    expected = []string{"slovo1", "slovo2", "slovo3"}
    top = Top10(text)
    actual = top[:3]
    require.Equal(t, expected, actual, "Analyze a string: "+text)
    
    text = "slovo1 slovo2 slovo3 slovo1 slovo2 slovo1 slovo2 slovo1 fkdsfj fds dfksv slovo3 dljflskd sLovo1 sLovo1 ddsa sloVO1 645ппавп павпва п вапвапа  пвап вп"
    expected = []string{"slovo1", "slovo2", "slovo3"}
    top = Top10(text)
    actual = top[:3]
    require.Equal(t, expected, actual, "Analyze a string: "+text)
}
