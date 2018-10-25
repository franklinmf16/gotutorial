package main

import "fmt"

func upodateSlice(s []int){
	s[0] = 100
}

func printSlice (slice []int){
	fmt.Printf("%v, len=%d, cap=%d\n", slice,len(slice), cap(slice))
}
func main()  {
	arr := [...]int{0,1,2,3,4,5,6,7}
	s := arr[2:6]
	fmt.Println(arr)
	fmt.Println(s)
	fmt.Println(arr[2:])

	fmt.Println("after update function")
	upodateSlice(arr[1:])
	fmt.Println(arr)

	s1 := arr[2:6]
	fmt.Println(s1)
	s2 := make([]int, 16)
	printSlice(s2)
	s3 := make([]int, 10, 32)
	printSlice(s3)


	printSlice(s1)
	fmt.Println("copying slice")
	copy(s2, s1)
	fmt.Println(s2)
	fmt.Println(s1)
	fmt.Println("deleting elements from slice")
	printSlice(	s2[:3])
	printSlice(	s2[4:])
	printSlice(append(s2[:3], s2[4:]...))

	fmt.Println("popping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front, s2)

	fmt.Println("popping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(tail, s2)



}