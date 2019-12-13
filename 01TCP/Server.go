package main

import (
	"net"
	"fmt"
	"os"
)

/*
·服务端在本机的8888端口建立TCP监听
·为接入的每一个客户端开辟一条独立的协程
·循环接收客户端消息，不管客户端说什么，都自动回复“已阅xxx”
·如果客户端说的是“im off”，则回复“bye”并断开其连接
*/

func ServerHandleError(err error, when string) {
	if err != nil {
		fmt.Println(err, when)
		os.Exit(1)
	}
}

func main() {

	//服务端在本机的8888端口建立TCP监听
	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	ServerHandleError(err, "net.Listen")

	for {
		/*循环接入所有客户端，得到专线连接*/
		conn, e := listener.Accept()
		ServerHandleError(e, "listener.Accept()")

		/*开辟一条独立协程，与该客户端聊天*/
		go ChatWith(conn)
	}

}

/*在Conn中与客户端对话*/
func ChatWith(conn net.Conn) {
	//创建消息缓冲区
	buffer := make([]byte, 1024)

	/*不断收发消息*/
	for {
		/*一个完整的消息回合*/

		//读取客户端发来的消息，存入缓冲区，消息长度为n字节
		n, err := conn.Read(buffer)
		ServerHandleError(err, "conn.Read(buffer)")

		//转换为字符串输出
		clientMsg := string(buffer[0:n])
		fmt.Printf("收到%v的消息:%s\n", conn.RemoteAddr(), clientMsg)

		//回复客户端消息：已阅或bye，bye时跳出会话
		if clientMsg != "im off" {
			conn.Write([]byte("已阅:" + clientMsg))
		} else {
			conn.Write([]byte("bye"))
			break
		}

	}

	//断开客户端连接
	conn.Close()
	fmt.Printf("客户端%s已经断开连接\n", conn.RemoteAddr())

}
