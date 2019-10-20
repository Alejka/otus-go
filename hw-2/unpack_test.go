package hw2

import (
	"testing"
	//"strconv"

	"github.com/stretchr/testify/require"
)

/*func TestRepeatRunes(t *testing.T) {
    var r rune
    var repeatCount int

    r = 'a'
    repeatCount = 0
    require.Equal(t, "", repeatRunes(r, repeatCount), "Repeating a rune '"+string(r)+"', "+strconv.Itoa(repeatCount)+" times")

    r = 'o'
    repeatCount = -7
    require.Equal(t, "", repeatRunes(r, repeatCount), "Repeating a rune '"+string(r)+"', "+strconv.Itoa(repeatCount)+" times")

    r = 'z'
    repeatCount = 1
    require.Equal(t, "z", repeatRunes(r, repeatCount), "Repeating a rune '"+string(r)+"', "+strconv.Itoa(repeatCount)+" times")

    r = 'q'
    repeatCount = 5
    require.Equal(t, "qqqqq", repeatRunes(r, repeatCount), "Repeating a rune '"+string(r)+"', "+strconv.Itoa(repeatCount)+" times")
}*/

func TestUnpack(t *testing.T) {
	// "a4bc2d5e" => "aaaabccddddde"
	// "abcd" => "abcd"
	// "45" => "" (некорректная строка)

	// qwe\4\5 => qwe45 (*)
	// qwe\45 => qwe44444 (*)
	// qwe\\5 => qwe\\\\\ (*)

	var str string

	str = `qwe\4\5`
	require.Equal(t, "qwe45", Unpack(str), "Unpacking a string: "+str)

	str = `qwe\45`
	require.Equal(t, "qwe44444", Unpack(str), "Unpacking a string: "+str)

	str = `qwe\\5`
	require.Equal(t, `qwe\\\\\`, Unpack(str), "Unpacking a string: "+str)

	str = "a4b11c2"
	require.Equal(t, "aaaabbbbbbbbbbbcc", Unpack(str), "Unpacking a string: "+str)

	str = "a4b0c0"
	require.Equal(t, "aaaa", Unpack(str), "Unpacking a string: "+str)

	str = "a4bc2d5e"
	require.Equal(t, "aaaabccddddde", Unpack(str), "Unpacking a string: "+str)

	str = "a14bc2d5e"
	require.Equal(t, "aaaaaaaaaaaaaabccddddde", Unpack(str), "Unpacking a string: "+str)

	str = "abcd"
	require.Equal(t, "abcd", Unpack(str), "Unpacking a string: "+str)

	str = ""
	require.Equal(t, "", Unpack(str), "Unpacking a string: "+str)

	str = "45"
	require.Equal(t, "", Unpack(str), "Unpacking a string: "+str)

	str = "45aa2c"
	require.Equal(t, "aaac", Unpack(str), "Unpacking a string: "+str)

	str = "世3界é2co"
	require.Equal(t, "世世世界ééco", Unpack(str), "Unpacking a string: "+str)
}
