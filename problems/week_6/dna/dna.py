import csv
import sys


def main():
    if len(sys.argv) < 3:
        print("Usage: python dna.py file.csv file.txt")
        exit(1)

    rows = []
    with open(sys.argv[1],'r') as file:
        reader = csv.DictReader(file)
        for row in reader:
            rows.append(row)

    with open(sys.argv[2],'r') as file:
        text = file.read()

    subsequence = list(rows[0].keys())[1:]
    results = {}

    for i in subsequence:
        results[i] = longest_match(text, i)


    for name in rows:
        counter = 0
        for i in subsequence:
            if int(name[i]) == results[i]:
                counter += 1

        if counter == len(subsequence):
            print(name["name"])
            return

    print("No match")
    return

def longest_match(sequence, subsequence):
    """Returns length of longest run of subsequence in sequence."""

    # Initialize variables
    longest_run = 0
    subsequence_length = len(subsequence)
    sequence_length = len(sequence)


    for i in range(sequence_length):
        count = 0
        while True:
            start = i + count * subsequence_length
            end = start + subsequence_length

            if sequence[start:end] == subsequence:
                count += 1
            else:
                break
        longest_run = max(longest_run, count)

    return longest_run


main()
