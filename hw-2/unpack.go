package hw2

import (
	"strconv"
	"strings"
	"unicode/utf8"

	valid "github.com/asaskevich/govalidator"
)

func Unpack(str string) string {
	runesCount := utf8.RuneCountInString(str)

	if !utf8.ValidString(str) || runesCount == 0 {
		return ""
	}

	var unpackedStrBuilder strings.Builder

	var repeatCountStrBuilder strings.Builder
	repeatRune := rune(-1)
	repeatCount := 1

	isEscapeLastRune := false

	for _, rune := range str {
		if valid.IsInt(string(rune)) && !isEscapeLastRune { // if rune is digit (not escape)
			repeatCountStrBuilder.WriteRune(rune)
			continue
		} else {
			// if rune is "escape" character
			if rune == '\\' && !isEscapeLastRune {
				isEscapeLastRune = true
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
				isEscapeLastRune = false
			}
		}
	}

	// process last rune
	if repeatRune != -1 {
		if repeatCountStrBuilder.Len() > 0 {
			repeatCount, _ = strconv.Atoi(repeatCountStrBuilder.String())
		} else {
			repeatCount = 1
		}
		unpackedStrBuilder.WriteString(strings.Repeat(string(repeatRune), repeatCount))
	}

	return unpackedStrBuilder.String()
}
