package readability

import (
	"math"
	"strconv"
)

func Readability(text string) string {
	var length = len(text)

	var wordCounter float64 = 0      // количество слов
	var letterCounter float64 = 0    // количество букв
	var sentencesCounter float64 = 0 // количество предложений

	i := 0

	for i < length { // данный цикл необходим для подсчитывания букв, слов и предложений в тексте

		if (text[i] >= 'a' && text[i] <= 'z') || (text[i] >= 'A' && text[i] <= 'Z') {

			letterCounter++ // отсчитывает буквы

		} else if text[i] == ' ' {

			wordCounter++ // при помощи пробела отсчитывает слова

		} else if text[i] == '.' || text[i] == '!' || text[i] == '?' {

			sentencesCounter++ // при помощи знаков препинания отсчитывает предложения

		}
		if length-1 == i && (text[i] == '.' || text[i] == '!' || text[i] == '?') {
			wordCounter++ // при помощи данного условия отсчитывает последнее слово в тексте
		}
		i++
	}
	averageLetter := (letterCounter / wordCounter) * 100                   //  среднее количество букв на 100 слов в тексте
	averageSentences := (sentencesCounter / wordCounter) * 100             // среднее количество предложений в тексте
	var idx = (0.0588 * averageLetter) - (0.296 * averageSentences) - 15.8 // Данная формула узнаёт к какому уровню сложности относится тексте
	if idx < 1 {
		return "Before Grade 1" // Если уровень сложности текста меньше первого класса
	} else if idx > 16 {
		return "Grade 16+" // если уровень сложности текста выше 16-го класса
	}
	var grade = int(math.Round(idx)) // Так как значение даётся с частным, его необходимо округлить для точной оценки класса

	class := "Grade " + strconv.Itoa(grade) // Строка, которая будет возвращать уровень класса

	return class
}
