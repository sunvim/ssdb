/*
	author: sunsc
	time: 2015-07-29
	positon:ShangHai
	License: MIT
*/
package ssdb

import (
	"fmt"
)

func (c *Client) Set(key string, val interface{}) (interface{}, error) {
	resp, err := c.Do("set", key, val)
	if err != nil {
		return nil, err
	}
	if len(resp) == 2 && resp[0] == RET_OK {
		return true, nil
	}
	return nil, fmt.Errorf("bad response")
}

// TODO: Will somebody write addition semantic methods?
func (c *Client) Get(key string) (interface{}, error) {
	resp, err := c.Do("get", key)
	if err != nil {
		return nil, err
	}
	if len(resp) == 2 && resp[0] == RET_OK {
		return resp[1], nil
	}
	if resp[0] == NOT_FOUND {
		return nil, nil
	}
	return nil, fmt.Errorf("bad response")
}

func (c *Client) Del(key string) (interface{}, error) {
	resp, err := c.Do("del", key)
	if err != nil {
		return nil, err
	}
	//response looks like this: [ok 1]
	if len(resp) > 0 && resp[0] == RET_OK {
		return true, nil
	}
	return nil, fmt.Errorf("bad response:resp:%v:", resp)
}

func (c *Client) Setx(key string, val interface{}, ttl int) (interface{}, error) {
	resp, err := c.Do("setx", key, val, ttl)
	if err != nil {
		return nil, err
	}
	if len(resp) == 2 && resp[0] == RET_OK {
		return true, nil
	}
	return nil, fmt.Errorf("bad response")
}

func (c *Client) Expire(key string, ttl int) (interface{}, error) {
	resp, err := c.Do("expire", key, ttl)
	if err != nil {
		return nil, err
	}
	if len(resp) == 2 && resp[0] == RET_OK {
		return true, nil
	}
	return nil, fmt.Errorf("bad response")
}

func (c *Client) Ttl(key string) (interface{}, error) {
	resp, err := c.Do("ttl", key)
	if err != nil {
		return nil, err
	}
	if len(resp) == 2 && resp[0] == RET_OK {
		return resp[1], nil
	}
	if resp[0] == NOT_FOUND {
		return nil, nil
	}
	return nil, fmt.Errorf("bad response")
}

func (c *Client) Getset(key string, val interface{}) (interface{}, error) {
	resp, err := c.Do("getset", key, val)
	if err != nil {
		return nil, err
	}
	if len(resp) == 2 && resp[0] == RET_OK {
		return resp[1], nil
	}
	if resp[0] == NOT_FOUND {
		return nil, nil
	}
	return nil, fmt.Errorf("bad response")
}

func (c *Client) Incr(key string, val ...int) (interface{}, error) {
	var num int
	if val == nil {
		num = 1
	} else {
		num = val[0]
	}
	resp, err := c.Do("incr", key, num)
	if err != nil {
		return nil, err
	}
	if len(resp) == 2 && resp[0] == RET_OK {
		return Value(resp[1]).Int(), nil
	}
	return nil, fmt.Errorf("bad response")
}

func (c *Client) Exists(key string) (interface{}, error) {
	resp, err := c.Do("exists", key)
	if err != nil {
		return nil, err
	}
	if len(resp) == 2 && resp[0] == RET_OK {
		return Value(resp[1]).Bool(), nil
	}
	return nil, fmt.Errorf("bad response")
}

func (c *Client) Getbit(key string, offset int) (interface{}, error) {
	resp, err := c.Do("getbit", key, offset)
	if err != nil {
		return nil, err
	}
	if len(resp) == 2 && resp[0] == RET_OK {
		return resp[1], nil
	}
	return nil, fmt.Errorf("bad response")
}

func (c *Client) Setbit(key string, offset int, val interface{}) (interface{}, error) {
	resp, err := c.Do("setbit", key, offset, val)
	if err != nil {
		return nil, err
	}
	if len(resp) == 2 && resp[0] == RET_OK {
		return resp[1], nil
	}
	return nil, fmt.Errorf("bad response")
}
