package binarysearchtree

import "fmt"

type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
	Key    int
}

type Tree struct {
	Root *Node
}

func InOrderTreeWalk(n *Node) {
	if n != nil {
		InOrderTreeWalk(n.Left)
		fmt.Println(n.Key)
		InOrderTreeWalk(n.Right)
	}
}

func TreeSearch(node *Node, key int) *Node {
	if node == nil {
		return nil
	}
	if node.Key == key {
		return node
	}
	if node.Key > key {
		return TreeSearch(node.Left, key)
	} else {
		return TreeSearch(node.Right, key)
	}
}

func Minimum(node *Node) *Node {
	if node == nil {
		return nil
	}
	if node.Left == nil {
		return node
	}
	return Minimum(node.Left)
}

func Maximum(node *Node) *Node {
	if node == nil {
		return nil
	}
	if node.Right == nil {
		return node
	}
	return Maximum(node.Right)
}

func TreeSuccessor(node *Node) *Node {
	if node == nil {
		return nil
	}
	if node.Right != nil {
		return Minimum(node.Right)
	}
	parent := node.Parent
	x := node
	for parent != nil && parent.Right == x {
		x = parent
		parent = parent.Parent
	}

	return parent
}

func TreeInsert(t *Tree, node *Node) {
	var y *Node
	x := t.Root
	for x != nil {
		y = x
		if x.Key > node.Key {
			x = x.Left
		} else {
			x = x.Right
		}
	}
	if y == nil {
		t.Root = node
	} else {
		node.Parent = y
		if y.Key > node.Key {
			y.Left = node
		} else {
			y.Right = node
		}
	}
}
