package problem

func isSubsequence(s string, t string) bool {
	i := 0
	if len(s) == 0{
		return true
	}
	if len(t) == 0{
		return false
	}
	for _, v := range s {
		if i >= len(t) {
			return false
		}
		for v != int32(t[i]) {
			i ++
			if i >= len(t) {
				return false
			}
		}
		i ++
	}
	return  true
}