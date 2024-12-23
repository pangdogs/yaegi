// Code generated by 'yaegi extract compress/gzip'. DO NOT EDIT.

//go:build go1.23
// +build go1.23

package stdlib

import (
	"compress/gzip"
	"go/constant"
	"go/token"
	"reflect"
)

func init() {
	Symbols["compress/gzip/gzip"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"BestCompression":    reflect.ValueOf(constant.MakeFromLiteral("9", token.INT, 0)),
		"BestSpeed":          reflect.ValueOf(constant.MakeFromLiteral("1", token.INT, 0)),
		"DefaultCompression": reflect.ValueOf(constant.MakeFromLiteral("-1", token.INT, 0)),
		"ErrChecksum":        reflect.ValueOf(&gzip.ErrChecksum).Elem(),
		"ErrHeader":          reflect.ValueOf(&gzip.ErrHeader).Elem(),
		"HuffmanOnly":        reflect.ValueOf(constant.MakeFromLiteral("-2", token.INT, 0)),
		"NewReader":          reflect.ValueOf(gzip.NewReader),
		"NewWriter":          reflect.ValueOf(gzip.NewWriter),
		"NewWriterLevel":     reflect.ValueOf(gzip.NewWriterLevel),
		"NoCompression":      reflect.ValueOf(constant.MakeFromLiteral("0", token.INT, 0)),

		// type definitions
		"Header": reflect.ValueOf((*gzip.Header)(nil)),
		"Reader": reflect.ValueOf((*gzip.Reader)(nil)),
		"Writer": reflect.ValueOf((*gzip.Writer)(nil)),
	}
}
