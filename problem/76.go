package problem

//func minWindow(s string, t string) string {
//	start := 0
//	end := 0
//	tList := [128]int{}
//	for _, v := range t {
//		tList[v]++
//	}
//	sl := len(s)
//	l := sl + 1
//	rStart := 0
//	rEnd := l
//	sList := [128]int{}
//	for end < sl {
//		if tList[s[end]] > 0{
//			sList[s[end]] ++
//			flag := true
//			for flag {
//				for i, v := range sList {
//					if v < tList[i] {
//						flag = false
//					}
//				}
//				if flag {
//					for start < sl && tList[s[start]] <= 0 {
//						start ++
//					}
//					ls := end - start
//					if ls >= 0 && ls < l {
//						rStart = start
//						rEnd = end + 1
//						l = ls
//					}
//					sList[s[start]]--
//					start ++
//					for start < sl && tList[s[start]] <= 0 {
//						start ++
//					}
//				}
//			}
//		}
//		end ++
//	}
//	if l == sl + 1{
//		return ""
//	}
//	return s[rStart: rEnd]
//}

func MinWindow(s string, t string) string {
	n := len(s)
	m := len(t)
	target := make(map[uint8]int)
	indexList := make([]int, 0)
	for i := 0; i < m; i++ {
		target[t[i]]++
	}
	indexMap := make(map[uint8]int)
	start, end := 0, 0
	min := 0x7fffffff
	for i := 0; i < n; i++ {
		if _, ok := target[s[i]]; ok {
			indexList = append(indexList, i)
			indexMap[s[i]]++
			isSame := true
			for isSame {
				for k, v := range target {
					if indexMap[k] < v {
						isSame = false
						break
					}
				}
				if isSame {
					j := indexList[0]
					if i-indexList[0] < min {
						min = i - j
						start = j
						end = i
					}
					indexMap[s[j]]--
					indexList = indexList[1:]
				}
			}
		}
	}
	if min != 0x7fffffff {
		return s[start : end+1]
	} else {
		return ""
	}
}
