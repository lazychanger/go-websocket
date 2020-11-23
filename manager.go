package websocket

import (
	"errors"
	"net/http"
)

type Manager struct {
	m  map[string]*Websocket
	hs []SetConfig
}

var (
	errorNotFound = errors.New("not found this connection")
)

func (m *Manager) SetConfig(hs ...SetConfig) {
	m.hs = append(m.hs, hs...)
	return
}

func (m *Manager) NewWebsocket(w http.ResponseWriter, r *http.Request, ws Server, hs ...SetConfig) {
	wsocket := NewWebsocket(w, r, ws, m.mergeHs(hs...)...)
	m.m[wsocket.Id] = wsocket
	return
}

func (m *Manager) Send(id string, data []byte) error {
	if ws, ok := m.m[id]; !ok {
		return errorNotFound
	} else {
		return ws.conn.WriteMessage(0, data)
	}
}

func (m *Manager) SendJson(id string, data interface{}) error {
	if ws, ok := m.m[id]; !ok {
		return errorNotFound
	} else {
		return ws.conn.WriteJSON(data)
	}
}

func (m *Manager) mergeHs(hs ...SetConfig) []SetConfig {
	return append(m.hs, hs...)
}
