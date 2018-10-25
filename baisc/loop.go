package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n/=2  {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile (filename string){
	file, err := os.Open(filename)
	if err != nil{
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}

}

func swap(a,b int) (outA, outB int){
	return b, a
}

func swapByRef(a, b *int)  {
	*b, *a = *a, *b
}




func main()  {
	fmt.Println(
			convertToBin(5), // 101
			convertToBin(9),
			convertToBin(13),
	)

	printFile("baisc/abc.txt")

	a, b := 2,3
	outA, outB := swap(a, b)
	fmt.Println(outA,outB)

	swapByRef(&a, &b)
	fmt.Println(a,b)
}

