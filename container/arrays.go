package main

import "fmt"

func printList(array *[5]int){
	array[1] = 100
	for i:= range array{
		fmt.Println(array[i])
	}
}

func main()  {
	var arr1 [5]int
	arr2 := [3]int{1,2,3}
	arr3 := [...]int{2,4,6,8,10}
	var grid [4][5]int

	fmt.Println(arr1,arr2,arr3,grid)


	for _,v:= range arr3{
		fmt.Println(v)
	}

	fmt.Println("here is the print list function")
	printList(&arr3)

	fmt.Println("here is the print list without function")
	for i:=0; i< len(arr3);i++  {
		fmt.Println(arr3[i])
	}



	
}
