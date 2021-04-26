package netdemo

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

//TCP 服务端
////////////////////////////////////////////////////////////////////////////////
func process(conn net.Conn) {
	//关闭链接
	defer conn.Close()

	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) //读取数据
		if err != nil {
			fmt.Println("读取客户端数据失败，err:", err)
			break
		}

		recvStr := string(buf[:n])
		fmt.Println("收到客服端数据：", recvStr)
		conn.Write([]byte("收到客户端发送的数据为：" + recvStr))
	}
}

//StartTCPServer 始监听服务端
func StartTCPServer(ip string, port int) {
	url := ip + ":" + strconv.Itoa(port)
	fmt.Println("监听地址：", url)
	listen, err := net.Listen("tcp", url)
	if err != nil {
		fmt.Println("接受失败：", err)
		return
	}

	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn) // 启动一个goroutine处理连接
	}

}

//////////////////////////tcp 服务端完//////////////////////////////////////////////////////

//////////////////////////tcp 客服端开始//////////////////////////////////////////////////////

//StartTCPClient tcp 客服端
func StartTCPClient(ip string, port int) {
	conn, err := net.Dial("tcp", ip+":"+strconv.Itoa(port))
	if err != nil {
		fmt.Println("err :", err)
		return
	}
	defer conn.Close() // 关闭连接
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n') // 读取用户输入
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" { // 如果输入q就退出
			return
		}
		_, err = conn.Write([]byte(inputInfo)) // 发送数据
		if err != nil {
			return
		}
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed, err:", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}

//////////////////////////tcp 客服端完//////////////////////////////////////////////////////
