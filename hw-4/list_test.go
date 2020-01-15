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
    l.PushBack("last")
    
    require.Equal(t, "last", l.Last().Value(), "Testing a length of the list: " + fmt.Sprintf("%v", l))
}

func TestPushFront(t *testing.T) {
    var l = List{}
    
    l.PushBack(1)
    l.PushBack(2)
    l.PushBack(3)
    l.PushFront("first")
    
    require.Equal(t, "first", l.First().Value(), "Testing a length of the list: " + fmt.Sprintf("%v", l))
    require.Equal(t, 4, l.Len(), "Testing a length of the list: " + fmt.Sprintf("%v", l))
}

func TestPushBack(t *testing.T) {
    var l = List{}
    
    l.PushBack(1)
    l.PushBack(2)
    l.PushBack(3)
    
    require.Equal(t, 3, l.Len(), "Testing a length of the list: " + fmt.Sprintf("%v", l))
}

func TestRemove(t *testing.T) {
    var l = List{}
    
    l.PushBack(1)
    l.PushBack(2)
    l.PushBack(3)
    
    //i := l.Last()
    //l.Remove(i.Prev())
    
    //require.Equal(t, 2, l.Len(), "Testing a length of the list: " + fmt.Sprintf("%v", l))
}
