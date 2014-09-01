package gomspec

import (
	"fmt"
)

// Expect is a light assertion pattern for testing currently embedded
// GoMspec's code.  Future work may be to expose it as an interface
// to allow registering custom matchers (e.g. testify's assert package).
type Expect func(val interface{}) *expectation

type matcher func(a, b interface{}) bool

type expectation struct {
	Output formatter
	Value  interface{}
}

// To tests a custom matching interface to the value.
func (e *expectation) To(desc string, match matcher, value interface{}) {
	e.Output.PrintFeature()
	e.Output.PrintContext()
	e.Output.PrintWhen()
	if !match(e.Value, value) {
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
		func(a, b interface{}) bool {
			return a == b
		},
		b,
	)
}

// ToNotEqual tests the inequality of the expectation to the value of b.
func (e *expectation) ToNotEqual(b interface{}) {
	e.To(
		"not equal",
		func(a, b interface{}) bool {
			return a != b
		},
		b,
	)
}

// ToExist tests that expectation is not equal to nil.
func (e *expectation) ToExist() {
	e.To(
		"exist",
		func(a, b interface{}) bool {
			return a != nil
		},
		nil,
	)
}

// ToNotExist tests that the expectation is equal to nil.
func (e *expectation) ToNotExist() {
	e.To(
		"not exist",
		func(a, b interface{}) bool {
			return a == nil
		},
		nil,
	)
}

func (e *expectation) notImplemented() {
	e.Output.PrintFeature()
	e.Output.PrintContext()
	e.Output.PrintWhen()
	e.Output.PrintTitleNotImplemented()
}
