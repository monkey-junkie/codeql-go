// Code generated by depstubber. DO NOT EDIT.
// This is a simple stub for k8s.io/apimachinery/pkg/conversion, strictly for use in testing.

// See the LICENSE file for information about the licensing of the original library.
// Source: k8s.io/apimachinery/pkg/conversion (exports: Scope; functions: )

// Package conversion is a stub of k8s.io/apimachinery/pkg/conversion, generated by depstubber.
package conversion

import (
	reflect "reflect"
)

type FieldMappingFunc func(string, reflect.StructTag, reflect.StructTag) (string, string)

type FieldMatchingFlags int

func (_ FieldMatchingFlags) IsSet(_ FieldMatchingFlags) bool {
	return false
}

type Meta struct {
	KeyNameMapping FieldMappingFunc
	Context        interface{}
}

type Scope interface {
	Convert(_ interface{}, _ interface{}, _ FieldMatchingFlags) error
	DestTag() reflect.StructTag
	Flags() FieldMatchingFlags
	Meta() *Meta
	SrcTag() reflect.StructTag
}
