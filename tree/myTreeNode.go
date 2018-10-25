package tree

type MyTreeNode struct {
	Node *Node
}

func (Node *MyTreeNode) PostOrder(){

	if Node == nil || Node.Node == nil {
		return
	}

	left := MyTreeNode{Node.Node.Left}
	right := MyTreeNode{Node.Node.Right}

	left.PostOrder()
	right.PostOrder()
	Node.Node.Print()
}