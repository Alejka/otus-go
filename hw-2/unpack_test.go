package hw2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	var str string
	var unpackedStr string
	var err error

	str = `qwe\4\5`
	unpackedStr, err = Unpack(str)
	require.Equal(t, "qwe45", unpackedStr, "Unpacking a string: "+str)
	require.Nil(t, err, "Unpacking a string: "+str)

	str = `qwe\45`
	unpackedStr, err = Unpack(str)
	require.Equal(t, "qwe44444", unpackedStr, "Unpacking a string: "+str)
	require.Nil(t, err, "Unpacking a string: "+str)

	str = `qwe\\5`
	unpackedStr, err = Unpack(str)
	require.Equal(t, `qwe\\\\\`, unpackedStr, "Unpacking a string: "+str)
	require.Nil(t, err, "Unpacking a string: "+str)

	str = `\45`
	unpackedStr, err = Unpack(str)
	require.Equal(t, "44444", unpackedStr, "Unpacking a string: "+str)
	require.Nil(t, err, "Unpacking a string: "+str)

	str = `\\`
	unpackedStr, err = Unpack(str)
	require.Equal(t, `\`, unpackedStr, "Unpacking a string: "+str)
	require.Nil(t, err, "Unpacking a string: "+str)

	str = `\7`
	unpackedStr, err = Unpack(str)
	require.Equal(t, "7", unpackedStr, "Unpacking a string: "+str)
	require.Nil(t, err, "Unpacking a string: "+str)

	str = "a4b11c2"
	unpackedStr, err = Unpack(str)
	require.Equal(t, "aaaabbbbbbbbbbbcc", unpackedStr, "Unpacking a string: "+str)
	require.Nil(t, err, "Unpacking a string: "+str)

	str = "a4b0c0"
	unpackedStr, err = Unpack(str)
	require.Equal(t, "aaaa", unpackedStr, "Unpacking a string: "+str)
	require.Nil(t, err, "Unpacking a string: "+str)

	str = "a4bc2d5e"
	unpackedStr, err = Unpack(str)
	require.Equal(t, "aaaabccddddde", unpackedStr, "Unpacking a string: "+str)
	require.Nil(t, err, "Unpacking a string: "+str)

	str = "a14bc2d5e"
	unpackedStr, err = Unpack(str)
	require.Equal(t, "aaaaaaaaaaaaaabccddddde", unpackedStr, "Unpacking a string: "+str)
	require.Nil(t, err, "Unpacking a string: "+str)

	str = "abcd"
	unpackedStr, err = Unpack(str)
	require.Equal(t, "abcd", unpackedStr, "Unpacking a string: "+str)
	require.Nil(t, err, "Unpacking a string: "+str)

	str = "a"
	unpackedStr, err = Unpack(str)
	require.Equal(t, "a", unpackedStr, "Unpacking a string: "+str)
	require.Nil(t, err, "Unpacking a string: "+str)

	str = ""
	unpackedStr, err = Unpack(str)
	require.Equal(t, "", unpackedStr, "Unpacking a string: "+str)
	require.Nil(t, err, "Unpacking a string: "+str)

	str = "世3界é2co"
	unpackedStr, err = Unpack(str)
	require.Equal(t, "世世世界ééco", unpackedStr, "Unpacking a string: "+str)
	require.Nil(t, err, "Unpacking a string: "+str)

	// errors

	str = `\`
	unpackedStr, err = Unpack(str)
	require.Equal(t, "", unpackedStr, "Unpacking a string: "+str)
	require.NotNil(t, err, "Unpacking a string: "+str)

	str = `\a`
	unpackedStr, err = Unpack(str)
	require.Equal(t, "", unpackedStr, "Unpacking a string: "+str)
	require.NotNil(t, err, "Unpacking a string: "+str)

	str = `g2\a`
	unpackedStr, err = Unpack(str)
	require.Equal(t, "", unpackedStr, "Unpacking a string: "+str)
	require.NotNil(t, err, "Unpacking a string: "+str)

	str = `d4\`
	unpackedStr, err = Unpack(str)
	require.Equal(t, "", unpackedStr, "Unpacking a string: "+str)
	require.NotNil(t, err, "Unpacking a string: "+str)

	str = "45aa2c"
	unpackedStr, err = Unpack(str)
	require.Equal(t, "", unpackedStr, "Unpacking a string: "+str)
	require.NotNil(t, err, "Unpacking a string: "+str)
}
