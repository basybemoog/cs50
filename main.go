package main

import (
	"CS50/problems/week_2/Substitution"
	"fmt"
)

func main() {
	chippedText := Substitution.Substitution("YTNSHKVEFXRBAUQZCLWDMIPGJO", "Hello!")
	fmt.Printf(chippedText)
}
