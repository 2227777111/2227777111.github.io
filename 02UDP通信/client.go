package main

import (
	"fmt"
	"os"
	"net"
	"bufio"
)

/*
·拨号连接服务端主机的8888端口，建立UDP连接
·循环从标准输入（命令行）读取一行用户输入，向服务端发送
·接收并打印服务端消息，如果消息是“bye”，就退出程序
*/
func main() {
	conn, e := net.Dial("udp", "127.0.0.1:8888")
	ClientHandleError(e,"net.Dial")

	reader := bufio.NewReader(os.Stdin)
	buffer := make([]byte, 1024)
	for {
		lineBytes, _, _:= reader.ReadLine()
		conn.Write(lineBytes)
		n, _ := conn.Read(buffer)
		serverMsg := string(buffer[:n])
		fmt.Println("服务端：",serverMsg)

		if serverMsg == "fuckoff"{
			//发出杀猪般的哀鸣
			conn.Write([]byte("呜呼狡兔死走狗烹飞鸟尽良弓藏，吾去也！"))
			break
		}
	}
	fmt.Println("GAME OVER!")
}

func ClientHandleError(err error, when string) {
	if err != nil {
		fmt.Println(err, when)
		os.Exit(1)
	}
}
