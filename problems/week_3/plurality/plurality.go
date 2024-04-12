package plurality

import (
	"fmt"
	"strings"
)

const max = 9

type candidate struct {
	name  string
	votes int
}

var candidates [max]candidate
var candidateCount int

func Plurality(args int, argv []string) string {

	if args < 2 {
		return "Usage: plurality [candidate ...]"
	}
	candidateCount = args - 1
	if candidateCount > max {
		return "Maximum number of candidates is 9"
	}

	for i := 0; i < candidateCount; i++ {
		candidates[i].name = argv[i+1]
		candidates[i].votes = 0
	}

	for i := 0; i < args; i++ {
		var name string
		fmt.Printf("Vote: ")
		fmt.Scan(&name)
		if !vote(name) {
			return "Invalid vote."
		}
	}
	return printWinner()
}

func vote(name string) bool {
	for i := 0; i < len(candidates); i++ {
		if strings.Compare(candidates[i].name, name) == 0 {
			candidates[i].votes += 1
			return true
		}
	}
	return false

}
func printWinner() string {

	changedVar := 0

	for i := 0; i < max-1; i++ {
		for j := 0; j < max-1; j++ {
			if candidates[j].votes > candidates[j+1].votes {
				changedVar = candidates[j+1].votes
				candidates[j+1].votes = candidates[j].votes
				candidates[j].votes = changedVar
			}
		}
	}
	var balls [9]int
	j := 0
	for i := len(balls) - 1; i > 0; i-- {
		balls[j] = candidates[i].votes
		j++
	}
	var winners string
	for i := 1; i < len(balls); i++ {
		if balls[i-1] > balls[i] && balls[i] > 0 {
			return candidates[i-1].name
		} else if candidates[i-1].votes == candidates[i].votes || candidates[i-1].votes == candidates[i-2].votes {
			winners += candidates[i-1].name + " "
		}
	}
	return winners
}
