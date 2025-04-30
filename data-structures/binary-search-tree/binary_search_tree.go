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

func Transplant(t *Tree, u, v *Node) {
	if u.Parent == nil {
		t.Root = v
	} else if u == u.Parent.Left {
		u.Parent.Left = v
	} else {
		u.Parent.Right = v
	}
	if v != nil {
		v.Parent = u.Parent
	}
}

func DeleteNode(t *Tree, z *Node) {
	if z.Left == nil {
		Transplant(t, z, z.Right)
	} else if z.Right == nil {
		Transplant(t, z, z.Left)
	} else {
		successor := Minimum(z.Right)
		if successor.Parent != z {
			Transplant(t, successor, successor.Right)
			successor.Right = z.Right
			if successor.Right != nil {
				successor.Right.Parent = successor
			}
		}
		Transplant(t, z, successor)
		successor.Left = z.Left
		if successor.Left != nil {
			successor.Left.Parent = successor
		}
	}
}
