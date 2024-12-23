// Code generated by 'yaegi extract testing/iotest'. DO NOT EDIT.

//go:build go1.23
// +build go1.23

package stdlib

import (
	"reflect"
	"testing/iotest"
)

func init() {
	Symbols["testing/iotest/iotest"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"DataErrReader":  reflect.ValueOf(iotest.DataErrReader),
		"ErrReader":      reflect.ValueOf(iotest.ErrReader),
		"ErrTimeout":     reflect.ValueOf(&iotest.ErrTimeout).Elem(),
		"HalfReader":     reflect.ValueOf(iotest.HalfReader),
		"NewReadLogger":  reflect.ValueOf(iotest.NewReadLogger),
		"NewWriteLogger": reflect.ValueOf(iotest.NewWriteLogger),
		"OneByteReader":  reflect.ValueOf(iotest.OneByteReader),
		"TestReader":     reflect.ValueOf(iotest.TestReader),
		"TimeoutReader":  reflect.ValueOf(iotest.TimeoutReader),
		"TruncateWriter": reflect.ValueOf(iotest.TruncateWriter),
	}
}
