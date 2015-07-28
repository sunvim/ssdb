/*
	author: sunsc
	time: 2015-07-29
	positon:ShangHai
	License: MIT
*/
package ssdb

func (this *Client) Hset(setName, key string, value interface{}) (err error) {
	resp, err := this.Do("hset", setName, key, this.encoding(value))
	if err != nil {
		return err
	}

	if len(resp) > 0 && resp[0] == RET_OK {
		return nil
	}
	return makeError(resp, setName, key)
}

func (this *Client) Hget(setName, key string) (value Value, err error) {
	resp, err := this.Do("hget", setName, key)
	if err != nil {
		return "", NewError(err, "Hget %s/%s error", setName, key)
	}
	if len(resp) == 2 && resp[0] == "ok" {
		return Value(resp[1]), nil
	}
	return "", makeError(resp, setName, key)
}

func (this *Client) Hdel(setName, key string) (err error) {
	resp, err := this.Do("hdel", setName, key)
	if err != nil {
		return NewError(err, "Hdel %s/%s error", setName, key)
	}
	if len(resp) > 0 && resp[0] == "ok" {
		return nil
	}
	return makeError(resp, setName, key)
}

func (this *Client) Hexists(setName, key string) (re bool, err error) {
	resp, err := this.Do("hexists", setName, key)
	if err != nil {
		return false, NewError(err, "Hexists %s/%s error", setName, key)
	}

	if len(resp) == 2 && resp[0] == "ok" {
		return resp[1] == "1", nil
	}
	return false, makeError(resp, setName, key)
}

func (this *Client) Hclear(setName string) (err error) {
	resp, err := this.Do("hclear", setName)
	if err != nil {
		return NewError(err, "Hclear %s error", setName)
	}

	if len(resp) > 0 && resp[0] == "ok" {
		return nil
	}
	return makeError(resp, setName)
}

func (this *Client) Hkeys(setName, key_start, key_end string, limit int) ([]Value, error) {
	var val []Value
	resp, err := this.Do("hkeys", setName, key_start, key_end, limit)
	if err != nil {
		return nil, NewError(err, "Hkeys %s error", setName)
	}
	for k, v := range resp {
		if k == 0 {
			continue
		} else {
			val = append(val, Value(v))
		}
	}
	return val, nil
}

func (this *Client) Hsize(setName string) (int, error) {
	var ret int
	resp, err := this.Do("hsize", setName)
	if err != nil {
		return 0, NewError(err, "Hsize %s error", setName)
	}
	if len(resp) > 0 && resp[0] == "ok" {
		ret = Value(resp[1]).Int()
		return ret, nil
	}
	return ret, makeError(resp, setName)
}
