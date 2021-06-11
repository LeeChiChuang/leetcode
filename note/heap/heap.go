package heap

type MinHeap struct {
	k int
	heap []int
}

func CreateMinHeap(k int, nums []int) *MinHeap {
	h := &MinHeap{k, []int{} }
	for i, v := range nums {
		h.heap = append(h.heap, v)
		h.up(i)
	}

	return h
}

func (h *MinHeap)Add(num int)  {
	if len(h.heap) < h.k {
		h.heap = append(h.heap, num)
		h.up(len(h.heap)-1)
	} else if num > h.heap[0] {
		h.heap[0] = num
		h.down(0)
	}
}

func (h *MinHeap) Index(i int) int {
	if len(h.heap) <= i {
		return -1
	}
	return h.heap[i]
}

func (h *MinHeap) up(i int)  {
	for i > 0 {
		parent := (i-1) / 2
		if h.heap[parent] > h.heap[i] {
			h.heap[parent], h.heap[i] = h.heap[i], h.heap[parent]
			i = parent
		} else {
			break
		}
	}
}

func (h *MinHeap) down(i int) {
	for 2*i + 1 < len(h.heap) {
		child := 2*i + 1
		if child+1 < len(h.heap) && h.heap[child] > h.heap[child+1] {
			child++
		}
		if h.heap[i] > h.heap[child] {
			h.heap[i], h.heap[child] = h.heap[child], h.heap[i]
			i = child
		} else {
			break
		}
	}
}
