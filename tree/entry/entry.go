package main

import ("fmt"
	 	"mygo/learngo/tree"
)

func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5,nil,nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(5)
	root.Left.Right.SetValue(3)
	fmt.Println("new value is ", root.Left.Right.Value)
	fmt.Println(root)

	fmt.Println("traverse")
	root.Traverse()


	fmt.Println("print node")
	root.Print()
	root.SetValue(100)

	fmt.Println("root pointer try")
	var pRoot *tree.Node

	pRoot.SetValue(200)
	pRoot = & root
	pRoot.Print()
	root.Print()
	fmt.Println()
	root.Traverse()

	
	fmt.Println()
	myRoot := tree.MyTreeNode{pRoot}
	myRoot.PostOrder()

}