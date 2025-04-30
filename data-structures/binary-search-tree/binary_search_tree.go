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
