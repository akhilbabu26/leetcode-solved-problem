package main

import "fmt"

type KthLargest struct {
	k    int
	data []int
}

func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{k: k}

	for _, v := range nums {
		kl.Add(v)
	}

	return kl
}

func (h *KthLargest) Add(val int) int {

	h.data = append(h.data, val)
	i := len(h.data) - 1

	for i > 0 {
		parent := (i - 1) / 2

		if h.data[parent] <= h.data[i] {
			break
		}

		h.data[parent], h.data[i] = h.data[i], h.data[parent]

		i = parent
	}

	if len(h.data) > h.k {
		h.removeMin()
	}

	return h.data[0]
}

func (h *KthLargest) removeMin() {
	n := len(h.data)

	h.data[0] = h.data[n-1]
	h.data = h.data[:n-1]

	i := 0

	for {
		smallest := i

		left := 2*i + 1
		right := 2*i + 2

		if left < len(h.data) && h.data[left] < h.data[smallest] {
			smallest = left
		}

		if right < len(h.data) && h.data[right] < h.data[smallest] {
			smallest = right
		}

		if smallest == i {
			break
		}

		h.data[i], h.data[smallest] = h.data[smallest], h.data[i]

		i = smallest
	}
}


/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */