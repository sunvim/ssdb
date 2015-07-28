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

var (
	FormatString = "%v\nthe trace error is\n%s"
)

//return new error info with old error info
func NewError(err error, format string, p ...interface{}) error {
	return fmt.Errorf(FormatString, fmt.Sprintf(format, p...), err)
}
