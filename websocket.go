package websocket

import (
	"github.com/gorilla/websocket"
	"log"
)

type Websocket struct {
	Id      string
	Conf    *Config
	conn    *websocket.Conn
	cfgSets []SetConfig
	serv    Server
	isClose bool
}

func (ws *Websocket) Start() {
	for _, set := range ws.cfgSets {
		set(ws.Conf)
	}

	ctx := &Context{
		ws: ws,
	}

	ws.conn.SetCloseHandler(func(code int, text string) error {
		ws.isClose = true
		ws.serv.OnClose(ctx)
		log.Println("close Handler", code, text)
		return nil
	})
	ws.conn.SetPingHandler(func(appData string) error {
		log.Print("ping handler", appData)
		return nil
	})

	ws.conn.SetPongHandler(func(appData string) error {
		log.Println("pong handler", appData)
		return nil
	})

	ws.serv.OnOpen(ctx)

	for !ws.isClose {
		_, b, e := ws.conn.ReadMessage()
		if e != nil {
			log.Println("[websocket] read message err. ", e)
			_ = ws.Close()
			break
		}

		ws.serv.OnMessage(ctx, b)
	}
}

func (ws *Websocket) Close() error {
	if !ws.isClose {
		if err := ws.conn.Close(); err != nil {
			return err
		}
	}
	ws.isClose = true
	return nil
}
