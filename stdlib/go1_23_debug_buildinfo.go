// Code generated by 'yaegi extract debug/buildinfo'. DO NOT EDIT.

//go:build go1.23
// +build go1.23

package stdlib

import (
	"debug/buildinfo"
	"reflect"
)

func init() {
	Symbols["debug/buildinfo/buildinfo"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Read":     reflect.ValueOf(buildinfo.Read),
		"ReadFile": reflect.ValueOf(buildinfo.ReadFile),

		// type definitions
		"BuildInfo": reflect.ValueOf((*buildinfo.BuildInfo)(nil)),
	}
}
