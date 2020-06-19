package problem

func isIsomorphic(s string, t string) bool {
	m := make(map[uint8]uint8)
	p := make(map[uint8]uint8)
	j := 0
	for i := 0; i < len(s); i++ {
		if j >= len(t) {
			return false
		}
		//if s[i] == t[j] {
		//	j++
		//	continue
		//}
		k, ok := m[s[i]]
		k2, ok2 := p[t[j]]
		if !ok && !ok2 {
			m[s[i]] = t[j]
			p[t[j]] = s[i]
		} else if ok && ok2{
			if t[j] != k || s[i] != k2{
				return false
			}
		} else {
			return false
		}
		j++
	}
	if j != len(t) {
		return false
	}
	return true
}