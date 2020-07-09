package problem

func fourSumCount(A []int, B []int, C []int, D []int) int {
	m1 := make(map[int]int)
	n := len(A)
	sum := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			s := A[i] + B[j]
			_, ok := m1[s]
			if ok {
				m1[s]++
			} else {
				m1[s] = 1
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			s := C[i] + D[j]
			v, ok := m1[-s]
			if ok {
				sum += v
			}
		}
	}
	return sum
}
