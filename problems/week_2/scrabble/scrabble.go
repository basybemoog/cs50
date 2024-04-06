package scrabble

import (
	"unicode"
)

func Scrabble(word1, word2 string) string {

	Player1 := summing(word1) // контрольная сумма от слова первого игрока
	Player2 := summing(word2) // контрольная сумма от слова второга игрока
	if Player1 > Player2 {
		return "Player 1 wins!"
	} else if Player2 > Player1 { // Данное условие проверяет чья контрольная сумма больше
		return "Player 2 wins!" // Условие возвращает строку
	}
	return "Tie!"
}

func summing(word string) int { //объявил функцию, с вводной переменной типа string, на выводе получает переменную типа int
	length1 := len(word) // здесь высчитываю длинну слова
	i := 0               // использую его как счётчик
	var word1 string     // объявляю строку, для того, чтобы перезаписать его в верхний регистр
	for i < length1 {    // цикл for перезаписывает слово в верхний регистр
		word1 += string(unicode.ToUpper(rune(word[i])))
		i++
	}
	i = 0             // сбрасываю счётчик, для повтороного использования
	summ1 := 0        // объявляю  переменную для контрольной суммы, которая будет высчитывать победителя
	for i < length1 { // счётчик, который суммирует значение каждого символа, для выведения контрольной суммы
		if word1[i] == 'A' || word1[i] == 'E' || word1[i] == 'I' || word1[i] == 'L' || word1[i] == 'N' ||
			word1[i] == 'O' || word1[i] == 'R' || word1[i] == 'S' || word1[i] == 'T' || word1[i] == 'U' {

			summ1 += 1 // значение букв A, E, I, L, N, O, R, S, T, U == 1, поэтому он прибавляет к сумме 1

		} else if word1[i] == 'D' || word1[i] == 'G' {

			summ1 += 2 // значение букв D, G == 2, поэтому он прибавляет к сумме 2

		} else if word1[i] == 'B' || word1[i] == 'C' || word1[i] == 'M' || word1[i] == 'P' {

			summ1 += 3 // значение букв B, C, M, P == 3 , поэтому он прибавляет к сумме 3

		} else if word1[i] == 'F' || word1[i] == 'H' || word1[i] == 'V' || word1[i] == 'W' || word1[i] == 'Y' {

			summ1 += 4 // значение букв F, H, V, W, Y == 4 , поэтому он прибавляет к сумме 4

		} else if word1[i] == 'K' {

			summ1 += 5 // Значение буквы К == 5 , прибавляет к сумме 5

		} else if word1[i] == 'J' || word1[i] == 'X' {
			summ1 += 8 // Значение букв J, X == 8, прибавляет к сумме 8
		} else if word1[i] == 'Q' || word1[i] == 'Z' {
			summ1 += 10 // Значение букв Q, Z == 10, прибавляет к сумме 10
		} else {
			summ1 += 0 // Данное условие нужно для того, чтобы любые другие символы по типу знаков препинания игнорировались
		}
		i++ // +1 к счётчику
	}
	return summ1 // возвращает значение контрольной суммы типа int
}
