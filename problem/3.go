package problem

func lengthOfLongestSubstring(s string) int {
	tmpMap := make(map[string]int)
	i := -1
	max := 0
	for index, value := range s {
		str := string(value)
		v, ok := tmpMap[str]
		if ok && v >= i {
			i = v
		}
		tmpMap[str] = index
		length := index - i
		if length > max{
			max = length
		}
	}
	return max
}
