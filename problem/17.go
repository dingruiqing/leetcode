package problem

func letterCombinations(digits string) []string {
	var a []string
	for _, value := range digits{
		v := value - 50
		var t []string
		if v >= 0 && v < 5{
			t = []string{string(97 + 3 * v), string(98 + 3 * v), string(99 + 3 * v)}
		} else if v == 5 {
			t = []string{"p", "q", "r", "s"}
		} else if v == 6 {
			t = []string{"t", "u", "v"}
		} else if v == 7 {
			t = []string{"w", "x", "y", "z"}
		}
		if len(a) == 0{
			a = t
		} else {
			var c []string
			for _, k := range a {
				for _, v := range t {
					c = append(c, k + v)
				}
			}
			a = c
		}
	}
	return a
}
