package main

import (
	"fmt"
	"go-Study/Go-project/tree"
)

type myTreenode struct {
	node *tree.Node
}

func (myNode *myTreenode) postOrder()  {
	if myNode == nil  || myNode.node == nil {
		return
	}
	left := myTreenode{myNode.node.Left}
	right := myTreenode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
	myNode.node.Print()

}

func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	root.Traverse()
	fmt.Println()
	myRoot := myTreenode{&root}
	myRoot.postOrder()
}
