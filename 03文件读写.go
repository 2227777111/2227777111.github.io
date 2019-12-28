package main

import (
	"os"
	"fmt"
	"time"
	"bufio"
	"io"
	"io/ioutil"
)

/*
√ 简易方式打开一个文件，拿着一顿骚操作，一秒后关闭；
√ 以只读方式打开一个文件，创建其带缓冲的读取器，逐行读取到末尾；
√ 使用ioutil包对《一些逼嗑》进行简易读取
√ 以创写追加或创写覆盖方式打开一个文件，缓冲式写出几行数据，倒干缓冲区后退出；
√ 使用ioutil包进行简易文件写出
√ 使用os包的状态检测结合os.IsNotExist(err)判断文件是否存在
*/

func main031() {
	file, err := os.Open("D:/iWorkspace/TeachingSkills/er些逼嗑.txt")
	if err == nil {
		fmt.Println("文件打开成功")
	} else {
		fmt.Println("文件打开失败，err=", err)
		return
	}
	defer func() {
		file.Close()
		fmt.Println("文件已关闭")
	}()

	fmt.Println("拿着文件一顿骚操作", file)
	time.Sleep(1 * time.Second)
}

/*
以只读方式打开一个文件，创建其带缓冲的读取器，逐行读取到末尾
4=readable,2=writable,1=executable
6=4+2
*/
func main032() {
	//以只读模式打开文件(0666代表所有人都有读写权限)
	file, err := os.OpenFile("D:/iWorkspace/TeachingSkills/一些逼嗑.txt", os.O_RDONLY, 0666)

	//判断文件打开是否成功
	if err == nil {
		fmt.Println("文件打开成功")
	} else {
		fmt.Println("文件打开失败，err=", err)
		return
	}

	//延时(函数返回前)关闭文件
	defer func() {
		file.Close()
		fmt.Println("文件已关闭")
	}()

	//创建该文件的缓冲读取器
	reader := bufio.NewReader(file)

	//循环读入数据
	for {
		//每次读入一行
		str, err := reader.ReadString('\n')

		//判断读入是否成功
		if err == nil {
			//打印读入的字符串
			fmt.Println(str)
		} else {
			if err == io.EOF {
				//已到文件末尾,跳出读取循环
				fmt.Println("已到文件末尾！")
				break
			} else {
				//有其它异常，打印异常并结束程序
				fmt.Println("读取失败,err=", err)
				return
			}
		}
	}
	fmt.Println("文件读取完毕！")

	//函数返回前会先执行defer引导的语句
}

/*使用ioutil包对《一些逼嗑》进行简易读取*/
func main033() {
	//读入指定文件的全部数据,返回[]byte类型的原始数据
	bytes, err := ioutil.ReadFile("D:/iWorkspace/TeachingSkills/二些逼嗑.txt")
	if err == nil {
		contentStr := string(bytes)
		fmt.Println(contentStr)
	} else {
		fmt.Println("读取失败,err=", err)
	}
}

/*以【创-写-追加】或【创-写-覆盖】方式打开一个文件，缓冲式写出几行数据，倒干缓冲区后退出；*/
func main034() {
	//打开指定文件，模式：不存在就创建+只写+追加，生成的文件的权限是-rw-rw-rw-
	file, err := os.OpenFile("D:/iWorkspace/TeachingSkills/三些逼嗑.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	//如果有异常直接打印退出
	if err != nil {
		fmt.Println("文件打开失败，err=", err)
		return
	}

	//延时关闭文件
	defer func() {
		file.Close()
		fmt.Println("文件已关闭")
	}()

	//创建文件的缓冲写出器
	writer := bufio.NewWriter(file)

	//分批次地写出一些数据
	writer.WriteString("当代四大屌丝\n")
	writer.WriteString("悔创阿里杰克马\n")
	writer.WriteString("一无所有王健林\n")
	writer.WriteString("不识妻美刘强东\n")
	writer.WriteString("普通家庭马化腾\n")
	writer.WriteRune('你')
	writer.WriteRune('妹')
	writer.WriteByte(123)
	writer.Write([]byte{123, 124, 125, 126})

	//倒干缓冲区(立刻将最后一桶数据写出到文件)
	writer.Flush()
	fmt.Println("写出完毕！")
}

/*使用ioutil包进行简易文件写出*/
func main035() {
	//反引号代表保留原始排版的字符串
	dataStr := `想当年，
	老夫拳打南山敬老院，
		脚踢北海幼儿园，
			我骄傲了吗？！`
	fmt.Printf("dataStr's type=%T，value=%v\n", dataStr, dataStr)

	//将数据字符串转换为原始字节切片
	dataBytes := []byte(dataStr)
	fmt.Printf("dataBytes's type=%T，value=%v\n", dataBytes, dataBytes)

	//向指定文件中写出上面诗篇中的字节数据
	err := ioutil.WriteFile("D:/iWorkspace/TeachingSkills/四些逼嗑.txt", dataBytes, 0666)

	//判断写出是否成功
	if err != nil {
		fmt.Println("托楠哥的福发生了一些错误，err=", err)
	} else {
		fmt.Println("楠哥对不起，写出成功了！~")
	}
}

/*使用os包的状态检测结合os.IsNotExist(err)判断文件是否存在*/
func main() {
	//获取指定文件的信息
	fileInfo, err := os.Stat("D:/iWorkspace/TeachingSkills/一些逼嗑.txt")
	if err != nil {
		fmt.Println("err=", err)

		//返回true代表错误为【文件不存在错误】
		if os.IsNotExist(err){
			fmt.Println("文件不存在！")
		}
	} else {
		fmt.Println("文件存在！")
		fmt.Println(fileInfo)
	}
}
