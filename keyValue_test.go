/*
	author: sunsc
	time: 2015-07-29
	positon:ShangHai
	License: MIT
*/
package ssdb

import (
	"testing"
	// "time"
)

func Test_Set(t *testing.T) {
	err := NewPool(Conn{SSDB_SERVER, SSDB_PORT, 2, 4})
	if err != nil {
		t.Error("create pool failed=>", err)
	}
	defer ClosePool()
	c := new(Client)
	ok, err := c.GetConn().Set("hello", "sunsc")
	if ok.(bool) && err == nil {
		t.Log("ok")
	} else {
		t.Fatal("set hello sunsc failed")
	}
}

func Test_Get(t *testing.T) {
	err := NewPool(Conn{SSDB_SERVER, SSDB_PORT, 2, 4})
	if err != nil {
		t.Error("create pool failed=>", err)
	}
	defer ClosePool()
	c := new(Client)
	val, err := c.GetConn().Get("hello")
	if val.(string) == "sunsc" && err == nil {
		t.Log("ok")
	} else {
		t.Fatal("key hello value is=>", val.(string))
	}
}

func Test_Del(t *testing.T) {
	err := NewPool(Conn{SSDB_SERVER, SSDB_PORT, 2, 4})
	if err != nil {
		t.Error("create pool failed=>", err)
	}
	defer ClosePool()
	c := new(Client)
	ok, err := c.GetConn().Del("hello")
	if ok.(bool) && err == nil {
		t.Log("ok")
	} else {
		t.Fatal("del key hello failed: ", err)
	}
}

func Test_Setx(t *testing.T) {
	err := NewPool(Conn{SSDB_SERVER, SSDB_PORT, 2, 4})
	if err != nil {
		t.Error("create pool failed=>", err)
	}
	defer ClosePool()
	c := new(Client)
	ok, err := c.GetConn().Setx("hello2", "sunsc", 10)
	if ok.(bool) && err == nil {
		t.Log("ok")
	} else {
		t.Fatal("setx hello2 sunsc failed")
	}
}

func Test_Expire(t *testing.T) {
	err := NewPool(Conn{SSDB_SERVER, SSDB_PORT, 2, 4})
	if err != nil {
		t.Error("create pool failed=>", err)
	}
	defer ClosePool()
	c := new(Client)
	c.GetConn().Set("hello3", "sunsc")
	ok, err := c.Expire("hello3", 10)
	if ok.(bool) && err == nil {
		t.Log("ok")
	} else {
		t.Fatal("expire hello3 10 failed")
	}
}

func Test_Ttl(t *testing.T) {
	err := NewPool(Conn{SSDB_SERVER, SSDB_PORT, 2, 4})
	if err != nil {
		t.Error("create pool failed=>", err)
	}
	defer ClosePool()
	c := new(Client)
	c.GetConn().Setx("hello4", "sunsc", 30)
	val, err := c.Ttl("hello4")
	if err == nil {
		t.Log("key hello4 ttl is :", val)
	} else {
		t.Fatal("ttl hello4  failed")
	}
}

func Test_Getset(t *testing.T) {
	err := NewPool(Conn{SSDB_SERVER, SSDB_PORT, 2, 4})
	if err != nil {
		t.Error("create pool failed=>", err)
	}
	defer ClosePool()
	c := new(Client)
	c.GetConn().Set("hello5", "sunsc")
	val, err := c.Getset("hello5", "world")
	if err == nil && val.(string) == "sunsc" {
		t.Log("getset hello5 return value is :", val.(string))
	} else {
		t.Fatal("getset hello5  failed")
	}
}

func Test_Incr(t *testing.T) {
	err := NewPool(Conn{SSDB_SERVER, SSDB_PORT, 2, 4})
	if err != nil {
		t.Error("create pool failed=>", err)
	}
	defer ClosePool()
	c := new(Client)
	c.GetConn().Del("hello6")
	val, err := c.Incr("hello6")
	if err == nil && val.(int) == 1 {
		t.Log("incr hello6 return value is :", val)
	} else {
		t.Fatal("incr hello6  failed")
	}
}

func Test_Exists(t *testing.T) {
	err := NewPool(Conn{SSDB_SERVER, SSDB_PORT, 2, 4})
	if err != nil {
		t.Error("create pool failed=>", err)
	}
	defer ClosePool()
	c := new(Client)
	c.GetConn().Del("hello7")
	val, err := c.Exists("hello7")
	if err == nil && !val.(bool) {
		t.Log("exists hello7 return value is :", val)
	} else {
		t.Fatal("exists hello7  failed")
	}
}

func Test_Getbit(t *testing.T) {
	err := NewPool(Conn{SSDB_SERVER, SSDB_PORT, 2, 4})
	if err != nil {
		t.Error("create pool failed=>", err)
	}
	defer ClosePool()
	c := new(Client)
	c.GetConn().Set("h8", 5)
	val, err := c.Getbit("h8", 2)
	if err == nil && val.(string) == "1" {
		t.Log("getbit h8 2 return value is :", val)
	} else {
		t.Fatal("getbit h8  failed")
	}
}

func Test_Setbit(t *testing.T) {
	err := NewPool(Conn{SSDB_SERVER, SSDB_PORT, 2, 4})
	if err != nil {
		t.Error("create pool failed=>", err)
	}
	defer ClosePool()
	c := new(Client)
	c.GetConn().Set("h9", 8)
	val, err := c.Setbit("h9", 2, 1)
	if err == nil && val.(string) == "0" {
		t.Log("setbit h9 2 return value is :", val)
	} else {
		t.Fatal("setbit h9  failed")
	}
}
