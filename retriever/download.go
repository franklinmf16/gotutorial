package main

import (
	"fmt"
	"mygo/learngo/retriever/mock"
	real2 "mygo/learngo/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string  {
	return r.Get("http://www.imooc.com")
}

func main(){
	var r Retriever
	r = mock.Retriever{"this is good demo"}
	r2 := real2.Retriever{}
	fmt.Println(download(r))
	fmt.Println(download(r2))
}

