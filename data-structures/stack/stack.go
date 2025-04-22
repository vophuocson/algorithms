package stack

type Node struct {
	Value int
	Next  *Node
}

type Stack struct {
	Top *Node
	L   int
}

func (s *Stack) StackEmpty() bool {
	return s.Top == nil
}

func (s *Stack) Push(node *Node) {
	node.Next = s.Top
	s.Top = node
}

func (s *Stack) Pop() *Node {
	if s.StackEmpty() {
		return nil
	}
	node := s.Top
	s.Top = s.Top.Next
	return node
}
