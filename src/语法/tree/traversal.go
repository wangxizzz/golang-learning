package tree

import "fmt"

func (node *Node) Traverse() {
	node.TraverseFunc(func(n *Node) {
		n.Print()
	})
	fmt.Println()
}

// 传一个函数，就可以做很多事情了，策略模式
func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}

	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}

func (node *Node) TraverseWithChannel() chan *Node {
	out := make(chan *Node)
	go func() {
		// 函数式编程，策略模式的体现
		node.TraverseFunc(func(node *Node) {
			out <- node
		})
		close(out)
	}()
	return out
}

// 后序遍历
func (node *Node) TraversePost() {
	if node == nil {
		return
	}

	node.Left.TraversePost()
	fmt.Println(node.Value)
	node.Right.TraversePost()
}
