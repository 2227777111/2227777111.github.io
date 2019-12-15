package main

import (
	"net/http"
	"strconv"
	"io/ioutil"
)

func HandleFuck (writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("PLEASE FUCK YOURSELF\n呵呵"))
}

func main() {
	http.HandleFunc("/shit", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("shit你妹啊"))
	})

	http.HandleFunc(
		//路由
		"/whoami",

		//路由的处理函数
		//writer 响应的书写器
		//request 用户的请求
		func(writer http.ResponseWriter, request *http.Request) {
			writer.Write([]byte("请求的方法="+request.Method))
			writer.Write([]byte("请求的内容长度="+strconv.Itoa(int(request.ContentLength))))
			writer.Write([]byte("请求的主机="+request.Host))
			writer.Write([]byte("请求的协议="+request.Proto))
			writer.Write([]byte("请求的远程地址="+request.RemoteAddr))
			writer.Write([]byte("请求的路由="+request.RequestURI))
			writer.Write([]byte("朕已收到请求"))
		})

	http.HandleFunc("/fuck",HandleFuck)

	http.HandleFunc("/hupu", func(writer http.ResponseWriter, request *http.Request) {
		huangyiBytes, _ := ioutil.ReadFile("D:/BJBlockChain1803/demos/W2/day5/03HTTP通信/虎扑.html")
		writer.Write(huangyiBytes)
	})

	//http.ListenAndServe("127.0.0.1:8080",nil)
	//允许通过本机的所有IP访问
	//10.15.110.7:8080/shit
	http.ListenAndServe("0.0.0.0:8080",nil)
}
