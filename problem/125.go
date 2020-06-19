package problem

func isPalindrome(s string) bool {
	start := 0
	end := len(s) - 1
	for start < end {
		st := s[start]
		e := s[end]
		if st < 48 || (st > 57 && st < 65) || (st > 90 && st < 97) || st > 122 {
			start ++
			continue
		}
		if e < 48 || (e > 57 && e < 65) || (e > 90 && e < 97) || e > 122 {
			end --
			continue
		}
		if st == e {
			start ++
			end --
		} else {
			if (st >= 65 && e >= 65) && ((st - e == 32) || (st - e) == 224) {
				start ++
				end --
			}else {
				return false
			}
		}
	}
	return true

}
