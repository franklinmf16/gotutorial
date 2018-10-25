package main

import "fmt"

func main() {
	m := map[string]string{
		"name": "ccmouse",
		"course": "golang",
		"site": "imooc",
		"quality": "good",
	}

	m2 := make(map[string]string)

	for key, value := range m {
		fmt.Println(key,value)
	}

	fmt.Println(m2)

	courseName, ok := m["name"]
	fmt.Println(courseName, ok)
	delete(m,"name")

	for key, value := range m {
		fmt.Println(key,value)
	}

}
