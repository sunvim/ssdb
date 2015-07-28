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

//check again  before return value
func makeError(resp []string, errKey ...interface{}) error {
	if len(resp) < 1 {
		return fmt.Errorf("ssdb respone error")
	}
	//if u want to catch this case , plz use "exist" command to check data
	if resp[0] == NOT_FOUND {
		return nil
	}
	if len(errKey) > 0 {
		return fmt.Errorf("access ssdb error, code is %v, parameter is %v", resp, errKey)
	} else {
		return fmt.Errorf("access ssdb error, code is %v", resp)
	}
}
