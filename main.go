package main

import (
	"fmt"
)

func main() {
	m1 := [1]int{}
	m1[0] = 1
	m2 := [1]int{}
	m2[0] = 1
	fmt.Println(m1 == m2)
}
