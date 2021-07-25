package netdemo

import (
	"fmt"
	"net"
)

//StartUDPService 启动UDP 服务端
func StartUDPService() {
	fmt.Println("启动udp服务端程序...")
	listen, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4(0, 0, 0, 0), Port: 30000})
	if err != nil {
		fmt.Println("监听失败，err", err)
	}

	defer listen.Close()
	for {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("接受udp数据错误，err:", err)
			continue
		}

		fmt.Printf("data:%v addr:%v count:%v\n", string(data[:n]), addr, n)

		response := []byte("服务的收到数据:" + string(data[:n]))
		_, err = listen.WriteToUDP(response, addr) // 发送数据
		if err != nil {
			fmt.Println("write to udp failed, err:", err)
			continue
		}

	}

}

//StartUDPClient 启动UDP客服端
func StartUDPClient() {
	fmt.Println("启动udp客服端程序。。。")
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("连接服务端失败，err:", err)
		return
	}
	defer socket.Close()
	sendData := []byte("Hello server")
	_, err = socket.Write(sendData) // 发送数据
	if err != nil {
		fmt.Println("发送数据失败，err:", err)
		return
	}
	fmt.Println("发送数据成功")
	data := make([]byte, 4096)
	n, remoteAddr, err := socket.ReadFromUDP(data) // 接收数据
	if err != nil {
		fmt.Println("接收数据失败，err:", err)
		return
	}
	fmt.Printf("recv:%v addr:%v count:%v\n", string(data[:n]), remoteAddr, n)
}
