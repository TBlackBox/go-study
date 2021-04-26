package websocket

import (
	"fmt"
	"mylog"
	"net/http"

	"github.com/gorilla/mux"
)

//StartWebsocketServer websocket 服务端
func StartWebsocketServer() {

	mylog.Println("开始启动websocket...")
	router :=  .NewRouter()
	go h.run()
	router.HandleFunc("/ws", myws)
	if err := http.ListenAndServe("127.0.0.1:8080", router); err != nil {
		fmt.Println("err:", err)
	}
}
