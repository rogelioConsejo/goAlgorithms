package binaryTree

type binaryTree struct {
	root *node
}

func (t *binaryTree) Root() *node {
	return t.root
}

func (t *binaryTree) SetRoot(root *node) {
	t.root = root
}

func BinaryTree() *binaryTree {
	return &binaryTree{root: new(node)}
}

func (t *binaryTree) LevelOrder() [][]interface{} {
	return t.root.LevelOrder()
}

func (t *binaryTree) MaxDepth() int {
	return t.root.MaxDepth()
}