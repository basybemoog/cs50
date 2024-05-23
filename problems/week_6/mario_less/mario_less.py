def main():
    value = height()
    conter =value
    for up in range(1,conter+1):
        print(" "* (value-up)+"#" * up)
def height():
    while True:
        try:
            var = int(input("Heigth: "))
            if var > 0 and var < 9:
                return var
        except ValueError:
            print("Not an integer")

main()