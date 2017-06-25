package priorityqueue

//PriorityQueue is a PriorityQueue
type PriorityQueue struct {
	data     []int
	size     int
	capacity int
}

//NewPriorityQueue gives you a new PriorityQueue of capacity capacity
func NewPriorityQueue(capacity int) *PriorityQueue {
	return &PriorityQueue{make([]int, capacity), 0, capacity}
}

//Insert an element in the PriorityQueue return true on success, false
//otherwise
func (p *PriorityQueue) Insert(k int) bool {

	if p.size == p.capacity {
		return false
	}

	p.data[p.size] = k
	p.size++
	p.BubbleUp(p.size - 1)

	return true
}

//BubbleUp the element at i until priority order is restored
func (p *PriorityQueue) BubbleUp(i int) {
	if parent(i) == -1 {
		return
	}

	if p.data[parent(i)] > p.data[i] {
		p.Swap(parent(i), i)
		p.BubbleUp(parent(i))
	}
}

//Swap the elements i and j
func (p *PriorityQueue) Swap(i, j int) {
	t := p.data[i]
	p.data[i] = p.data[j]
	p.data[j] = t
}

//ExtractMin returns the minimum and removes it from the queue
func (p *PriorityQueue) ExtractMin() int {
	min := -1

	if p.size != 0 {
		min = p.data[0]
		p.data[0] = p.data[p.size-1]
		p.size--
		p.BubbleDown(0)
	}

	return min
}

//BubbleDown the element at i until priority order is restored
func (p *PriorityQueue) BubbleDown(i int) {

	min := p.data[i]
	minIndex := i

	l := left(i)
	if l < p.size {
		if min > p.data[l] {
			min = p.data[l]
			minIndex = l
		}
	}

	r := right(i)
	if r < p.size {
		if min > p.data[r] {
			min = p.data[r]
			minIndex = r
		}
	}

	if minIndex != i {
		p.Swap(i, minIndex)
		p.BubbleDown(minIndex)
	}

}

//Size returns the number of elements in the PriorityQueue
func (p *PriorityQueue) Size() int {
	return p.size
}

func parent(i int) int {
	if i == 0 {
		return -1
	}

	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}
