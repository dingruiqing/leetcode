package problem

func findAnagrams(s string, p string) []int {
	n := len(p) - 1
	pList := [26]uint{}
	for _, v := range p {
		pList[v - 97]++
	}
	var result []int
	sList := [26]uint{}
	for i, v := range s {
		sList[v - 97]++
		if i >= n {
			if pList[v - 97] > 0 && pList[s[i - n] - 97] > 0 {
				if pList == sList {
					result = append(result, i - n)
				}
			}
			sList[s[i - n] - 97]--
		}
	}
	return result
}

//func findAnagrams(s string, p string) []int {
//	pHash := [26]int{}
//
//	for i := 0; i < len(p); i++ {
//		pHash[p[i] - 'a']++
//	}
//
//	result := make([]int, 0)
//	tempHash := [26]int{}
//
//	for i := 0; i < len(s); i++ {
//		tempHash[s[i] - 'a']++
//
//		if i + 1 >= len(p) {
//			if pHash[s[i] - 'a'] > 0 && pHash[s[i - len(p) + 1] - 'a'] > 0 {
//				if tempHash == pHash {
//					result = append(result, i - len(p) + 1)
//				}
//			}
//
//			tempHash[s[i - len(p) + 1] - 'a']--
//		}
//	}
//
//	return result
//}
