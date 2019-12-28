package main

/*
√ 对一个文本文件做字符统计，包含字母、数字、空白（'\t',' ','\n','r')和其它；
*/

import (
	"fmt"
	"io/ioutil"
)

//字符串的遍历
func test() {
	str := "hello你妹123"
	var c1 = 'a'
	var c2 = '震'
	fmt.Println(c1, c2)

	for i, v := range str {
		fmt.Printf("序号%d,字符%c,字符的类型是%T\n", i, v, v)
	}
}

//字节切片和字符串可以相互做类型转换
func test2() {
	fmt.Println([]byte("你妹"))
	fmt.Println(string([]byte{228, 189, 160, 229, 166, 185}))
}

func main() {
	bytes, _ := ioutil.ReadFile("D:/BJBlockChain1803/demos/W2/day2/你妹.txt")
	contentStr := string(bytes)
	fmt.Println(contentStr)

	var numCount, letterCount, spaceCount int
	for _, c := range contentStr {
		if c >= '0' && c <= '9' {
			numCount ++
		} else if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			letterCount++
		} else if c == '\n' || c == ' ' || c == '\r' || c == '\t' {
			fmt.Printf("space-c=%c\n",c)
			spaceCount++
		}
	}
	fmt.Println("numCount=", numCount)
	fmt.Println("letterCount=", letterCount)
	fmt.Println("spaceCount=", spaceCount)
}
