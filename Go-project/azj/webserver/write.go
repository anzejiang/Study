package main

import (
	"go-Study/Go-project/azj/webserver/filelisting"
	"log"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request)  {
	return func(writer http.ResponseWriter, request *http.Request) {

	err := handler(writer, request)
	if err != nil {
		// 日志输出
		log.Println(err.Error())
		//log.Output(2, err.Error())

		code := http.StatusOK

		// 对不同的错误进行处理
		switch {
		// 文件不存在的情况
		case os.IsNotExist(err):
			code = http.StatusNotFound
		// 权限不足的情况
		case os.IsPermission(err):
			code = http.StatusForbidden
		// 未知错误
		default:
			code = http.StatusInternalServerError
			//http.Error(writer, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
		http.Error(writer, http.StatusText(code), code)
	}
	}
}

func main() {
	http.HandleFunc("/list/", errWrapper(filelisting.HandleFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil{
		panic(err)
	}
}
