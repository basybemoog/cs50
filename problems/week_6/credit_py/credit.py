def main():
    number = get_int("Number: ")
    while number < 1000000000 and number < 5700000000000000:
        number = get_int("Number: ")
    cardlen = len(str(number))
    cardnum = str(number)
    summa = 0
    if cardlen > 12 and cardlen < 17:
        for start in range(1, 17):
            if start % 2 == 0:
                chet = (number % 10) * 2
                if chet >= 10:
                    summa += (chet % 10) + 1
                else:
                    summa += chet
            elif start % 2 != 0:
                summa += number % 10
            number //= 10
    else:
        print("INVALID")
        return
    if summa % 10 != 0:
        print("INVALID")
        return
    if cardlen == 13:
        if int(cardnum[0:2]) in range(40, 50):
            print("VISA")
            return
        else:
            pass
    elif cardlen == 15:
        if int(cardnum[0:2]) == 34 or int(cardnum[0:2]) == 37:
            print("AMEX")
            return
        else:
            pass
    elif cardlen == 16:
        if int(cardnum[0:2]) in range(51, 56):
            print("MASTERCARD")
            return
        elif int(cardnum[0:2]) in range(40, 50):
            print("VISA")
            return
        else:
            pass
    print("INVALID")


def get_int(value):
    while True:
        try:
            return int(input(value))
        except:
            pass


main()
