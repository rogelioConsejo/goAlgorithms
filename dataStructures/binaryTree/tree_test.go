package binaryTree

import (
	"fmt"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	tree := BinaryTree()
	root := tree.Root()
	root.SetValue(3)
	root.SetLeft(Node(6, Node(9, nil, Node(5, Node(8, Node(7,nil,nil), Node(1,nil,nil)), nil)), Node(2, nil, nil)))
	root.SetRight(Node(1, nil, Node(4,nil,nil)))
	fmt.Printf("%+v\n", tree.LevelOrder())
	fmt.Printf("%+v\n", tree.MaxDepth())
}
