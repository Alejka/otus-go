package hw2

import (
	"fmt"
	"strconv"
	"strings"

	//"unicode"
	"unicode/utf8"

	valid "github.com/asaskevich/govalidator"
)

func Unpack(str string) string {
	// "a4bc2d5e" => "aaaabccddddde"
	// "abcd" => "abcd"
	// "45" => "" (некорректная строка)

	// qwe\4\5 => qwe45 (*)
	// qwe\45 => qwe44444 (*)
	// qwe\\5 => qwe\\\\\ (*)

	// qwe\5 => qwe5 (*)
	// qwe\\5 => qwe\\\\\ (*)

	runesCount := utf8.RuneCountInString(str)

	if !utf8.ValidString(str) || runesCount == 0 {
		return ""
	}

	var unpackedStrBuilder strings.Builder
	var repeatCountStrBuilder strings.Builder
	repeatRune := rune(-1)
	repeatCount := 1
	runeIndex := 0
	isEscapedLastRune := false

	// a4bc2d5e => aaaabccddddde
	fmt.Printf("-----------\nNEXT STR: %q\nCount of runes: %d\n-----------\n", str, runesCount)
	for _, rune := range str {
		runeIndex++
		fmt.Printf("runeIndex: %d\nrune: %q\n", runeIndex, rune)

		// https://stackoverflow.com/questions/25540951/differences-between-isdigit-and-isnumber-in-unicode-in-go
		// http://www.fileformat.info/info/unicode/category/Nd/list.htm
		//if unicode.IsDigit(rune) {
		if valid.IsInt(string(rune)) && !isEscapedLastRune { // if rune is digit
			repeatCountStrBuilder.WriteRune(rune)
			continue
		} else {
			// if rune is "escape" character
			if rune == '\\' && !isEscapedLastRune {
				isEscapedLastRune = true
			} else {
				// if rune is char: repeat previous rune
				if repeatRune != -1 {
					if repeatCountStrBuilder.Len() > 0 {
						repeatCount, _ = strconv.Atoi(repeatCountStrBuilder.String())
					} else {
						repeatCount = 1
					}
					unpackedStrBuilder.WriteString(strings.Repeat(string(repeatRune), repeatCount))
				}

				repeatRune = rune
				repeatCountStrBuilder.Reset()
				isEscapedLastRune = false
			}
		}

		/*
		   if runeIndex == runesCount {
		       if repeatRune != -1 {
		           if repeatCountStrBuilder.Len() > 0 {
		               repeatCount, _ = strconv.Atoi(repeatCountStrBuilder.String())
		           } else {
		               repeatCount = 1
		           }
		           unpackedStrBuilder.WriteString(strings.Repeat(string(repeatRune), repeatCount))
		       }
		   }
		*/
	}

	fmt.Printf("process last rune: \n")
	if repeatRune != -1 {
		if repeatCountStrBuilder.Len() > 0 {
			repeatCount, _ = strconv.Atoi(repeatCountStrBuilder.String())
		} else {
			repeatCount = 1
		}
		fmt.Printf("repeatCount: %d\nrepeatRune: %q\n", repeatCount, repeatRune)
		unpackedStrBuilder.WriteString(strings.Repeat(string(repeatRune), repeatCount))
	}

	tmpstr := unpackedStrBuilder.String()
	fmt.Printf("READY STR: %q\n", tmpstr)
	return tmpstr
	//return unpackedStrBuilder.String()
}
