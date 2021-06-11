package kth_largest_element_in_a_stream

import (
	"leetcode/note/heap"
)

type KthLargest struct {
	h *heap.MinHeap
}

func Constructor(k int, nums []int) KthLargest {
	h := heap.CreateMinHeap(k, nums)
	return KthLargest{h}
}

// 使用最小堆 最小堆的大小为k 堆顶为第k大的值
func (this *KthLargest) Add(val int) int {
	this.h.Add(val)
	return this.h.Index(0)
}


/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

// 题目意思
// 2 3 4 5 8
// 2 3 4 5 5 8
// 2 3 4 5 5 8 10
// 2 3 4 5 5 8 9 10
// 2 3 4 4 5 5 8 9 10
