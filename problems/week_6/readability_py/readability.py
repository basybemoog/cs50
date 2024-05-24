import math


def main():
    text = get_string("Text:")
    text = text.upper()
    word = 0
    letter = 0
    sentence = 0
    for i in range(len(text)):
        if text[i] >= 'A' and text[i] < 'Z':
            letter += 1
        elif text[i] == ' ':
            word += 1
        elif text[i] == '.' or text[i] == '!' or text[i] == '?':
            sentence += 1
        if i == len(text) - 1 and (text[i] == '.' or text[i] == '!' or text[i] == '?'):
            word += 1
    averageL = (float(letter) / word) * 100
    averageS = (float(sentence) / word) * 100
    idx = round((0.0588 * averageL) - (0.296 * averageS) - 15.8)
    if idx < 1:
        print("Before grade 1")
        return
    elif idx > 16:
        print("Grade 16+")
        return

    print("Grade:", idx)


def get_string(value):
    while True:
        try:
            return input(value)
        except:
            pass


main()
