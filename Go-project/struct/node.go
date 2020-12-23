package tree

import "fmt"

type Node struct {
	Value int
	Left, Right *Node
}

// 值接收者，不会改变结构体中的元数据
func (node Node) Print()  {
}

// 指针接收者，会改变源结构体中的数据
func (node *Node) SetValue(value int)  {
	if node == nil{
		fmt.Println("Setting value to nil " +
			"node. Ignored.")
		return
	}
	node.Value = value
}

func (node *Node) traverse()  {
	if node == nil {
		return
	}
	node.Left.traverse()
	node.Print()
	node.Right.traverse()
}

func CreateNode(value int) *Node {
	return &Node{Value: value}
}


