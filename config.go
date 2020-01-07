package melody

import "time"

// Config melody configuration struct.
type Config struct {
	WriteWait         time.Duration // Milliseconds until write times out.
	PongWait          time.Duration // Timeout for waiting on pong.
	PingPeriod        time.Duration // Milliseconds between pings.
	MaxMessageSize    int64         // Maximum size in bytes of a message.
	MessageBufferSize int           // The max amount of messages that can be in a sessions buffer before it starts dropping them.
}

func newConfig() *Config {
	return &Config{
		WriteWait:         10 * time.Second,
		PongWait:          (14 * 24 * time.Hour), // Because some browsers no ping pong, so we don't do this.
		PingPeriod:        (14 * 24 * time.Hour), // Because some browsers no ping pong, so we don't do this.
		MaxMessageSize:    2048,                  // 512 is something tiny.
		MessageBufferSize: 256,
	}
}
