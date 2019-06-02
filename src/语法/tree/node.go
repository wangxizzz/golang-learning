package tree

import "fmt"

type Node struct {
	Value int
	// Left、Right为指针类型，指向这个节点(是这样的，如果是值传递，根本无法改变Left与Right的引用)
	Left, Right *Node
}

// (node Node)函数的接收者(值传递)
func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

// （引用传递）
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting Value to nil " +
			"node. Ignored.")
		return
	}
	node.Value = value
}

// go中没有构造函数，因此可以通过factory控制构造
func CreateNode(value int) *Node {
	return &Node{Value: value}
}
