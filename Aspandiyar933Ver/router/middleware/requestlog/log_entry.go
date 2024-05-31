package requestlog

import (
	"net"
	"time"
)

type logEntry struct {
	RequestID         string
	ReceivedTime      time.Time
	RequestMethod     string
	RequestURL        string
	RequestHeaderSize int64
	RequestBodySize   int64
	UserAgent         string
	Referer           string
	Proto             string

	RemoteIP string
	ServerIP string

	Status             int
	ResponceHeaderSize int64
	ResponceBodySize   int64
	Latency            time.Duration
}

func ipFromHostPort(hp string) string {
	h, _, err := net.SplitHostPort(hp)
	if err != nil {
		return ""
	}
	if len(h) > 0 && h[0] == '[' {
		return h[1 : len(h)-1]
	}
	return h 
}





