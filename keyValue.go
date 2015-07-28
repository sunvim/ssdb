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

func (c *Client) Set(key string, val string) (interface{}, error) {
	resp, err := c.Do("set", key, val)
	if err != nil {
		return nil, err
	}
	if len(resp) == 2 && resp[0] == "ok" {
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
	if len(resp) == 2 && resp[0] == "ok" {
		return resp[1], nil
	}
	if resp[0] == "not_found" {
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
	if len(resp) > 0 && resp[0] == "ok" {
		return true, nil
	}
	return nil, fmt.Errorf("bad response:resp:%v:", resp)
}
