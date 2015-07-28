/*
	author: sunsc
	time: 2015-07-29
	positon:ShangHai
	License: MIT
*/
package ssdb

import (
	"encoding/json"
	"time"
)

type Value string

func (this Value) String() string {
	return string(this)
}

func (this Value) Int64() int64 {
	return Int64(this)
}

func (this Value) Int32() int32 {
	return int32(this.Int64())
}

func (this Value) Int16() int16 {
	return int16(this.Int64())
}

func (this Value) Int8() int8 {
	return int8(this.Int64())
}

func (this Value) Int() int {
	return int(this.Int64())
}

func (this Value) UInt64() uint64 {
	return Uint64(this)
}

func (this Value) UInt32() uint32 {
	return uint32(this.UInt64())
}

func (this Value) UInt16() uint16 {
	return uint16(this.UInt64())
}

func (this Value) UInt8() uint8 {
	return uint8(this.UInt64())
}

func (this Value) Byte() byte {
	return this.UInt8()
}

func (this Value) UInt() uint {
	return uint(this.UInt64())
}

func (this Value) Float64() float64 {
	return Float64(this)
}

func (this Value) Float32() float32 {
	return float32(this.Float64())
}

func (this Value) Bool() bool {
	return Bool(this)
}

func (this Value) Time() time.Time {
	return Time(this)
}

func (this Value) Duration() time.Duration {
	return Duration(this)
}

func (this Value) Bytes() []byte {
	return []byte(this)
}

func (this Value) IsEmpty() bool {
	return this == ""
}

//this function for turn the right value
//json encoding format
func (this Value) As(value interface{}) (err error) {
	err = json.Unmarshal(this.Bytes(), value)
	return
}
