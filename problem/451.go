package problem

import (
	"container/heap"
)

type IntFrequency struct {
	Value, Frequency int32
}

type PriorityQueue []*IntFrequency

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Frequency > pq[j].Frequency
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*IntFrequency)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n - 1]
	*pq = old[0: n - 1]
	return item
}



func frequencySort(s string) string {
	a := map[int32]int32{}
	for _, v := range s {
		a[v]++
	}
	b := ""
	pq := make(PriorityQueue, len(a))
	index := 0
	for k , v := range a {
		pq[index] = &IntFrequency{k, v}
		index ++
	}
	heap.Init(&pq)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*IntFrequency)
		frequency := item.Frequency
		for frequency > 0 {
			b += string(item.Value)
			frequency--
		}
	}
	return b
	//box := make([][]int, len(s) + 1)
	//for i, k := range a {
	//	if k > 0 {
	//		if box[k] == nil {
	//			box[k] = make([]int, len(s) + 1)
	//		}
	//		box[k] = append(box[k], i)
	//	}
	//}
	//for i := len(s); i >= 0; i-- {
	//	if box[i] != nil {
	//		for _, k := range box[i] {
	//			if k != 0 {
	//				p := i
	//				for p > 0 {
	//					b += string(k)
	//					p --
	//				}
	//			}
	//		}
	//	}
	//}
	//return b
}