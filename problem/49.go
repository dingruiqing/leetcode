package problem

func groupAnagrams(strs []string) [][]string {
	mapList := make(map[[26]int]int, 0)
	res := make([][]string, 0)
	i := 0
	for _, v := range strs {
		m := [26]int{}
		for _, c := range v {
			m[c-97]++
		}
		mp, ok := mapList[m]
		if ok {
			res[mp] = append(res[mp], v)
		} else {
			res = append(res, []string{v})
			mapList[m] = i
			i++
		}
	}
	return res
}
