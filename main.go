package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

var ch = make(chan string)

func WirthFile(msg string) {
	str, _ := os.Getwd()
	file, err := os.OpenFile(str+"\\log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	_, err = file.WriteString(msg)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateServer() {
	ServerSocket, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer ServerSocket.Close()
	fmt.Printf("\u001B[34m\u001B[45m Listen %s \u001B[0m \n", ServerSocket.Addr())
	ch <- ServerSocket.Addr().String()
	WirthFile(ServerSocket.Addr().String() + "\n")
	for {
		// Wait for a connection.
		conn, err := ServerSocket.Accept()
		if err != nil {
			log.Fatal(err)
		}
		m := make(map[string]net.Conn)
		go func(c net.Conn) {
			buf := make([]byte, 5*1024)
			m[c.RemoteAddr().String()] = c
			for {
				n, err2 := c.Read(buf)
				if err2 != nil {
					fmt.Printf("\u001B[31m\u001B[40m 收到 %s \u001B[0m \n", "出现问题了")
					break
				}
				fmt.Printf("\u001B[34m\u001B[45m 收到了来自客服端%s的消息\n内容是 \u001B[0m\n%s", c.RemoteAddr(), strings.Trim(string(buf[:n]), "\r\n"))
				WirthFile("收到了来自客服端的消息->" + strings.Trim(string(buf[:n]), "\r\n") + "\n")
				_, e3 := c.Write([]byte(time.Now().Format(" 15:04:05 ") + strings.Trim(string(buf[:n]), "\r\n") + "\n"))
				if e3 != nil {
					fmt.Printf("\u001B[31m\u001B[40m 回 %s \u001B[0m \n", "出现问题了")
					continue
				}
			}
		}(conn)
	}
}
func main() {
	go CreateServer()
	time.Sleep(3 * time.Second)
	v := <-ch
	c, err := net.Dial("tcp", v)
	close(ch)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\u001B[34m\u001B[45m join %s \u001B[0m \n", c.RemoteAddr())
	go func() {
		buf := make([]byte, 5*1024)
		for {
			n, err := c.Read(buf)
			if err != nil {
				return
			}
			msg := string(buf[:n])
			fmt.Printf("\u001B[35m\u001B[45m 收到了来自服务器%s的消息\n内容是 \u001B[0m\n%s", c.RemoteAddr(), msg)
		}
	}()
	// 向服务端发送信息
	for {
		var msg string
		_, _ = fmt.Scanf("%s", &msg)
		msg = strings.Trim(msg, "\r\n")
		_, _ = c.Write([]byte(msg))
	}
}
