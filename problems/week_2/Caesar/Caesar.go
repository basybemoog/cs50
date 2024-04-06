package Caesar

import (
	"fmt"
)

func Caesar(key int) string {
	if key < 0 {
		return "Usage: ./caesar key"
	}
	var plain_t []string
	fmt.Printf("plain text: ")
	fmt.Scanf("%s", &plain_t[0])
	length_of_text := len(plain_t)
	var cipher_text string
	for i := 0; i < length_of_text; i++ {
		if plain_t[0][i] >= 'A' && plain_t[0][i] <= 'Z' {

		} else if plain_t[0][i] >= 'a' && plain_t[0][i] <= 'z' {

		}
	}
}
