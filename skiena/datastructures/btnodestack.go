package datastructures

//BTNodeStack is a stack for BTNode
type BTNodeStack struct {
	data []*BTNode
}

//Push a BTNode on top of the stack
func (s *BTNodeStack) Push(node *BTNode) {
	s.data = append(s.data, node)
}

//Pop a BTNode from the top of the stack
func (s *BTNodeStack) Pop() *BTNode {
	if len(s.data) == 0 {
		return nil
	}

	node := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return node
}

//NewBTNodeStack returns an initialized BTNodeStack
func NewBTNodeStack() *BTNodeStack {
	return &BTNodeStack{make([]*BTNode, 0)}
}
