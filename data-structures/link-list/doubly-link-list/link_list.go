package doublylinklist

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

type LinkList struct {
	L    int
	Head *Node
}

func (l *LinkList) Search(k int) *Node {
	for x := l.Head; x != nil; x = x.Next {
		if x.Value == k {
			return x
		}
	}
	return nil
}

func (l *LinkList) Insert(node *Node) {
	node.Prev = nil
	if l.Head == nil {
		l.Head = node
	} else {
		l.Head.Prev = node
		node.Next = l.Head
		l.Head = node
	}
	l.L += 1
}

func (l *LinkList) Delete(node *Node) {
	if node.Prev == nil {
		l.Head = node.Next
	} else {
		node.Prev.Next = node.Next
		if node.Next != nil {
			node.Next.Prev = node.Prev
		}
	}
	l.L--
}
