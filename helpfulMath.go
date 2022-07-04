package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)
	s := strings.Split(str, "")
	nums := make([]string, 0)
	for i := 0; i < len(s); i++ {
		if s[i] != "+" {
			nums = append(nums, s[i])
		}
	}
	sort.Strings(nums)
	for i := 0; i < len(nums); i++ {
		fmt.Printf("%s", nums[i])
		if i < len(nums)-1 {
			fmt.Printf("+")
		}
	}

}
