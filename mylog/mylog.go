//Package mylog 第三方的日志库，如logrus、zap等。
package mylog

import (
	"fmt"
	"log"
)

func init() {

	//输出到某个日志文件下
	// logFile, err := os.OpenFile("./test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	// if err != nil {
	// 	fmt.Println("open log file failed, err:", err)
	// 	return
	// }
	// log.SetOutput(logFile)

	//设置标志
	flag := log.Flags() //获取flag 选项
	fmt.Println("flag:", flag)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)

	//设置前缀
	prefix := log.Prefix()
	fmt.Println("orefix:", prefix)
	log.SetPrefix("[mylog]")

}

// Print 打印
func Print(v ...interface{}) {
	log.Print(v...)
}

// Printf 打印
func Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

// Println 打印
func Println(v ...interface{}) {
	log.Println(v...)
}

// Fatal 打印
func Fatal(v ...interface{}) {
	log.Fatal(v...)
}

// Fatalf 打印
func Fatalf(format string, v ...interface{}) {
	log.Fatalf(format, v...)
}

// Fatalln 打印
func Fatalln(v ...interface{}) {
	log.Fatalln(v...)
}

// Panic 打印
func Panic(v ...interface{}) {
	log.Panic(v...)
}

// Panicf 打印
func Panicf(format string, v ...interface{}) {
	log.Panicf(format, v...)
}

// Panicln 打印
func Panicln(v ...interface{}) {
	log.Panicln(v...)
}
