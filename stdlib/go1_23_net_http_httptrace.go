// Code generated by 'yaegi extract net/http/httptrace'. DO NOT EDIT.

//go:build go1.23
// +build go1.23

package stdlib

import (
	"net/http/httptrace"
	"reflect"
)

func init() {
	Symbols["net/http/httptrace/httptrace"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"ContextClientTrace": reflect.ValueOf(httptrace.ContextClientTrace),
		"WithClientTrace":    reflect.ValueOf(httptrace.WithClientTrace),

		// type definitions
		"ClientTrace":      reflect.ValueOf((*httptrace.ClientTrace)(nil)),
		"DNSDoneInfo":      reflect.ValueOf((*httptrace.DNSDoneInfo)(nil)),
		"DNSStartInfo":     reflect.ValueOf((*httptrace.DNSStartInfo)(nil)),
		"GotConnInfo":      reflect.ValueOf((*httptrace.GotConnInfo)(nil)),
		"WroteRequestInfo": reflect.ValueOf((*httptrace.WroteRequestInfo)(nil)),
	}
}
