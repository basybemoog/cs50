package caesar

import "strings"

func Cipher(r rune, key int) rune {
	char := int(r) + key
	if r >= 'a' && r <= 'z' {

		if char > 'z' {
			return rune(char - 26)
		} else if char < 'a' {
			return rune(char + 26)
		}
	} else if r >= 'A' && r <= 'Z' {

		if char > 'Z' {
			return rune(char - 26)
		} else if char < 'A' {
			return rune(char + 26)
		}
	} else {
		return r
	}
	return rune(char)
}
func Caesar(plaintText string, key int) string {
	for key > 26 {
		key -= 26
	}
	chipperText := "ciphertext: " + strings.Map(func(r rune) rune {
		return Cipher(r, key)
	}, plaintText)
	return chipperText
}
