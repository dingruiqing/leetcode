package problem

import (
	"strconv"
)

func RestoreIpAddresses(s string) []string {
	a := []string{}
	ipAddresses(s, "", 4, &a)
	return a
}

func ipAddresses(s string, ip string, n int, a *[]string) {
	if n == 0 {
		if len(s) == 0 {
			*a = append(*a, ip)
		} else {
			return
		}
	}
	n--
	for i := 1; i <= len(s); i++ {
		if i > 3 {
			break
		}
		tmp := s[:i]
		if len(tmp) != 1 && tmp[0] == 48 {
			continue
		}
		num, _ := strconv.Atoi(tmp)
		if num > 255 {
			break
		}
		if n == 3 {
			ipAddresses(s[i:], strconv.Itoa(num), n, a)
		} else {
			ipAddresses(s[i:], ip+"."+strconv.Itoa(num), n, a)

		}
	}
}
