package tree

import "fmt"

type Node struct {
	Value int
	Left, Right *Node
}

func CreateNode (value int) *Node {
	return &Node{Value: value}
}

func (node Node) Print(){
	fmt.Print(node.Value, " ")
}

func (node *Node) SetValue(value int){
	if node == nil {
		fmt.Println("setting value to nil node.")
		return
	}
	node.Value = value
}





