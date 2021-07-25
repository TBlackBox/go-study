package websocket

import (
	"encoding/json"
	"fmt"
)

var h = hub{
	c: make(map[*connection]bool),
	u: make(chan *connection),
	b: make(chan []byte),
	r: make(chan *connection),
}

type hub struct {
	c map[*connection]bool
	b chan []byte
	r chan *connection
	u chan *connection
}

func (h *hub) run() {
	for {
		select {
		case c := <-h.r:
			fmt.Println("从通道和h.r中获取到了数据,登录数据")
			h.c[c] = true
			c.data.IP = c.ws.RemoteAddr().String()
			c.data.Type = "handshake"
			c.data.UserList = userList
			dataB, _ := json.Marshal(c.data)
			c.sc <- dataB
		case c := <-h.u:
			fmt.Println("从通道和h.u中获取到了数据")
			if _, ok := h.c[c]; ok {
				delete(h.c, c)
				close(c.sc)
			}
		case data := <-h.b:
			fmt.Println("从通道和h.b中获取到了数据")
			for c := range h.c {
				select {
				case c.sc <- data:
				default:
					delete(h.c, c)
					close(c.sc)
				}
			}
		}
	}
}
