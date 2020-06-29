package problem

import "container/heap"

type Queue []int

func (q Queue) Less(i, j int) bool {
	return q[i] < q[j]
}

func (q Queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *Queue) Push(x interface{}) {
	*q = append(*q, x.(int))
}

func (q *Queue) Pop() interface{} {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[:n-1]
	return x
}

func (q Queue) Len() int {
	return len(q)
}

func FindKthLargest(nums []int, k int) int {
	q := &Queue{}
	heap.Init(q)
	min := nums[0]
	for _, v := range nums {
		if q.Len() == k {
			if v > min {
				heap.Push(q, v)
				min = heap.Pop(q).(int)
			}
		} else {
			if v <= min {
				min = v
			}
			heap.Push(q, v)
		}
	}
	return heap.Pop(q).(int)
}
