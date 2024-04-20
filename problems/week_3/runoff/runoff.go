package runoff

import (
	"fmt"
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

func Runoff(candidatesCount int, voterCounter int, argv []string) string {

	if candidatesCount < 2 {
		return "Usage: runoff [candidate ...]"
	}
	candidateCounter = candidatesCount

	if candidateCounter > maxCandidates {
		return "Maximum number of candidates is 9"
	}
	for i := 0; i < candidateCounter; i++ {
		candidates[i].name = argv[i]
		candidates[i].votes = 0
		candidates[i].eliminated = false
	}
	if voterCounter > maxVoters {
		return "Maximum number of voters is 100"
	}
	for i := 0; i < voterCounter; i++ {

		for j := 0; j < candidateCounter; j++ {
			var name string
			fmt.Printf("Rank %d: ", j+1)
			fmt.Scan(&name)

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
					return candidates[i].name
				}
			}
			break
		}
		eliminate(min)

		for i := 0; i < candidateCounter; i++ {
			candidates[i].votes = 0
		}
	}
	return ""
}

func vote(voter, rank int, name string) bool {

	for i := 0; i < candidateCounter; i++ {
		if candidates[i].name == name {
			prefences[voter][rank] = i

			return true
		}
	}

	return false
}
func tabulate() {
	for i := 0; i < voterCounter; i++ {
		for j := 0; j < candidateCounter; j++ {
			if candidates[prefences[j][i]].eliminated == false {
				candidates[prefences[j][i]].votes++
				break
			}
		}
	}
	return
}

func printWinner() bool {
	majority := voterCounter/2 + 1
	for i := 0; i < voterCounter; i++ {
		if candidates[i].votes >= majority && candidates[i].eliminated == false {
			fmt.Printf("%s", candidates[i].name)
			return true
		}
	}
	return false
}
func findMin() int {
	minVote := 0
	for i := 0; i < candidateCounter; i++ {
		if candidates[i].votes < minVote && candidates[i].eliminated == false {
			minVote = candidates[i].votes
		}

	}
	return minVote
}
func isTie(min int) bool {
	for i := 0; i < candidateCounter; i++ {
		if candidates[i].votes == min && candidates[i].eliminated == false {
			return true
		}
	}
	return false
}
func eliminate(min int) {
	for i := 0; i < candidateCounter; i++ {
		if candidates[i].votes == min && candidates[i].eliminated == false {
			candidates[i].eliminated = true
			break
		}
	}
	return
}
