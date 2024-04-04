package credit

import (
	"strconv"
)

func Validate(card_number int) string {

	delitelb := 10
	modul := 10
	place := 1
	chastnoe := 0
	mnozit := 10
	summchet := 0
	summnechet := 0

	var answer string

	var card_valid bool

	length := len(strconv.Itoa(card_number))

	for place <= length {
		if place == 1 {
			chastnoe = card_number % modul
			summnechet += chastnoe
			place++
			modul *= mnozit
		}
		if place%2 == 0 {
			chastnoe = (card_number % modul) / delitelb
			chastnoe *= 2
			if chastnoe > 10 {
				chastnoe = (chastnoe % 10) + 1
				summchet += chastnoe
			} else if chastnoe == 10 {
				chastnoe = 1
				summchet += chastnoe
			} else {
				summchet += chastnoe
			}
			place++
			modul *= mnozit
			delitelb *= mnozit
		} else {
			chastnoe = (card_number % modul) / delitelb
			summnechet += chastnoe
			place++
			modul *= mnozit
			delitelb *= mnozit
		}
	}

	end_sum := summchet + summnechet

	if end_sum%10 == 0 {
		card_valid = true
	} else {
		card_valid = false
		return "INVALID"
	}
	if card_valid == true {
		lenVISA1 := 1000000000000
		lenVISA2 := 1000000000000000
		lenAMEX := 10000000000000
		lenMC := 100000000000000

		if card_number/lenVISA1 == 4 {
			answer = "VISA"
		} else if card_number/lenVISA2 == 4 {
			answer = "VISA"
		} else if card_number/lenAMEX == 34 || card_number/lenAMEX == 37 {
			answer = "AMEX"
		} else if card_number/lenMC >= 51 && card_number/lenMC <= 55 {
			answer = "MASTERCARD"
		} else {
			answer = "INVALID"
		}

	}
	return answer
}
