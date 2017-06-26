package binarytree

//BTNodeQueue is a queue for BTNode
type BTNodeQueue struct {
	data []*BTNode
}

//Enqueue a node
func (q *BTNodeQueue) Enqueue(node *BTNode) {
	q.data = append(q.data, node)
}

//Dequeue a node
func (q *BTNodeQueue) Dequeue() *BTNode {
	if len(q.data) == 0 {
		return nil
	}

	node := q.data[0]
	q.data = q.data[1:]

	return node
}

func NewBTNodeQueue() *BTNodeQueue {
	return &BTNodeQueue{make([]*BTNode, 0)}
}
