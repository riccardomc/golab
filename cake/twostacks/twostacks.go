package twostacks

//IntStack is a Stack for integers
type IntStack struct {
	data []int
}

func newIntStack() *IntStack {
	return &IntStack{make([]int, 0)}
}

func (s *IntStack) push(d int) {
	s.data = append(s.data, d)
}

func (s *IntStack) empty() bool {
	if len(s.data) == 0 {
		return true
	}
	return false
}

func (s *IntStack) pop() int {
	d := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return d
}

//IntQueue is a queue that uses two stacks
type IntQueue struct {
	stack1 *IntStack
	stack2 *IntStack
}

//NewIntQueue gives you a new queue
func NewIntQueue() *IntQueue {
	return &IntQueue{
		newIntStack(),
		newIntStack(),
	}
}

//EnqueueNaive classic
func (q *IntQueue) EnqueueNaive(d int) {
	q.stack1.push(d)
}

//DequeueNaive classic
func (q *IntQueue) DequeueNaive() int {
	for !q.stack1.empty() {
		q.stack2.push(q.stack1.pop())
	}
	d := q.stack2.pop()
	for !q.stack2.empty() {
		q.stack1.push(q.stack2.pop())
	}
	return d
}

//Enqueue using two stacks
func (q *IntQueue) Enqueue(d int) {
	q.stack1.push(d)
}

//Dequeue using two stacks
func (q *IntQueue) Dequeue() int {
	if !q.stack2.empty() {
		return q.stack2.pop()
	}

	for !q.stack1.empty() {
		q.stack2.push(q.stack1.pop())
	}

	return q.stack2.pop()
}
