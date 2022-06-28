package main

import "fmt"

func main() {
	var n int
	var k int
	fmt.Scan(&n, &k)
	scoresEarned := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&scoresEarned[i])
	}

	score := scoresEarned[k-1]
	count := 0

	for i := 0; i < n; i++ {
		if scoresEarned[i] >= score && scoresEarned[i] != 0 {
			count++
		}
	}
	fmt.Println(count)
}
