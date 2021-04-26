package netdemo

import (
	"fmt"
	"io"
	"net/http"
)

//StartHTTPService 开始http 服务端
func StartHTTPService() {
	//http://127.0.0.1:8000/go
	// 单独写回调函数
	http.HandleFunc("/go", myHandler)
	//http.HandleFunc("/ungo",myHandler2 )
	// addr：监听的地址
	// handler：回调函数
	http.ListenAndServe("127.0.0.1:8000", nil)
}

// handler函数
func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr, "连接成功")
	// 请求方式：GET POST DELETE PUT UPDATE
	fmt.Println("method:", r.Method)
	// /go
	fmt.Println("url:", r.URL.Path)
	fmt.Println("header:", r.Header)
	fmt.Println("body:", r.Body)
	// 回复
	w.Write([]byte("http请求响应数据"))
}

/////////////////////////////////////////////////////

//StartHTTPClient 开始http客服端
func StartHTTPClient() {
	//resp, _ := http.Get("http://www.baidu.com")
	//fmt.Println(resp)
	resp, err := http.Get("http://127.0.0.1:8000/go")
	if err != nil {
		fmt.Println("http get 请求错误，err:", err)
	}
	defer resp.Body.Close()
	// 200 OK
	fmt.Println(resp.Status)
	fmt.Println(resp.Header)

	buf := make([]byte, 1024)
	for {
		// 接收服务端信息
		n, err := resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		} else {
			fmt.Println("读取完毕")
			res := string(buf[:n])
			fmt.Println(res)
			break
		}
	}
}
