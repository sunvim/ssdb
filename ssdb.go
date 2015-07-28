/*
	author: sunsc
	time: 2015-07-29
	positon:ShangHai
	License: MIT
*/
package ssdb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
)

var pool Pool

type Client struct {
	sock     net.Conn
	recv_buf bytes.Buffer
}

func NewPool(conn Conn) error {
	var err error
	pool, err = conn.NewPool(conn)
	// fmt.Println("init pool num=>", pool.Len())
	return err
}

func ClosePool() {
	pool.Close()
}

func (this *Client) GetConn() *Client {
	if this.sock != nil {
		// fmt.Println("had a connect,don't again!")
		return this
	}
	var err error
	this.sock, err = pool.Get()
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println("first get connect!")
	// fmt.Println("current pool num=>", pool.Len())
	return this
}

func (this *Client) Do(args ...interface{}) ([]string, error) {
	err := this.send(args)
	if err != nil {
		return nil, err
	}
	resp, err := this.recv()
	return resp, err
}

func (this *Client) Send(args ...interface{}) error {
	return this.send(args)
}

func (this *Client) send(args []interface{}) error {
	var buf bytes.Buffer
	for _, arg := range args {
		var s string
		switch arg := arg.(type) {
		case string:
			s = arg
		case []byte:
			s = string(arg)
		case []string:
			for _, s := range arg {
				buf.WriteString(fmt.Sprintf("%d", len(s)))
				buf.WriteByte('\n')
				buf.WriteString(s)
				buf.WriteByte('\n')
			}
			continue
		case int:
			s = fmt.Sprintf("%d", arg)
		case int64:
			s = fmt.Sprintf("%d", arg)
		case float64:
			s = fmt.Sprintf("%f", arg)
		case bool:
			if arg {
				s = "1"
			} else {
				s = "0"
			}
		case nil:
			s = ""
		default:
			return fmt.Errorf("bad arguments")
		}
		buf.WriteString(fmt.Sprintf("%d", len(s)))
		buf.WriteByte('\n')
		buf.WriteString(s)
		buf.WriteByte('\n')
	}
	buf.WriteByte('\n')
	_, err := this.sock.Write(buf.Bytes())
	return err
}

func (this *Client) Recv() ([]string, error) {
	return this.recv()
}

func (this *Client) recv() ([]string, error) {
	var tmp [8192]byte
	for {
		resp := this.parse()
		if resp == nil || len(resp) > 0 {
			return resp, nil
		}
		n, err := this.sock.Read(tmp[0:])
		if err != nil {
			return nil, err
		}
		this.recv_buf.Write(tmp[0:n])
	}
}

func (this *Client) parse() []string {
	resp := []string{}
	buf := this.recv_buf.Bytes()
	var idx, offset int
	idx = 0
	offset = 0

	for {
		idx = bytes.IndexByte(buf[offset:], '\n')
		if idx == -1 {
			break
		}
		p := buf[offset : offset+idx]
		offset += idx + 1
		//fmt.Printf("> [%s]\n", p);
		if len(p) == 0 || (len(p) == 1 && p[0] == '\r') {
			if len(resp) == 0 {
				continue
			} else {
				this.recv_buf.Next(offset)
				return resp
			}
		}

		size, err := strconv.Atoi(string(p))
		if err != nil || size < 0 {
			return nil
		}
		if offset+size >= this.recv_buf.Len() {
			break
		}

		v := buf[offset : offset+size]
		resp = append(resp, string(v))
		offset += size + 1
	}

	//fmt.Printf("buf.size: %d packet not ready...\n", len(buf))
	return []string{}
}

// Close The Client Connection
func (this *Client) Close() error {
	err := this.sock.Close()
	this.sock = nil
	return err
}

//encoding value for ssdb
func (this *Client) encoding(value interface{}) string {
	switch t := value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, complex64, complex128:
		return String(t)
	case string: //byte==uint8
		return t
	case []byte:
		return string(t)
	case bool:
		if t {
			return "1"
		} else {
			return "0"
		}
	case nil:
		return ""
	default:
		if bs, err := json.Marshal(value); err == nil {
			return string(bs)
		} else {
			return "encoding value failed"
		}
	}
}
