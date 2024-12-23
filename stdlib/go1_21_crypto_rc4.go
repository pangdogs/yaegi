// Code generated by 'yaegi extract crypto/rc4'. DO NOT EDIT.

//go:build go1.21 && !go1.22 && !go1.23
// +build go1.21,!go1.22,!go1.23

package stdlib

import (
	"crypto/rc4"
	"reflect"
)

func init() {
	Symbols["crypto/rc4/rc4"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"NewCipher": reflect.ValueOf(rc4.NewCipher),

		// type definitions
		"Cipher":       reflect.ValueOf((*rc4.Cipher)(nil)),
		"KeySizeError": reflect.ValueOf((*rc4.KeySizeError)(nil)),
	}
}
