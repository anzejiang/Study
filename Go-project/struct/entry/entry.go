package main



func main() {
	var root tree.TreeNode

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)


	root.left.right = createNode(2)
	root.Print()

	root.right.left.SetValue(4)
	root.right.left.print()

	root.traverse()
}
