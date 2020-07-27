package problem

func IsBipartite(graph [][]int) bool {
	n := len(graph)
	if n == 0 {
		return true
	}
	queue := make([]int, 0)
	visit := make([]int, n)
	index := 0
	for index < n {
		visit[index] = 1
		queue = append(queue, index)
		for len(queue) != 0 {
			tmp := make([]int, 0)
			for _, v := range queue {
				for _, k := range graph[v] {
					if visit[k] == visit[v] {
						return false
					} else {
						if visit[k] == 0 {
							visit[k] = -visit[v]
							tmp = append(tmp, k)
						}
					}
				}
			}
			queue = tmp
		}
		for index < n && visit[index] != 0 {
			index++
		}
	}
	return true
}
