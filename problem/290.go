package problem

import (
	"strings"
)

func wordPattern(pattern string, str string) bool {
	s := strings.Split(str, " ")
	m := make(map[int32]string)
	p := make(map[string]int32)
	j := 0
	for _, i := range pattern {
		if j > len(s) {
			return false
		}
		k, ok := m[i]
		k2, ok2 := p[s[j]]
		if ok || ok2 {
			if s[j] != k || i != k2{
				return false
			}
		} else {
			m[i] = s[j]
			p[s[j]] = i
		}
		j++
	}
	if j != len(s) {
		return false
	}
	return true
}
