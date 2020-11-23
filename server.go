package websocket

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"net/http"
)

type Server interface {
	OnOpen(ctx *Context)
	OnMessage(ctx *Context, body []byte)
	OnClose(ctx *Context)
	OnConnect(ctx *Context, w http.ResponseWriter, r *http.Request) error
}

func NewWebsocket(w http.ResponseWriter, r *http.Request, ws Server, hs ...SetConfig) *Websocket {
	upgrade := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	ctx := &Context{}
	if err := ws.OnConnect(ctx, w, r); err != nil {
		return nil
	}

	id := uuid.New().String()
	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		return nil
	}

	return &Websocket{
		Id: id,
		Conf: &Config{
			msg: &JsonMessage{},
		},
		cfgSets: hs,
		conn:    conn,
		serv:    ws,
	}
}
