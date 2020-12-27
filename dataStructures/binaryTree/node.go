package binaryTree

type node struct {
	value interface{}
	left  *node
	right *node
}

//Constructor, Getters & Setters
func Node(value interface{}, left *node, right *node) *node {
	return &node{value: value, left: left, right: right}
}

func (n *node) Value() interface{} {
	return n.value
}

func (n *node) SetValue(value interface{}) {
	n.value = value
}

func (n *node) Right() *node {
	return n.right
}

func (n *node) SetRight(right *node) {
	n.right = right
}

func (n *node) Left() *node {
	return n.left
}

func (n *node) SetLeft(left *node) {
	n.left = left
}

//Other Methods
func (n *node) LevelOrder() (levelOrder [][]interface{}) {
	if n != nil {
		levelOrder = [][]interface{}{{n.value}}
		if n.left != nil || n.right != nil {
			leftBranch := n.left.LevelOrder()
			rightBranch := n.right.LevelOrder()
			levelOrder = append(levelOrder, mergeArraysOfArrays(leftBranch, rightBranch)...)
		}
	}

	return levelOrder
}

func mergeArraysOfArrays(firstArrayOfArrays [][]interface{}, secondArrayOfArrays [][]interface{}) (merged [][]interface{}) {
	mergeAndAppend := func(firstArrayOfArrays [][]interface{}, secondArrayOfArrays [][]interface{}) [][]interface{} {
		mergeExisting := func(firstArrayOfArrays [][]interface{}, secondArrayOfArrays [][]interface{}) (lastIndex int, merged [][]interface{}) {
			for lastIndex = 0; lastIndex < len(firstArrayOfArrays) && lastIndex < len(secondArrayOfArrays); lastIndex++ {
				merge := append(firstArrayOfArrays[lastIndex], secondArrayOfArrays[lastIndex]...)
				merged = append(merged, merge)
			}
			return lastIndex, merged
		}
		appendRest := func(firstArrayOfArrays [][]interface{}, secondArrayOfArrays [][]interface{}, index int, merged [][]interface{}) [][]interface{} {
			for ; index < len(firstArrayOfArrays); index++ {
				merged = append(merged, firstArrayOfArrays[index])
			}
			for ; index < len(secondArrayOfArrays); index++ {
				merged = append(merged, secondArrayOfArrays[index])
			}
			return merged
		}

		var merged [][]interface{}
		var index int
		var cache [][]interface{}
		index, cache = mergeExisting(firstArrayOfArrays, secondArrayOfArrays)
		merged = append(merged, cache...)
		merged = appendRest(firstArrayOfArrays, secondArrayOfArrays, index, merged)
		return merged
	}

	isFirstNil := firstArrayOfArrays != nil
	isSecondNil := secondArrayOfArrays != nil

	if isFirstNil && isSecondNil {
		merged = mergeAndAppend(firstArrayOfArrays, secondArrayOfArrays)
	} else if isFirstNil {
		merged = append(merged, firstArrayOfArrays...)
	} else {
		merged = append(merged, secondArrayOfArrays...)
	}

	return merged
}

func (n node) MaxDepth() (maxDepth int) {
	bothNil := n.left == nil && n.right == nil
	bothNonNil := n.left != nil && n.right != nil
	if bothNil {
		maxDepth = 0
	} else if bothNonNil {
		if n.left.MaxDepth() > n.right.MaxDepth() {
			maxDepth = n.left.MaxDepth()
		} else {
			maxDepth = n.right.MaxDepth()
		}
	} else if n.left != nil {
		maxDepth = n.left.MaxDepth()
	} else {
		maxDepth = n.right.MaxDepth()
	}
	return maxDepth + 1
}