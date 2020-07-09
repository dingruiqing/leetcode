package main

import (
	"leetcode/problem"
)

func main() {
	h := problem.BuildListNode([]int{4, 2, 3, 1})

	problem.PrintListNode(problem.SortList(h))
}
