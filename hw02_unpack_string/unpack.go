package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(srcString string) (string, error) {
	var builder strings.Builder

	if srcString == "" {
		return "", nil
	}

	isCorrect, err := verify(srcString)
	if err != nil {
		return "", err
	}

	if isCorrect {
		runes := []rune(srcString)
		size := len(runes)
		for i := 0; i < size-1; i++ {
			if !isDigit(runes[i+1]) && !isDigit(runes[i]) {
				builder.WriteRune(runes[i])
			}
			if !isDigit(runes[i]) && isDigit(runes[i+1]) {
				repeatCount, _ := strconv.Atoi(string(runes[i+1]))
				repeatRuneWriter(runes[i], repeatCount, &builder)
			}
		}
		if !isDigit(runes[size-1]) {
			builder.WriteRune(runes[size-1])
		}
	}

	return builder.String(), nil
}

func verify(srcString string) (bool, error) {
	runes := []rune(srcString)
	if isDigit(runes[0]) {
		return false, ErrInvalidString
	}

	for i := 1; i < len(runes); i++ {
		if isDigit(runes[i]) && isDigit(runes[i-1]) {
			return false, ErrInvalidString
		}
	}

	return true, nil
}

func repeatRuneWriter(r rune, repeatCount int, builder *strings.Builder) {
	if repeatCount != 0 {
		builder.WriteString(strings.Repeat(string(r), repeatCount))
	}
}

func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}
