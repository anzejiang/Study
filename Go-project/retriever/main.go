package main

import (
	"fmt"
	mock "go-Study/Go-project/retriever/mockretriever"
	"go-Study/Go-project/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string  {
	return r.Get("http://www.imooc.com")
}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent", v.UserAgent)
	}
}

func main() {
	var r Retriever
	r = mock.Retriever{"The is a fake Content"}
	inspect(r)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut: time.Minute,
	}

	//fmt.Println(download(r))
	inspect(r)

	if mokeRetriever, ok := r.(mock.Retriever); ok {
		fmt.Println(mokeRetriever.Contents)
	}else {
		fmt.Println("not a mock retriever")
	}

}

