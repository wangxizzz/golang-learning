package mock

import "fmt"

type RetrieverMock struct {
	Contents string
}

// 实现了Stringer接口的String(),定义了指针接收者
func (r *RetrieverMock) String() string {
	return fmt.Sprintf(
		"Retriever: {Contents=%s}", r.Contents)
}

func (r *RetrieverMock) Post(url string,
	form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}

func (r *RetrieverMock) Get(url string) string {
	return r.Contents
}
