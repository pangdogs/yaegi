// Code generated by 'yaegi extract testing/slogtest'. DO NOT EDIT.

//go:build go1.22 && !go1.23
// +build go1.22,!go1.23

package stdlib

import (
	"reflect"
	"testing/slogtest"
)

func init() {
	Symbols["testing/slogtest/slogtest"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Run":         reflect.ValueOf(slogtest.Run),
		"TestHandler": reflect.ValueOf(slogtest.TestHandler),
	}
}
