package hw2

import (
	"errors"
	"strconv"
	"strings"
	"unicode/utf8"

	valid "github.com/asaskevich/govalidator"
)

// Unpack unpacks the string according to the format string
func Unpack(str string) (string, error) {
	var err error

	if !utf8.ValidString(str) {
		err = errors.New("It is not valid utf8 string")
		return "", err
	}

	runesCount := utf8.RuneCountInString(str)
	if runesCount == 0 {
		return "", err
	}

	var unpackedStrBuilder strings.Builder

	var repeatCountStrBuilder strings.Builder
	repeatRune := rune(-1)
	repeatCount := 1

	escapedRune := '\\'
	isEscapeLastRune := false

	runeIndex := 0
	strRune := ""
	for _, rune := range str {
		runeIndex++
		strRune = string(rune)

		// first rune cannot be digit
		if runeIndex == 1 && valid.IsInt(strRune) {
			err = errors.New("Invalid string: First rune cannot be digit")
			return "", err
		}
		// if escaped char locates before lette
		if isEscapeLastRune && !valid.IsInt(strRune) && rune != escapedRune {
			err = errors.New("Invalid string: You cannot escape letter")
			return "", err
		}

		if valid.IsInt(strRune) && !isEscapeLastRune { // if rune is digit (not escape)
			repeatCountStrBuilder.WriteRune(rune)
			continue
		} else {
			// if rune is "escape" character
			if rune == escapedRune && !isEscapeLastRune {
				// last rune cannot be `\`
				if runeIndex == runesCount {
					err = errors.New("Invalid string: Last rune cannot be '\\' ")
					return "", err
				}

				isEscapeLastRune = true
			} else {
				// if rune is char: repeat previous rune
				if repeatRune != -1 {
					if repeatCountStrBuilder.Len() > 0 {
						repeatCount, err = strconv.Atoi(repeatCountStrBuilder.String())
						if err != nil {
							return "", err
						}
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
			repeatCount, err = strconv.Atoi(repeatCountStrBuilder.String())
			if err != nil {
				return "", err
			}
		} else {
			repeatCount = 1
		}
		unpackedStrBuilder.WriteString(strings.Repeat(string(repeatRune), repeatCount))
	}

	if err != nil {
		unpackedStrBuilder.Reset()
	}

	return unpackedStrBuilder.String(), err
}
