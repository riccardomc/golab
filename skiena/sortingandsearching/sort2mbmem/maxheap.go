package sort2mbmem

//MaxHeap is a MaxHeap
type MaxHeap struct {
	data     []int32
	size     int
	capacity int
}

//Insert k in the Heap
func (h *MaxHeap) Insert(k int32) {
	if h.size == h.capacity {
		return
	}

	h.data[h.size] = k
	h.BubbleUp(h.size)
	h.size++
}

//BubbleUp the element at i
func (h *MaxHeap) BubbleUp(i int) {
	parent := parent(i)

	if parent < 0 {
		return
	}

	if h.data[i] > h.data[parent] {
		h.Swap(i, parent)
		h.BubbleUp(parent)
	}
}

//Max extracts the max from the heap
func (h *MaxHeap) Max() int32 {
	if h.size <= 0 {
		return -1
	}

	max := h.data[0]
	h.size--
	h.data[0] = h.data[h.size]
	h.BubbleDown(0)

	return max
}

//NewMaxHeap gets you a new MaxHeap
func NewMaxHeap(capacity int) *MaxHeap {
	return &MaxHeap{make([]int32, capacity), 0, capacity}
}

//BubbleDown the element at i
func (h *MaxHeap) BubbleDown(i int) {

	max := h.data[i]
	maxIndex := i

	l := left(i)
	if l < h.size {
		if max < h.data[l] {
			max = h.data[l]
			maxIndex = l
		}
	}

	r := right(i)
	if r < h.size {
		if max < h.data[r] {
			max = h.data[r]
			maxIndex = r
		}
	}

	if maxIndex != i {
		h.Swap(i, maxIndex)
		h.BubbleDown(maxIndex)
	}

}

//Swap the elements at i and j
func (h *MaxHeap) Swap(i, j int) {
	t := h.data[i]
	h.data[i] = h.data[j]
	h.data[j] = t
}

func left(i int) int {
	return i*2 + 1
}

func right(i int) int {
	return i*2 + 2
}

func parent(i int) int {

	if i <= 0 {
		return -1
	}

	return (i - 1) / 2
}
