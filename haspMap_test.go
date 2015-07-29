/*
	author: sunsc
	time: 2015-07-29
	positon:ShangHai
	License: MIT
*/
package ssdb

import (
	"testing"
)

func Test_Hincr(t *testing.T) {
	err := NewPool(Conn{SSDB_SERVER, SSDB_PORT, 2, 4})
	if err != nil {
		t.Error("create pool failed=>", err)
	}
	defer ClosePool()
	c := new(Client)
	c.GetConn().Hdel("hktest", "hkey1")
	val, err := c.Hincr("hktest", "hkey1")
	if err == nil && val.(int) == 1 {
		t.Log("hincr hktest hkey1 value is:", val)
	} else {
		t.Fatal("hincr hktest hkey1 value failed:", err)
	}
}
