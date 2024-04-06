package scrabble

import (
	"cs50/problems/week_2/scrabble"
	"testing"
)

func TestScrabble(t *testing.T) {
	tests := []struct {
		name       string
		wordFirst  string
		wordSecond string
		want       string
	}{
		{"handles letter cases correctly", "LETTERCASE", "lettercase", "Tie!"},
		{"handles punctuation correctly", "Punctuation!?!?", "punctuation", "Tie!"},
		{"correctly identifies 'Question?' and 'Question!' as a tie", "Question?", "Question!", "Tie!"},
		{"correctly identifies 'drawing' and 'illustration' as a tie", "drawing", "illustration", "Tie!"},
		{"correctly identifies 'hai!' as winner over 'Oh,'", "Oh,", "hai!", "Player 2 wins!"},
		{"correctly identifies 'COMPUTER' as winner over 'science'", "COMPUTER", "science", "Player 1 wins!"},
		{"correctly identifies 'Scrabble' as winner over 'wiNNeR'", "Scrabble", "wiNNeR", "Player 1 wins!"},
		{"correctly identifies 'pig' as winner over 'dog'", "pig", "dog", "Player 1 wins!"},
		{"correctly identifies 'Skating!' as winner over 'figure?'", "figure?", "Skating!", "Player 2 wins!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scrabble.Scrabble(tt.wordFirst, tt.wordSecond); got != tt.want {
				t.Errorf("Scrabble() = %v, want %v", got, tt.want)
			}
		})
	}
}
