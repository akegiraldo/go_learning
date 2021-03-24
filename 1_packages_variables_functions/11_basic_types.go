package main

import "fmt"

var (

	_bool       bool       = false
	_string     string     = "text"
	_int        int        = 1
	_int8       int8       = 1<<7 - 1
	_int16      int16      = 1<<15 - 1
	_int32      int32      = 1<<31 - 1
	_int64      int64      = 1<<63 - 1
	_uint       uint       = 1
	_uint8      uint8      = 1<<8 - 1
	_uint16     uint16     = 1<<16 - 1
	_uint32     uint32     = 1<<32 - 1
	_uint64     uint64     = 1<<64 - 1
	_byte       byte       = 1<<8 - 1
	_rune       rune       = 1<<31 - 1
	_float32    float32    = 1<<32 - 1.1
	_float64    float64    = 1<<64 - 1.1
	_complex64  complex64  = 1<<64 - 1i
	_complex128 complex128 = 1<<128 - 1i
)

func main() {

	fmt.Printf("Type: %T Value: %v\n", _bool, _bool)
	fmt.Printf("Type: %T Value: %v\n", _string, _string)
	fmt.Printf("Type: %T Value: %v\n", _int, _int)
	fmt.Printf("Type: %T Value: %v\n", _int8, _int8)
	fmt.Printf("Type: %T Value: %v\n", _int16, _int16)
	fmt.Printf("Type: %T Value: %v\n", _int32, _int32)
	fmt.Printf("Type: %T Value: %v\n", _int64, _int64)
	fmt.Printf("Type: %T Value: %v\n", _uint, _uint)
	fmt.Printf("Type: %T Value: %v\n", _uint8, _uint8)
	fmt.Printf("Type: %T Value: %v\n", _uint16, _uint16)
	fmt.Printf("Type: %T Value: %v\n", _uint32, _uint32)
	fmt.Printf("Type: %T Value: %v\n", _uint64, _uint64)
	fmt.Printf("Type: %T Value: %v\n", _byte, _byte)
	fmt.Printf("Type: %T Value: %v\n", _rune, _rune)
	fmt.Printf("Type: %T Value: %v\n", _float32, _float32)
	fmt.Printf("Type: %T Value: %v\n", _float64, _float64)
	fmt.Printf("Type: %T Value: %v\n", _complex64, _complex64)
	fmt.Printf("Type: %T Value: %v\n", _complex128, _complex128)
}