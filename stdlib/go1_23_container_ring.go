// Code generated by 'yaegi extract container/ring'. DO NOT EDIT.

//go:build go1.23
// +build go1.23

package stdlib

import (
	"container/ring"
	"reflect"
)

func init() {
	Symbols["container/ring/ring"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"New": reflect.ValueOf(ring.New),

		// type definitions
		"Ring": reflect.ValueOf((*ring.Ring)(nil)),
	}
}
