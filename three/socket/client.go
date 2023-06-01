package socket

import (
	"fmt"
	"net"
	"strings"
)

func RunSocketClient() {
	host := "localhost"
	port := "8888"
	// 创建连接
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("已连接到", conn.RemoteAddr())
	// 监听服务端消息
	go func() {
		buf := make([]byte, 5*1024)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				return
			}
			msg := string(buf[:n])
			fmt.Print(msg)
		}
	}()

	// 向服务端发送信息
	for {
		var msg string
		_, _ = fmt.Scanf("%s", &msg)
		msg = strings.Trim(msg, "\r\n")
		if "exit" == msg {
			fmt.Printf("close connect...\n")
			return
		}
		_, _ = conn.Write([]byte(msg))
	}
}
