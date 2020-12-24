package main

import (
	"fmt"
	"go-Study/Go-project/retriever/mockretriever"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string  {
	return r.Get("www.imooc.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{"The is a fake Content"}
	fmt.Println(download(r))
}
