// Code generated by 'yaegi extract encoding/asn1'. DO NOT EDIT.

//go:build go1.22 && !go1.23
// +build go1.22,!go1.23

package stdlib

import (
	"encoding/asn1"
	"go/constant"
	"go/token"
	"reflect"
)

func init() {
	Symbols["encoding/asn1/asn1"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"ClassApplication":     reflect.ValueOf(constant.MakeFromLiteral("1", token.INT, 0)),
		"ClassContextSpecific": reflect.ValueOf(constant.MakeFromLiteral("2", token.INT, 0)),
		"ClassPrivate":         reflect.ValueOf(constant.MakeFromLiteral("3", token.INT, 0)),
		"ClassUniversal":       reflect.ValueOf(constant.MakeFromLiteral("0", token.INT, 0)),
		"Marshal":              reflect.ValueOf(asn1.Marshal),
		"MarshalWithParams":    reflect.ValueOf(asn1.MarshalWithParams),
		"NullBytes":            reflect.ValueOf(&asn1.NullBytes).Elem(),
		"NullRawValue":         reflect.ValueOf(&asn1.NullRawValue).Elem(),
		"TagBMPString":         reflect.ValueOf(constant.MakeFromLiteral("30", token.INT, 0)),
		"TagBitString":         reflect.ValueOf(constant.MakeFromLiteral("3", token.INT, 0)),
		"TagBoolean":           reflect.ValueOf(constant.MakeFromLiteral("1", token.INT, 0)),
		"TagEnum":              reflect.ValueOf(constant.MakeFromLiteral("10", token.INT, 0)),
		"TagGeneralString":     reflect.ValueOf(constant.MakeFromLiteral("27", token.INT, 0)),
		"TagGeneralizedTime":   reflect.ValueOf(constant.MakeFromLiteral("24", token.INT, 0)),
		"TagIA5String":         reflect.ValueOf(constant.MakeFromLiteral("22", token.INT, 0)),
		"TagInteger":           reflect.ValueOf(constant.MakeFromLiteral("2", token.INT, 0)),
		"TagNull":              reflect.ValueOf(constant.MakeFromLiteral("5", token.INT, 0)),
		"TagNumericString":     reflect.ValueOf(constant.MakeFromLiteral("18", token.INT, 0)),
		"TagOID":               reflect.ValueOf(constant.MakeFromLiteral("6", token.INT, 0)),
		"TagOctetString":       reflect.ValueOf(constant.MakeFromLiteral("4", token.INT, 0)),
		"TagPrintableString":   reflect.ValueOf(constant.MakeFromLiteral("19", token.INT, 0)),
		"TagSequence":          reflect.ValueOf(constant.MakeFromLiteral("16", token.INT, 0)),
		"TagSet":               reflect.ValueOf(constant.MakeFromLiteral("17", token.INT, 0)),
		"TagT61String":         reflect.ValueOf(constant.MakeFromLiteral("20", token.INT, 0)),
		"TagUTCTime":           reflect.ValueOf(constant.MakeFromLiteral("23", token.INT, 0)),
		"TagUTF8String":        reflect.ValueOf(constant.MakeFromLiteral("12", token.INT, 0)),
		"Unmarshal":            reflect.ValueOf(asn1.Unmarshal),
		"UnmarshalWithParams":  reflect.ValueOf(asn1.UnmarshalWithParams),

		// type definitions
		"BitString":        reflect.ValueOf((*asn1.BitString)(nil)),
		"Enumerated":       reflect.ValueOf((*asn1.Enumerated)(nil)),
		"Flag":             reflect.ValueOf((*asn1.Flag)(nil)),
		"ObjectIdentifier": reflect.ValueOf((*asn1.ObjectIdentifier)(nil)),
		"RawContent":       reflect.ValueOf((*asn1.RawContent)(nil)),
		"RawValue":         reflect.ValueOf((*asn1.RawValue)(nil)),
		"StructuralError":  reflect.ValueOf((*asn1.StructuralError)(nil)),
		"SyntaxError":      reflect.ValueOf((*asn1.SyntaxError)(nil)),
	}
}
