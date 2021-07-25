package websocket

import (
	"encoding/json"
	"fmt"
	"mylog"
	"net/http"

	"github.com/gorilla/websocket"
)

//定义一个链接结构体
type connection struct {
	ws   *websocket.Conn
	sc   chan []byte
	data *Data
}

var wu = &websocket.Upgrader{ReadBufferSize: 512,
	WriteBufferSize: 512, CheckOrigin: func(r *http.Request) bool { return true }}

func myws(w http.ResponseWriter, r *http.Request) {

	fmt.Println("接受到一个请求")

	ws, err := wu.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	c := &connection{sc: make(chan []byte, 256), ws: ws, data: &Data{}}
	//将链接发送到通道
	h.r <- c
	//启动一个写数据的协程
	go c.writer()
	//主协程用来读数据
	c.reader()
	defer func() {
		c.data.Type = "logout"
		userList = del(userList, c.data.User)
		c.data.UserList = userList
		c.data.Content = c.data.User
		dataB, _ := json.Marshal(c.data)
		h.b <- dataB
		h.r <- c
	}()
}

//启动一个写数据的协程
func (c *connection) writer() {
	for message := range c.sc {
		c.ws.WriteMessage(websocket.TextMessage, message)
	}
	c.ws.Close()
}

var userList = []string{}

//读数据
func (c *connection) reader() {
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			mylog.Println("未读到消息：", c.data.User)
			h.r <- c
			break
		}

		mylog.Println("读取的消息为：", string(message))

		json.Unmarshal(message, &c.data)
		switch c.data.Type {
		case "login":
			c.data.User = c.data.Content
			c.data.From = c.data.User
			userList = append(userList, c.data.User)
			c.data.UserList = userList
			dataB, _ := json.Marshal(c.data)
			h.b <- dataB
		case "user":
			c.data.Type = "user"
			dataB, _ := json.Marshal(c.data)
			h.b <- dataB
		case "logout":
			c.data.Type = "logout"
			userList = del(userList, c.data.User)
			dataB, _ := json.Marshal(c.data)
			h.b <- dataB
			h.r <- c
		default:
			fmt.Print("========default================")
		}
	}
}

func del(slice []string, user string) []string {
	count := len(slice)
	if count == 0 {
		return slice
	}
	if count == 1 && slice[0] == user {
		return []string{}
	}
	var nSlice = []string{}
	for i := range slice {
		if slice[i] == user && i == count {
			return slice[:count]
		} else if slice[i] == user {
			nSlice = append(slice[:i], slice[i+1:]...)
			break
		}
	}
	fmt.Println(nSlice)
	return nSlice
}
