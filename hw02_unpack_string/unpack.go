package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

const (
	// ServiceSymbol ...
	ServiceSymbol = '\\'
)

// ErrInvalidString ...
var ErrInvalidString = errors.New("invalid string")

func isServiceSymbol(ra []rune, index *int) bool {
	if ra[*index] == rune(ServiceSymbol) &&
		(unicode.IsDigit(ra[*index+1]) || ra[*index+1] == rune(ServiceSymbol)) {
		*index++
		return true
	}

	return false
}

// Unpack ...
func Unpack(str string) (string, error) {
	var resultStr strings.Builder
	ra := []rune(str)
	for i := 0; i < len(ra); i++ {
		if !isServiceSymbol(ra, &i) && unicode.IsDigit(ra[i]) {
			return "", ErrInvalidString
		}
		if i < len(ra)-1 {
			if unicode.IsDigit(ra[i+1]) {
				number, _ := strconv.Atoi(string(ra[i+1]))
				for j := 0; j < number; j++ {
					resultStr.WriteRune(ra[i])
				}
				i++
				continue
			}
		}
		resultStr.WriteRune(ra[i])
	}

	return resultStr.String(), nil
}
