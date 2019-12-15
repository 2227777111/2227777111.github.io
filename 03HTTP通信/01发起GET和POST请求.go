package main

import (
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)

/*
·发起百度搜索的GET请求："http://www.baidu.com/s?wd=肉丝"，打印回复的内容
·对https://httpbin.org/post发起post请求，携带表单数据，键值自拟，打印回复的内容
·表单数据类型 application/x-www-form-urlencoded，表单读取API：strings.NewReader("rmb=0.5&hobby=接客和约汉"))
*/

func mainGet() {
	resp, err := http.Get("http://www.baidu.com/s?wd=去浪")
	ClientHandleError(err,"http.Get")
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bytes))
}

func main() {
	resp, err := http.Post(
		"https://httpbin.org/post?name=jack&teacher=john",
		"application/x-www-form-urlencoded",
		strings.NewReader("姓名=接客&师承=约汉&心愿=男人脑壳全坏掉"))

	ClientHandleError(err,"http.Get")
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bytes))
}

func ClientHandleError(err error, when string) {
	if err != nil {
		fmt.Println(err, when)
		os.Exit(1)
	}
}