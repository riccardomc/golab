package linkedlistktolast

type node struct {
	key  int
	next *node
}

func ktolast(l *node, k int) *node {

	cursor := l.next
	ktolast := l.next

	for i := 0; i < k; i++ {
		if cursor == nil {
			return nil
		}
		cursor = cursor.next

	}

	for cursor != nil {
		ktolast = ktolast.next
		cursor = cursor.next
	}

	return ktolast
}
