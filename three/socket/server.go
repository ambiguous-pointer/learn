package socket

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func RunSocketServer() {
	// 监听聊天室端口
	port := "8888"
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("聊天室开启成功！正在监听%s端口!\n", port)
	/*
		创建 map 储存连接
		key：conn
		value：name
	*/
	connMap := make(map[net.Conn]string)

	// 开始创建连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		connMap[conn] = conn.RemoteAddr().String()
		notifyAllNewUserLogin(connMap[conn], connMap)
		go HandlerConn(conn, connMap)
	}
}

// HandlerConn 处理每一个连接
func HandlerConn(conn net.Conn, connMap map[net.Conn]string) {
	buf := make([]byte, 1024)
	// 处理连接断开
	defer handlerConnClose(conn, connMap, 1)

	//开启一个携程监听超时
	keepAlive := make(chan bool)
	go func(conn net.Conn) {
		for {
			select {
			case <-keepAlive:
			case <-time.After(1 * time.Minute):
				// 该连接超时
				handlerConnClose(conn, connMap, 2)
				return
			}
		}
	}(conn)

	// 打印帮助列表
	//help(conn)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			return
		}
		keepAlive <- true
		msg := strings.Trim(string(buf[:n]), "\r\n")
		handleMsg(conn, connMap, msg)
	}
}

// 处理连接断开
func handlerConnClose(conn net.Conn, connMap map[net.Conn]string, t int) {
	// t == 1 : 主动断开连接
	// t == 2 : 超时断开连接
	userExit(conn, connMap, t)
	_ = conn.Close()
}

// 提醒所有人新用户上线
func notifyAllNewUserLogin(name string, connMap map[net.Conn]string) {
	for curConn := range connMap {
		msg := time.Now().Format(" 15:04:05 ") + " === [" + name + "] " + "login! === \n"
		_, err := curConn.Write([]byte(msg))
		if err != nil {
			fmt.Printf("提醒[%s]新用户上线失败\n", connMap[curConn])
			continue
		}
	}
}

// 提醒所有人用户下线
func notifyAllNewUserLogout(name string, connMap map[net.Conn]string, t int) {
	for curConn := range connMap {
		msg := ""
		if t == 1 {
			msg = time.Now().Format(" 15:04:05 ") + " === [" + name + "] " + "logout! === \n"
		} else if t == 2 {
			msg = time.Now().Format(" 15:04:05 ") + " === [" + name + "] " + "timeout! === \n"
		}
		_, err := curConn.Write([]byte(msg))
		if err != nil {
			fmt.Printf("提醒[%s]新用户下线失败\n", connMap[curConn])
			continue
		}
	}
}

// 处理连接传递的信息
func handleMsg(conn net.Conn, connMap map[net.Conn]string, msg string) {
	parseArr := strings.Split(msg, "|")
	if len(parseArr) > 1 && parseArr[0] == "func" {
		switch parseArr[1] {
		case "rename":
			{
				if len(parseArr) == 3 {
					rename(conn, connMap, parseArr[2])
				} else {
					syntaxError(conn)
				}
			}
		case "list":
			{
				if len(parseArr) == 2 {
					list(conn, connMap)
				} else {
					syntaxError(conn)
				}
			}
		case "exit":
			{
				if len(parseArr) == 2 {
					userExit(conn, connMap, 1)
				} else {
					syntaxError(conn)
				}
			}
		}
	} else {
		broadcast(conn, connMap, msg)
	}
}

// 广播
func broadcast(conn net.Conn, connMap map[net.Conn]string, msg string) {
	for curConn := range connMap {
		_, err := curConn.Write([]byte(time.Now().Format(" 15:04:05 ") + connMap[conn] + " : " + msg + "\n"))
		if err != nil {
			fmt.Printf("向[%s]发送广播失败\n", connMap[curConn])
			return
		}
	}
}

// 退出登陆
func userExit(conn net.Conn, connMap map[net.Conn]string, t int) {
	name, exist := connMap[conn]
	if exist {
		delete(connMap, conn)
		notifyAllNewUserLogout(name, connMap, t)
	}
}

// 提醒语法错误
func syntaxError(conn net.Conn) {
	_, err := conn.Write([]byte("syntaxError!\n"))
	if err != nil {
		fmt.Printf("向[%s]发送<语法错误>失败\n", conn)
		return
	}
}

// 修改名称
func rename(conn net.Conn, connMap map[net.Conn]string, newName string) {
	_, err := conn.Write([]byte("success!\n"))
	if err != nil {
		fmt.Printf("向[%s]发送<修改名称成功>失败\n", conn)
		return
	}
	connMap[conn] = newName
}

// 获取当前列表
func list(conn net.Conn, connMap map[net.Conn]string) {
	msg := "=====================\n"
	msg += "User List:\n"
	for curConn, name := range connMap {
		msg += name
		if conn == curConn {
			msg += " (me) "
		}
		msg += "\n"
	}
	msg += "=====================\n"
	_, err := conn.Write([]byte(msg))
	if err != nil {
		fmt.Printf("向[%s]发送<当前用户列表>失败\n", conn)
		return
	}
}

// help 帮助
func help(conn net.Conn) {
	msg :=
		`
==================================================================================================
	Introduction:
	This is a easy chat room! 
	You can chat with other people by this client!
	As you input anything, other people in the room can see it, also, you can see their input!
	If you do nothing within 3 minutes, the client will go offline!

	Command List:
	<anything>                      - chat with others
	func|help                       - get some help
	func|rename|<Your New Name>     - change your name
	func|list                       - to see who is online
	func|exit                       - go offline
	
    by Aurora~
=================================================================================================

`
	_, err := conn.Write([]byte(msg))
	if err != nil {
		fmt.Printf("向[%s]发送<语法错误>失败\n", conn)
		return
	}
}
