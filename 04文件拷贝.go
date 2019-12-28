package main

import (
	"io/ioutil"
	"fmt"
	"io"
	"os"
	"bufio"
)

/*
√ 使用ioutil包做一个傻瓜式拷贝
√ 使用io.Copy进行文件拷贝
√ 使用缓冲1K的缓冲区配合缓冲读写器进行图片拷贝
*/

func main041() {
	bytes, _ := ioutil.ReadFile("D:/iWorkspace/TeachingSkills/imgs2/fuckoff2.jpg")
	err := ioutil.WriteFile("D:/iWorkspace/TeachingSkills/imgs2/fuckoff2222222.jpg", bytes, 0666)
	if err == nil {
		fmt.Println("拷贝成功！")
	} else {
		fmt.Println("拷贝失败，err=", err)
	}
}

/*使用io.Copy进行文件拷贝*/
func main042() {

	//打开拷贝源文件，注意模式为只读模式
	srcFile, _ := os.OpenFile("D:/iWorkspace/TeachingSkills/imgs2/fuckoff2.jpg", os.O_RDONLY, 0666)
	//打开要拷贝到的目标文件，注意模式：创建+写出
	dstFile, _ := os.OpenFile("D:/iWorkspace/TeachingSkills/imgs2/fuckoff2333.jpg", os.O_WRONLY|os.O_CREATE, 0666)

	//执行源文件到目标文件的拷贝
	written, err := io.Copy(dstFile, srcFile)

	//判断执行结果
	if err == nil {
		fmt.Println("拷贝成功，字节数=", written)
	} else {
		fmt.Println("拷贝失败，err=", err)
	}
}

/*使用缓冲1K的缓冲区配合缓冲读写器进行图片拷贝*/
func main() {
	//打开源文件
	srcFile, _ := os.OpenFile("D:/iWorkspace/TeachingSkills/imgs2/dragon.gif", os.O_RDONLY, 0666)
	//打开目标文件
	dstFile, _ := os.OpenFile("D:/iWorkspace/TeachingSkills/imgs2/my_dragon.gif", os.O_WRONLY|os.O_CREATE, 0666)
	//延时关闭他们
	defer func() {
		srcFile.Close()
		dstFile.Close()
		fmt.Println("文件全部关闭！")
	}()

	//创建源文件的缓冲读取器
	reader := bufio.NewReader(srcFile)
	//创建目标文件的写出器
	writer := bufio.NewWriter(dstFile)

	//创建小水桶（缓冲区）
	buffer := make([]byte, 1024)

	//一桶一桶地读入数据到水桶（缓冲区），直到io.EOF
	for{
		_, err := reader.Read(buffer)
		if err != nil{
			if err == io.EOF{
				fmt.Println("源文件读取完毕！")
				break
			}else{
				fmt.Println("读取源文件发生错误，err=",err)
				return
			}
		}else{
			//将每桶数据写出到目标文件
			_, err := writer.Write(buffer)
			if err !=nil{
				fmt.Println("写出错误，err=",err)
				return
			}
		}
	}

	fmt.Println("拷贝完毕！")
}

