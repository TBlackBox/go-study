package main

import (
	"fmt"
	"sqliteTest"
)

func init() {

	fmt.Printf("主函数开始初始化操作。。。\n")
}

func main() {

	fmt.Printf("进入main函数，程序开始执行。。。\n")

	sqliteTest.StartGromTest()

	//ginTest.GinTest()

	// ffmpegutil.TestVideo()

	//返回变量占的字节数
	// var a int
	// fmt.Println(unsafe.Sizeof(a))

	//websocket.StartWebsocketServer()

	//测试定时器
	//base.StartTimer();

	//测试 求产生随机数之和 写得还不是很好  还需要不断的优化
	//example.StartGoroutinedDemo(1000, 1)

	//http测试
	// netdemo.StartHTTPService()
	// netdemo.StartHTTPClient()

	//udp 测试
	// netdemo.StartUDPService()
	// netdemo.StartUDPClient()

	//tcp 测试
	// netdemo.StartTCPServer("127.0.0.1", 9090)
	// netdemo.StartTCPClient("127.0.0.1", 9090)

	//基础测试
	//base.Test()

	//data := [...]int{0,1,2,3,4,5,10:9}
	//
	//fmt.Printf("arr: %p ,%v,%d,%d\n",&data,data,len(data),cap(data))
	//
	//slice := data[:3]
	//fmt.Printf("arr: %p ,%v,%d,%d,%p\n",&slice,slice,len(slice),cap(slice),&slice[0])
	//
	//slice2 := data[:2]
	//fmt.Printf("arr: %p ,%v,%d,%d,%p\n",&slice2,slice2,len(slice2),cap(slice2),&slice2[0])
	//
	//slice22 := slice2[:1]
	//fmt.Printf("slice22: %p ,%v,%d,%d,%p\n",&slice22,slice22,len(slice22),cap(slice22),&slice22[0])
	//
	//slice22 = append(slice22, 1,2,3,4,2,3,2,1,2,1,2,2,2,22,2,2,2,2,2,22,2,2,2)
	//fmt.Printf("slice22: %p ,%v,%d,%d,%v,%p\n",&slice22,slice22,len(slice22),cap(slice22),data,&slice22[0])
	//
	//slice222 := slice22[:5]
	//fmt.Printf("slice222: %p ,%v,%d,%d\n",&slice222,slice222,len(slice222),cap(slice222))

}
