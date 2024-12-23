// Code generated by 'yaegi extract net/rpc/jsonrpc'. DO NOT EDIT.

//go:build go1.22 && !go1.23
// +build go1.22,!go1.23

package stdlib

import (
	"net/rpc/jsonrpc"
	"reflect"
)

func init() {
	Symbols["net/rpc/jsonrpc/jsonrpc"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Dial":           reflect.ValueOf(jsonrpc.Dial),
		"NewClient":      reflect.ValueOf(jsonrpc.NewClient),
		"NewClientCodec": reflect.ValueOf(jsonrpc.NewClientCodec),
		"NewServerCodec": reflect.ValueOf(jsonrpc.NewServerCodec),
		"ServeConn":      reflect.ValueOf(jsonrpc.ServeConn),
	}
}
