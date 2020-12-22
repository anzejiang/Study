package main

type treeNode struct {
	value int
	left, right *treeNode
}

func createNode(value int) *treeNode {
	return &treeNode{value: value}

}

func main() {
	var root treeNode
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	//fmt.Println(root.right.left, root.right.right, root.right.value, root.value)
	root.right.left = new(treeNode)
	root.left.right = createNode(2)

}
