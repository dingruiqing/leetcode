package main

import (
	"fmt"
	"leetcode/problem"
)

func main() {
	pointMap := make(map[int]map[int]bool)
	for i := 0; i <= 3; i++ {
		pointMap[i] = make(map[int]bool)
	}
	pointMap[1][2] = true
	_, ok := pointMap[1][3]
	fmt.Println(ok)
	fmt.Println(problem.UniquePaths(3, 2))
}
