package hw2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	// "a4bc2d5e" => "aaaabccddddde"
	// "abcd" => "abcd"
	// "45" => "" (некорректная строка)

	// qwe\4\5 => qwe45 (*)
	// qwe\45 => qwe44444 (*)
	// qwe\\5 => qwe\\\\\ (*)

	var str string

	str = "a4bc2d5e"
	require.Equal(t, "aaaabccddddde", Unpack(str), "Unpacking string: "+str)

	str = "a14bc2d5e"
	require.Equal(t, "aaaaaaaaaaaaaabccddddde", Unpack(str), "Unpacking string: "+str)

	str = "abcd"
	require.Equal(t, "abcd", Unpack(str), "Unpacking string: "+str)

	str = "45"
	require.Equal(t, "", Unpack(str), "Unpacking string: "+str)

	str = "45aa2c"
	require.Equal(t, "aaac", Unpack(str), "Unpacking string: "+str)

	str = ""
	require.Equal(t, "", Unpack(str), "Unpacking string: "+str)
}
