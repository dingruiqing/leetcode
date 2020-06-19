package problem

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	lenG := len(g)
	indexG := lenG - 1
	lenS := len(s)
	indexS := lenS - 1
	num := 0
	for indexG >= 0 && indexS >= 0{
		if g[indexG] <= s[indexS] {
			indexG --
			indexS --
			num ++
		} else {
			indexG --
		}
	}
	return num
}