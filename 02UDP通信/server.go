package main

import (
	"net"
	"fmt"
	"os"
)

/*
·服务端在本机的8888端口建立UDP监听，得到广口连接
·循环接收客户端消息，不管客户端说什么，都自动回复“已阅xxx”
·如果客户端说的是“im off”，则回复“bye”
*/
func main() {
	//解析得到UDP地址
	udpAddr, err := net.ResolveUDPAddr("udp", "localhost:8888")
	ServerHandleError(err, "net.ResolveUDPAddr")

	//建立UDP监听，得到广口连接
	udpConn, err := net.ListenUDP("udp", udpAddr)
	ServerHandleError(err, "net.ListenUDP")

	//创建消息缓冲区
	buffer := make([]byte, 1024)

	//从广口连接中源源不断地读取（来自任何客户端的）数据包
	for {
		//读取一个数据包到消息缓冲区，同时获得该数据包的客户端信息
		n, remoteAddr, _ := udpConn.ReadFromUDP(buffer)

		//打印数据包消息内容
		clientMsg := string(buffer[:n])
		fmt.Printf("收到来自%v的消息：%s\n", remoteAddr, clientMsg)

		//回复该数据包的客户端
		if clientMsg != "im off"{
			udpConn.WriteToUDP([]byte("已阅："+clientMsg), remoteAddr)
		}else{
			udpConn.WriteToUDP([]byte("fuckoff"), remoteAddr)
		}
	}

}

func ServerHandleError(err error, when string) {
	if err != nil {
		fmt.Println(err, when)
		os.Exit(1)
	}
}
