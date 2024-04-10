package readability

import (
	"math"
	"strconv"
)

func Readability(text string) string {
	var length = len(text)

	var word_counter float64 = 0      // количество слов
	var letter_counter float64 = 0    // количество букв
	var sentences_counter float64 = 0 // количество предложений

	i := 0

	for i < length { // данный цикл необходим для подсчитывания букв, слов и предложений в тексте

		if (text[i] >= 'a' && text[i] <= 'z') || (text[i] >= 'A' && text[i] <= 'Z') {

			letter_counter++ // отсчитывает буквы

		} else if text[i] == ' ' {

			word_counter++ // при помощи пробела отсчитывает слова

		} else if text[i] == '.' || text[i] == '!' || text[i] == '?' {

			sentences_counter++ // при помощи знаков препинания отсчитывает предложения

		}
		if length-1 == i && (text[i] == '.' || text[i] == '!' || text[i] == '?') {
			word_counter++ // при помощи данного условия отсчитывает последнее слово в тексте
		}
		i++
	}
	average_letter := (letter_counter / word_counter) * 100                  //  среднее количество букв на 100 слов в тексте
	average_sentences := (sentences_counter / word_counter) * 100            // среднее количество предложений в тексте
	var idx = (0.0588 * average_letter) - (0.296 * average_sentences) - 15.8 // Данная формула узнаёт к какому уровню сложности относится тексте
	if idx < 1 {
		return "Before Grade 1" // Если уровень сложности текста меньше первого класса
	} else if idx > 16 {
		return "Grade 16+" // если уровень сложности текста выше 16-го класса
	}
	var grade = int(math.Round(idx)) // Так как значение даётся с частным, его необходимо округлить для точной оценки класса

	class := "Grade " + strconv.Itoa(grade) // Строка, которая будет возвращать уровень класса

	return class
}
