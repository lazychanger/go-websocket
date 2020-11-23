package websocket

import (
	"errors"
	"sync"
)

var (
	keyAlreadyExist error = errors.New("key already exist")
	keyNotExist     error = errors.New("key not exist")
)

type Context struct {
	state map[string]interface{}
	lock  *sync.Mutex
	ws    *Websocket
}

func (ctx *Context) Next() {
}

func (ctx *Context) Set(name string, val interface{}) error {
	ctx.lock.Lock()
	if _, ok := ctx.state[name]; ok {
		return keyAlreadyExist
	}

	ctx.state[name] = val
	ctx.lock.Unlock()
	return nil
}

func (ctx *Context) SetIfNotExist(name string, val interface{}) (ok bool) {
	ctx.lock.Lock()
	if _, ok := ctx.state[name]; ok {
		return ok
	}

	ctx.state[name] = val
	ctx.lock.Unlock()
	return false
}

func (ctx *Context) FocusSet(name string, val interface{}) {
	ctx.lock.Lock()
	ctx.state[name] = val
	ctx.lock.Unlock()
}

func (ctx *Context) Get(name string) (interface{}, error) {
	if val, ok := ctx.state[name]; ok {
		return val, nil
	}
	return nil, keyNotExist
}

func (ctx *Context) GetWithDefault(name string, d interface{}) interface{} {

	if val, ok := ctx.state[name]; ok {
		return val
	}

	return d
}

func (ctx *Context) GetId() string {
	return ctx.ws.Id
}

func (ctx *Context) SendJson(data interface{}) error {
	return ctx.ws.conn.WriteJSON(data)
}
