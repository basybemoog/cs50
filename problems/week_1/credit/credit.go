package credit

func Validate(cardNum int) string {

	var cardNumReserve = cardNum

	answer := 0

	counter := 0

	for cardNum > 0 {
		temp := cardNum % 10

		if counter%2 == 0 {
			answer += temp
		} else {
			temp *= 2

			if temp > 9 {
				answer += temp % 10
				answer += temp / 10
			} else {
				answer += temp
			}
		}
		counter++
		cardNum /= 10
	}

	if answer%10 != 0 {
		return "INVALID"
	}

	lenVISA1 := 1000000000000    // длина карты VISA в 13 чисел
	lenVISA2 := 1000000000000000 // длина карты VISA в 16 чисел
	lenAMEX := 10000000000000    // длина карты AMEX в 15 чисел
	lenMC := 100000000000000     // длина карты MASTERCARD в 16 чисел

	if cardNumReserve/lenVISA1 == 4 || cardNumReserve/lenVISA2 == 4 {
		return "VISA" // если при делении на VISA1 число равно 4, то это карта VISA
	} else if cardNumReserve/lenAMEX == 34 || cardNumReserve/lenAMEX == 37 {
		return "AMEX" // если при делении на AMEX число равно 34 или 37, то это карта AMEX
	} else if cardNumReserve/lenMC >= 51 && cardNumReserve/lenMC <= 55 {
		return "MASTERCARD" // если при делении на MC число в диапазоне от 51 до 55, то это MASTERCARD
	}
	return "INVALID" // В иных случаях, карта не проходит валидацию
}
