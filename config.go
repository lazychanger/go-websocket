package websocket

type Config struct {
	msg Message
}

type SetConfig func(config *Config) *Config

func SetMessage(message Message) SetConfig {
	return func(config *Config) *Config {
		config.msg = message
		return config
	}
}
