package main

import (
	"fmt"
	"math"
)

func variableZeroValue(){
	var a int
	var s string
	fmt.Printf("%d %q", a,s)
}


func variableInitialValue(){
	var a, b int = 3,4
	var s string = "abc"
	fmt.Println(a,b,s)
}

func enums(){
	const(
		cpp = iota
		java
		python
		golang
		javascript
	)
	fmt.Println(cpp,java)
}

const filename = "abc.txt"

func consts(){
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)

}

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	consts()
	enums()
}













