package Substitution

import (
	"strings"
)

func ChipText(r rune, chip string) rune {
	if len(chip) != 26 {
		return 0
	}
	char := r
	alph := "abcdefghijklmnopqrstuvwxyz"
	bigalph := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	j := 0
	if r >= 'a' && r <= 'z' {

		if char == rune(alph[j]) {
			char = rune(chip[j] + 32)
			j = 0
		} else {
			for char != rune(alph[j]) {
				j++
			}
			char = rune(chip[j] + 32)
			j = 0
		}

	} else if r >= 'A' && r <= 'Z' {

		if char == rune(bigalph[j]) {
			char = rune(chip[j])
		} else {
			for char != rune(bigalph[j]) {
				j++
			}
			char = rune(chip[j])
			j = 0
		}

	} else {
		return r

	}
	return char
}
func Substitution(key string, word string) string {
	if len(key) != 26 {
		return "Key must contain 26 characters"
	} else {
		chippedText := strings.Map(func(r rune) rune {
			return ChipText(r, key)
		}, word)

		return chippedText
	}
}
