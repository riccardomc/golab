package braces

import "fmt"

type byteStack struct {
	stack []byte
}

func newByteStack() *byteStack {
	return &byteStack{make([]byte, 0)}
}

func (s *byteStack) push(b byte) {
	s.stack = append(s.stack, b)
}

func (s *byteStack) pop() byte {
	b := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return b
}

func (s *byteStack) empty() bool {
	return len(s.stack) == 0
}

func validBraces(input string) bool {

	closer := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}

	s := newByteStack()
	bytes := []byte(input)
	for _, b := range bytes {
		switch b {
		case '{':
			fallthrough
		case '[':
			fallthrough
		case '(':
			s.push(b)
		case '}':
			fallthrough
		case ']':
			fallthrough
		case ')':
			if s.empty() {
				fmt.Println("nowoo")
				return false
			}
			c := s.pop()
			if c != closer[b] {
				return false
			}
		}
	}

	return s.empty()
}
