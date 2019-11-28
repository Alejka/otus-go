package hw4

import (
    "fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLen(t *testing.T) {
    var l = List{}
    
    require.Equal(t, 0, l.Len(), "Testing a length of the list: " + fmt.Sprintf("%v", l))
    
    l.PushBack(1)
    l.PushBack(2)
    l.PushBack(3)
    
    require.Equal(t, 3, l.Len(), "Testing a length of the list: " + fmt.Sprintf("%v", l))
}

func TestFirst(t *testing.T) {
    var l = List{}
    
    l.PushBack(1)
    l.PushBack(2)
    l.PushBack(3)
    
    require.Equal(t, 1, l.First().Value(), "Testing a length of the list: " + fmt.Sprintf("%v", l))
}

func TestLast(t *testing.T) {
    var l = List{}
    
    l.PushBack(1)
    l.PushBack(2)
    l.PushBack(3)
    
    require.Equal(t, 3, l.Last().Value(), "Testing a length of the list: " + fmt.Sprintf("%v", l))
}

func TestPushFront(t *testing.T) {
    var l = List{}
    
    l.PushBack(1)
    l.PushBack(2)
    l.PushBack(3)
    l.PushFront(4)
    
    require.Equal(t, 4, l.First().Value(), "Testing a length of the list: " + fmt.Sprintf("%v", l))
}

func TestPushBack(t *testing.T) {
    var l = List{}
    
    l.PushBack(1)
    l.PushBack(2)
    l.PushBack(3)
    
    require.Equal(t, 3, l.Len(), "Testing a length of the list: " + fmt.Sprintf("%v", l))
}
