package maxstack

import "math"

type stack struct {
	stack []int
}

func newStack() *stack {
	return &stack{make([]int, 0)}
}

func (s *stack) push(item int) {
	s.stack = append(s.stack, item)
}

func (s *stack) pop() int {
	item := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return item
}

func (s *stack) empty() bool {
	return len(s.stack) == 0
}

func (s *stack) peek() int {
	if !s.empty() {
		return s.stack[len(s.stack)-1]
	}

	return math.MinInt64
}

//MaxStack is a stack O(1) getMax
type MaxStack struct {
	data *stack
	max  *stack
}

//NewMaxStack gets you a new MaxStack
func NewMaxStack() *MaxStack {
	return &MaxStack{
		newStack(),
		newStack(),
	}
}

//Push an item in the MaxStack
func (s *MaxStack) Push(item int) {
	s.data.push(item)
	if s.max.empty() {
		s.max.push(item)
	} else if s.max.peek() <= item {
		s.max.push(item)
	}
}

//Pop an item from the MaxStack
func (s *MaxStack) Pop() int {
	item := s.data.pop()

	if item == s.max.peek() {
		s.max.pop()
	}

	return item
}

//Empty returns true if the stack is empty, false otherwise
func (s *MaxStack) Empty() bool {
	return s.data.empty()
}

//Max returns the current max in the MaxStack
func (s *MaxStack) Max() int {
	return s.max.peek()
}
