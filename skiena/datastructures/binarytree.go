package datastructures

//BTNode is a binary tree node
type BTNode struct {
	Key   int
	Left  *BTNode
	Right *BTNode
}

//BTNodeDFSApply applies a function to all elements of the tree in DFS order
func BTNodeDFSApply(node *BTNode, f func(*BTNode)) {
	stack := NewBTNodeStack()

	n := node

	for n != nil {
		if n.Right != nil {
			stack.Push(n.Right)
		}
		if n.Left != nil {
			stack.Push(n.Left)
		}
		f(n)
		n = stack.Pop()
	}

}

//BTNodeBFSApply applies a function to all elements of the tree in BFS order
func BTNodeBFSApply(node *BTNode, f func(*BTNode)) {
	queue := NewBTNodeQueue()

	n := node

	for n != nil {
		if n.Left != nil {
			queue.Enqueue(n.Left)
		}
		if n.Right != nil {
			queue.Enqueue(n.Right)
		}
		f(n)
		n = queue.Dequeue()
	}

}

//Find a key
func Find(t *BTNode, key int) *BTNode {
	if t == nil {
		return nil
	}

	if t.Key == key {
		return t
	}

	left := Find(t.Left, key)
	if left != nil {
		return left
	}
	right := Find(t.Right, key)
	if right != nil {
		return right
	}

	return nil
}
