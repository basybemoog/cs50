package scrabble

import (
	"cs50/problems/week_2/scrabble"
	"testing"
)

const TIE = "Tie!"
const Player1Win = "Player 1 wins!"
const Player2Win = "Player 2 wins!"

func TestScrabble(t *testing.T) {
	tests := []struct {
		name       string
		wordFirst  string
		wordSecond string
		want       string
	}{
		{"handles letter cases correctly", "LETTERCASE", "lettercase", TIE},
		{"handles punctuation correctly", "Punctuation!?!?", "punctuation", TIE},
		{"correctly identifies 'Question?' and 'Question!' as a tie", "Question?", "Question!", TIE},
		{"correctly identifies 'drawing' and 'illustration' as a tie", "drawing", "illustration", TIE},
		{"correctly identifies 'hai!' as winner over 'Oh,'", "Oh,", "hai!", Player2Win},
		{"correctly identifies 'COMPUTER' as winner over 'science'", "COMPUTER", "science", Player1Win},
		{"correctly identifies 'Scrabble' as winner over 'wiNNeR'", "Scrabble", "wiNNeR", Player1Win},
		{"correctly identifies 'pig' as winner over 'dog'", "pig", "dog", Player1Win},
		{"correctly identifies 'Skating!' as winner over 'figure?'", "figure?", "Skating!", Player2Win},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scrabble.Scrabble(tt.wordFirst, tt.wordSecond); got != tt.want {
				t.Errorf("Scrabble() = %v, want %v", got, tt.want)
			}
		})
	}
}
