package filelisting

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

type userError1 string

func (e userError1) Error() string {
	return e.Message()
}

// 实现了userError接口
func (e userError1) Message() string {
	return string(e)
}

func HandleFileList(writer http.ResponseWriter,
	request *http.Request) error { //返回error,丢给调用此方法的地方处理错误
	fmt.Println()
	if strings.Index(
		request.URL.Path, prefix) != 0 {
		return userError1(fmt.Sprintf("path %s must start "+
			"with %s",
			request.URL.Path, prefix))
	}
	path := request.URL.Path[len(prefix):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(all)
	return nil
}
