// Code generated by 'yaegi extract go/build/constraint'. DO NOT EDIT.

//go:build go1.22 && !go1.23
// +build go1.22,!go1.23

package stdlib

import (
	"go/build/constraint"
	"reflect"
)

func init() {
	Symbols["go/build/constraint/constraint"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"GoVersion":      reflect.ValueOf(constraint.GoVersion),
		"IsGoBuild":      reflect.ValueOf(constraint.IsGoBuild),
		"IsPlusBuild":    reflect.ValueOf(constraint.IsPlusBuild),
		"Parse":          reflect.ValueOf(constraint.Parse),
		"PlusBuildLines": reflect.ValueOf(constraint.PlusBuildLines),

		// type definitions
		"AndExpr":     reflect.ValueOf((*constraint.AndExpr)(nil)),
		"Expr":        reflect.ValueOf((*constraint.Expr)(nil)),
		"NotExpr":     reflect.ValueOf((*constraint.NotExpr)(nil)),
		"OrExpr":      reflect.ValueOf((*constraint.OrExpr)(nil)),
		"SyntaxError": reflect.ValueOf((*constraint.SyntaxError)(nil)),
		"TagExpr":     reflect.ValueOf((*constraint.TagExpr)(nil)),

		// interface wrapper definitions
		"_Expr": reflect.ValueOf((*_go_build_constraint_Expr)(nil)),
	}
}

// _go_build_constraint_Expr is an interface wrapper for Expr type
type _go_build_constraint_Expr struct {
	IValue  interface{}
	WEval   func(ok func(tag string) bool) bool
	WString func() string
}

func (W _go_build_constraint_Expr) Eval(ok func(tag string) bool) bool { return W.WEval(ok) }
func (W _go_build_constraint_Expr) String() string {
	if W.WString == nil {
		return ""
	}
	return W.WString()
}
