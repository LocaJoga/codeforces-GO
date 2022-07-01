package main

import "fmt"

var stepLength int

func elephant(stepLength int) int {
	minStep := stepLength / 5
	if stepLength%5 > 0 {
		minStep++
	}
	return minStep
}

func main() {
	fmt.Scan(&stepLength)
	minimumSteps := elephant(stepLength)
	fmt.Println(minimumSteps)
}
