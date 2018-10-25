package main

import "fmt"

func main() {
	s := "Yes 我爱imooc!"
	//fmt.Printf("%X\n", []byte(s))

	for _, b := range []byte(s) {
		fmt.Printf("%X ",b)
	}
	fmt.Println()

}