package substitution

import (
	"strings"
)

func ChipText(r rune, chip string) rune {
	char := r
	alph := "abcdefghijklmnopqrstuvwxyz"
	bigalph := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	j := 0
	if r >= 'a' && r <= 'z' {
		if char == rune(alph[j]) {
			if chip[j] >= 'A' && chip[j] <= 'Z' {
				char = rune(chip[j] + 32)
				return char
			} else {
				char = rune(chip[j])
				return char
			}

		} else {
			for char != rune(alph[j]) {
				j++
			}
			if chip[j] >= 'A' && chip[j] <= 'Z' {
				char = rune(chip[j] + 32)
				return char
			} else {
				char = rune(chip[j])
				return char
			}
		}
	} else if r >= 'A' && r <= 'Z' {

		if char == rune(bigalph[j]) {
			if chip[j] >= 'A' && chip[j] <= 'Z' {
				char = rune(chip[j])
				return char
			} else {
				char = rune(chip[j] - 32)
			}
		} else {
			for char != rune(bigalph[j]) {
				j++
			}
			if chip[j] >= 'A' && chip[j] <= 'Z' {
				char = rune(chip[j])
				return char
			} else {
				char = rune(chip[j] - 32)
				return char
			}
		}

	}
	return r
}
func Substitution(key string, word string) string {
	if len(key) != 26 {
		return "Key must contain 26 characters"
	}
	for i := 0; i < len(key); i++ {
		if ((key[i] >= 'a' && key[i] <= 'z') || (key[i] >= 'A' && key[i] <= 'Z')) == false {
			return "Invalid characters in key"
		}
	}
	for j := 0; j < len(key); j++ {
		k := 0
		for k < len(key) {
			if key[j] == key[k] && k != j {
				return "duplicate characters"
			}
			k++
		}
		k = 0
	}
	chippedText := "ciphertext: " + strings.Map(func(r rune) rune {
		return ChipText(r, key)
	}, word)

	return chippedText
}
