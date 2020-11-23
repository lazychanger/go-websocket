package main

import (
	"github.com/gorilla/websocket"
	"log"
)

func main() {
	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws1", nil)
	if err != nil {
		log.Print(err)
		return
	}
	c.WriteJSON(map[string]interface{}{"hello": "world"})
	c.Close()
}
