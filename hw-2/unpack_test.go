package hw2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
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

	str = "5"
	require.Equal(t, "", Unpack(str), "Unpacking a string: "+str)

	str = "a"
	require.Equal(t, "a", Unpack(str), "Unpacking a string: "+str)

	str = ""
	require.Equal(t, "", Unpack(str), "Unpacking a string: "+str)

	str = "45"
	require.Equal(t, "", Unpack(str), "Unpacking a string: "+str)

	str = "45aa2c"
	require.Equal(t, "aaac", Unpack(str), "Unpacking a string: "+str)

	str = "世3界é2co"
	require.Equal(t, "世世世界ééco", Unpack(str), "Unpacking a string: "+str)
}
