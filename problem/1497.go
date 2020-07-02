package problem

func CanArrange(arr []int, k int) bool {
	n := len(arr)
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		arr[i] %= k
		_, ok := m[arr[i]]
		if ok {
			m[arr[i]]++
		} else {
			m[arr[i]] = 1
		}
	}
	if m[0]%2 != 0 {
		return false
	}
	for i := 1; i <= k-1; i++ {
		if m[i]+m[i-k] != m[k-i]+m[-i] {
			return false
		}
	}
	return true
}
