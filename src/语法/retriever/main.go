package main

import (
	"fmt"
	"语法/retriever/mock"
	"语法/retriever/real"

	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string,
		form map[string]string) string
}

const url = "http://www.imooc.com"

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name":   "ccmouse",
			"course": "golang",
		})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}

func main() {
	var r Retriever // retriever 中文：实现者
	// 传的是值。
	mockRetriever := mock.RetrieverMock{Contents: "this is a fake imooc.com"}
	fmt.Printf("%T \n", mockRetriever)
	/**
	在利用接口的实现，来初始化接口变量r时，如果接口的实现定义的是指针接收者，那么初始化时，需要取实现的地址
	*/
	r = &mockRetriever
	fmt.Printf("%T \n", r)
	//inspect(&mockRetriever)
	// 传的是地址
	r = &real.RetrieverReal{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	fmt.Printf("%T \n", r)
	inspect(r)

	// Type assertion(断言)
	if mockRetriever, ok := r.(*mock.RetrieverMock); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("r is not a mock retriever")
	}

	fmt.Println(
		"Try a session with mockRetriever")
	fmt.Println(session(&mockRetriever))
}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf(" > Type:%T Value:%v\n", r, r)
	fmt.Print(" > Type switch: ")
	switch v := r.(type) {
	case *mock.RetrieverMock:
		fmt.Println("Contents:", v.Contents)
	case *real.RetrieverReal:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}
