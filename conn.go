package ssdb

import (
	"github.com/sunvim/spool"
	// "log"
	"net"
)

type Conn struct {
	Host string
	Port string
	Min  int
	Max  int
}

func (this *Conn) NewPool(conf Conn) (Pool, error) {
	this = &conf
	if this.Host == "" {
		this.Host = "127.0.0.1"
	}
	if this.Port == "" {
		this.Port = "7530"
	}
	if this.Max == 0 {
		this.Max = 20
	}
	if this.Min == 0 {
		this.Min = 5
	}
	addr := this.Host + ":" + this.Port
	fac := func() (net.Conn, error) { return net.Dial("tcp", addr) }
	return spool.NewChannelPool(this.Min, this.Max, fac)
}
