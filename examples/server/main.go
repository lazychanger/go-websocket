package main

import (
	"github.com/lazychanger/go-websocket"
	"log"
	"net"
	"net/http"
)

type ExampleWSServer struct {
	Name string
}

func (e *ExampleWSServer) OnOpen(ctx *websocket.Context) {
	log.Printf("[wsserver %s] on open", e.Name)
}

func (e *ExampleWSServer) OnMessage(ctx *websocket.Context, body []byte) {
	log.Printf("[wsserver %s] on message: %s", e.Name, string(body))
}

func (e *ExampleWSServer) OnClose(ctx *websocket.Context) {
	log.Printf("[wsserver %s] on close", e.Name)
}

func (e *ExampleWSServer) OnConnect(ctx *websocket.Context, writer http.ResponseWriter, request *http.Request) error {
	log.Printf("[wsserver %s] on connect", e.Name)
	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panicln(err)
	}
	mux := &http.ServeMux{}
	mux.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		ws := websocket.NewWebsocket(writer, request, &ExampleWSServer{
			Name: "ws-v1",
		})
		ws.Start()
	})
	mux.HandleFunc("/ws-v2", func(writer http.ResponseWriter, request *http.Request) {
		ws := websocket.NewWebsocket(writer, request, &ExampleWSServer{
			Name: "ws-v2",
		})
		ws.Start()
	})
	log.Println("server start")
	if err := http.Serve(listener, mux); err != nil {
		log.Println(err)
	}
}
