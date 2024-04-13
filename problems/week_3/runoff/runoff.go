package runoff

import (
	"fmt"
	"strings"
)

const maxVoters = 100
const maxCandidates = 9

var prefences [maxVoters][maxCandidates]int

type candidate struct {
	name       string
	votes      int
	eliminated bool
}

var candidates [maxCandidates]candidate

var voterCounter int
var candidateCounter int

func Runoff(argc int, argv []string) string {
	if argc < 2 {
		return "Usage: runoff [candidate ...]"
	}

	candidateCounter = argc - 1
	if candidateCounter > maxCandidates {
		return "Maximum number of candidates is 9"
	}
	for i := 0; i < candidateCounter; i++ {
		candidates[i].name = argv[i+1]
		candidates[i].votes = 0
		candidates[i].eliminated = false
	}
	if voterCounter > maxVoters {
		return "Maximum number of voters is 100"
	}
	for i := 0; i < voterCounter; i++ {
		for j := 0; j < candidateCounter; j++ {
			var name string
			fmt.Printf("Rank: %d", j+1)
			fmt.Scanf("%s", &name)

			if !vote(i, j, name) {
				return "Invalid vote."
			}
		}
		fmt.Printf("\n")
	}
	for true {
		tabulate()

		var won = printWinner()
		if won {
			break
		}
		var min = findMin()
		var tie = isTie(min)

		if tie {
			for i := 0; i < candidateCounter; i++ {
				if !candidates[i].eliminated {
					fmt.Printf("%s\n", candidates[i].name)
				}
			}
			break
		}
		eliminate(min)

		for i := 0; i < candidateCounter; i++ {
			candidates[i].votes = 0
		}
	}
	return "hello"
}

func vote(voter, rank int, name string) bool {

	for i := 0; i < voter; i++ {
		if strings.Compare(candidates[i].name, name) == 0 {
			prefences[voter][rank] = i
			return true
		}
	}
	return false
}
func tabulate() {

	return
}
func printWinner() bool {
	moreOfHalf := voterCounter/2 + 1
	for i := 0; i < voterCounter; i++ {
		if candidates[i].votes > moreOfHalf && !candidates[i].eliminated {
			return true
		}
	}
	return false
}
func findMin() int {
	minVote := 0
	for i := 0; i < candidateCounter; i++ {
		if candidates[i].votes < minVote && !candidates[i].eliminated {
			minVote = candidates[i].votes
		}

	}
	return minVote
}
func isTie(min int) bool {
	for i := 0; i < candidateCounter; i++ {
		if candidates[i].votes == min && candidates[i].eliminated {
			return true
		}
	}
	return false
}
func eliminate(min int) {

	return
}
