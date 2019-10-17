package hw2

import (
	"fmt"
	"strconv"
	"unicode"
	"unicode/utf8"
)

func Unpack(str string) string {
	unpackedStr := ""

	if !utf8.ValidString(str) || utf8.RuneCountInString(str) == 0 {
		return unpackedStr
	}

	lastRune := rune(-1)
	lastCharRune := rune(-1)
	repeatCountStr := "0"
	repeatCount := 0

	// a4bc2d5e => aaaabccddddde
	fmt.Printf("-----------\nNEXT STR: %q\n-----------\n", str)
	for _, rune := range str {
		fmt.Printf("rune: %q\n", rune)

		if unicode.IsDigit(rune) { // if digit
			repeatCountStr += string(rune)
		} else { // if char
			if lastCharRune != -1 {
				repeatCount, _ = strconv.Atoi(repeatCountStr)
				if repeatCount > 0 {
					unpackedStr += repeatRunes(lastCharRune, repeatCount)
				} else {
					unpackedStr += string(lastCharRune)
				}
			}

			lastCharRune = rune
			repeatCountStr = "0"
		}

		lastRune = rune
		fmt.Printf("unpackedStr: %q\n-----------\n", unpackedStr)
	}

	fmt.Printf("lastRune: %q\n-----------\n", lastRune)
	if lastRune != -1 && !unicode.IsDigit(lastRune) { //@todo remove condition: lastRune != 1
		unpackedStr += string(lastCharRune)
	}
	fmt.Printf("unpackedStr: %q\n-----------\n", unpackedStr)

	return unpackedStr
}

func repeatRunes(rune rune, count int) string {
	str := ""

	if count < 1 {
		return str
	}

	for i := 1; i <= count; i++ {
		str += string(rune)
	}

	return str
}
