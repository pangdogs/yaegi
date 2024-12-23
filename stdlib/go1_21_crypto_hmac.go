// Code generated by 'yaegi extract crypto/hmac'. DO NOT EDIT.

//go:build go1.21 && !go1.22 && !go1.23
// +build go1.21,!go1.22,!go1.23

package stdlib

import (
	"crypto/hmac"
	"reflect"
)

func init() {
	Symbols["crypto/hmac/hmac"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Equal": reflect.ValueOf(hmac.Equal),
		"New":   reflect.ValueOf(hmac.New),
	}
}
