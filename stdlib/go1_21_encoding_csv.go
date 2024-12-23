// Code generated by 'yaegi extract encoding/csv'. DO NOT EDIT.

//go:build go1.21 && !go1.22 && !go1.23
// +build go1.21,!go1.22,!go1.23

package stdlib

import (
	"encoding/csv"
	"reflect"
)

func init() {
	Symbols["encoding/csv/csv"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"ErrBareQuote":     reflect.ValueOf(&csv.ErrBareQuote).Elem(),
		"ErrFieldCount":    reflect.ValueOf(&csv.ErrFieldCount).Elem(),
		"ErrQuote":         reflect.ValueOf(&csv.ErrQuote).Elem(),
		"ErrTrailingComma": reflect.ValueOf(&csv.ErrTrailingComma).Elem(),
		"NewReader":        reflect.ValueOf(csv.NewReader),
		"NewWriter":        reflect.ValueOf(csv.NewWriter),

		// type definitions
		"ParseError": reflect.ValueOf((*csv.ParseError)(nil)),
		"Reader":     reflect.ValueOf((*csv.Reader)(nil)),
		"Writer":     reflect.ValueOf((*csv.Writer)(nil)),
	}
}
