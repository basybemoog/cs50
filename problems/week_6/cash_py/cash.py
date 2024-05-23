def main():
    value = get_float()
    value = round(value*100)
    counter = 0
    while value >=25:
        value -= 25
        counter+=1
    while value >=10:
        value -= 10
        counter+=1
    while value >=5:
        value -= 5
        counter+=1
    while value >=1:
        value-=1
        counter+=1

    print(counter)

def get_float():
    while True:
        try:
            n = float(input('Change owed: '))
            if n > 0:
                return n
        except ValueError:
            print('Not an integer')

main()