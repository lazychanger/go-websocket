package websocket

type Message interface {
	Bytes() []byte
}

type JsonMessage struct {
}

func (j *JsonMessage) Bytes() []byte {
	panic("implement me")
}
