package problem

func minWindow(s string, t string) string {
	start := 0
	end := 0
	tList := [128]int{}
	for _, v := range t {
		tList[v]++
	}
	sl := len(s)
	l := sl + 1
	rStart := 0
	rEnd := l
	sList := [128]int{}
	for end < sl {
		if tList[s[end]] > 0{
			sList[s[end]] ++
			flag := true
			for flag {
				for i, v := range sList {
					if v < tList[i] {
						flag = false
					}
				}
				if flag {
					for start < sl && tList[s[start]] <= 0 {
						start ++
					}
					ls := end - start
					if ls >= 0 && ls < l {
						rStart = start
						rEnd = end + 1
						l = ls
					}
					sList[s[start]]--
					start ++
					for start < sl && tList[s[start]] <= 0 {
						start ++
					}
				}
			}
		}
		end ++
	}
	if l == sl + 1{
		return ""
	}
	return s[rStart: rEnd]
}
