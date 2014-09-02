package gomspec

import (
	"fmt"
	"reflect"
)

// Expect is a light assertion pattern for testing currently embedded
// GoMspec's code.  Future work may be to expose it as an interface
// to allow registering custom matchers (e.g. testify's assert package).
type Expect func(val interface{}) *expectation

type expectation struct {
	Value  interface{}
	Output formatter
}

// To tests a custom matching interface to the value.
func (e *expectation) To(desc string, value interface{}, matcher func(a, b interface{}) bool) {
	e.Output.PrintFeature()
	e.Output.PrintContext()
	e.Output.PrintWhen()
	if !matcher(e.Value, value) {
		e.Output.PrintTitleWithError()
		e.Output.PrintError(fmt.Sprintf("Expected `%v` to %s `%v`", e.Value, desc, value))
	} else {
		e.Output.PrintTitle()
	}
}

// ToEqual tests the equality of the expectation to the value of b.
func (e *expectation) ToEqual(b interface{}) {
	e.To(
		"equal",
		b,
		func(a, b interface{}) bool {
			return a == b
		},
	)
}

// ToNotEqual tests the inequality of the expectation to the value of b.
func (e *expectation) ToNotEqual(b interface{}) {
	e.To(
		"not equal",
		b,
		func(a, b interface{}) bool {
			return a != b
		},
	)
}

// ToNotBeNil tests that expectation is not equal to nil.
func (e *expectation) ToNotBeNil() {
	e.To(
		"exist",
		nil,
		func(a, b interface{}) bool {
			// TODO either inspect the TypeOf a for types that
			// support IsNil(); or, capture any panics that occur
			// and format the error nicely.
			return !reflect.ValueOf(a).IsNil()
		},
	)
}

// ToBeNil tests that the expectation is equal to nil.
func (e *expectation) ToBeNil() {
	e.To(
		"not exist",
		nil,
		func(a, b interface{}) bool {
			// TODO either inspect the TypeOf a for types that
			// support IsNil(); or, capture any panics that occur
			// and format the error nicely.
			return reflect.ValueOf(a).IsNil()
		},
	)
}

func (e *expectation) notImplemented() {
	e.Output.PrintFeature()
	e.Output.PrintContext()
	e.Output.PrintWhen()
	e.Output.PrintTitleNotImplemented()
}
