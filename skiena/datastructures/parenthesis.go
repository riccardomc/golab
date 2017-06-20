package datastructures

// Stack is a stack, man
type Stack struct {
	data []int
	head int
}

func (s *Stack) push(v int) {
	s.data = append(s.data, v)
	s.head++
}

func (s *Stack) pop() (bool, int) {
	if s.head == -1 {
		return false, -1
	}

	v := s.data[s.head]
	s.head--

	return true, v
}

func (s *Stack) empty() bool {
	return s.head == -1
}

func new() *Stack {
	return &Stack{make([]int, 0), -1}
}

// Parenthesis returns (false, n) if the input parenthesization is unbalanced,
// with n the position of the first offending parenthesis. (true, -1) if the
// input paranthesization is balanced.
func Parenthesis(input string) (bool, int) {
	stack := new()
	for i := 0; i < len(input); i++ {
		if input[i] == '(' {
			stack.push(i)
		} else {
			found, _ := stack.pop()
			if !found {
				return false, i
			}
		}
	}

	found, offending := stack.pop()

	if !found {
		return true, -1
	}

	return false, offending
}
