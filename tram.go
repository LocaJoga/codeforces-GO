package main

import (
	"fmt"
)

func main() {
	var exit int
	var enter int
	var stops int
	passenger := 0
	tramCapacity := 0
	fmt.Scan(&stops)
	for i := 0; i < stops; i++ {
		fmt.Scan(&exit, &enter)
		passenger -= exit
		passenger += enter
		tramCapacity = Max(tramCapacity, passenger)
	}
	fmt.Println(tramCapacity)
}

func Max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}
