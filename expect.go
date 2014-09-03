package gomspec

import (
	"fmt"
	"reflect"
	"time"
)

// Expect is a light assertion pattern for testing currently embedded
// GoMspec's code.  Future work may be to expose it as an interface
// to allow registering custom matchers (e.g. testify's assert package).
type Expect func(val interface{}) *Expectation

// Expectation represents the value being asserted.
type Expectation struct {
	Value  interface{}
	Output formatter
}

// To tests a custom matching interface to the value.
func (e *Expectation) To(desc string, value interface{}, matcher func(a, b interface{}) bool) {
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

// ToEqual tests the equality of the Expectation to the value of b.
func (e *Expectation) ToEqual(b interface{}) {
	e.To(
		"to equal",
		b,
		func(a, b interface{}) bool {
			return objectsAreEqual(a, b)
		},
	)
}

// ToNotEqual tests the inequality of the Expectation to the value of b.
func (e *Expectation) ToNotEqual(b interface{}) {
	e.To(
		"to equal",
		b,
		func(a, b interface{}) bool {
			return !objectsAreEqual(a, b)
		},
	)
}

// ToMatchExactly matches both the type and the values.
func (e *Expectation) ToMatchExactly(b interface{}) {
	e.To(
		"to exactly match values and type of",
		b,
		func(a, b interface{}) bool {
			aType := reflect.TypeOf(a)
			bType := reflect.TypeOf(b)

			if aType != bType {
				return false
			}

			return objectsAreEqual(a, b)
		},
	)
}

func objectsAreEqual(expected, actual interface{}) bool {

	if expected == nil || actual == nil {
		return expected == actual
	}

	if reflect.DeepEqual(expected, actual) {
		return true
	}

	expectedValue := reflect.ValueOf(expected)
	actualValue := reflect.ValueOf(actual)
	if expectedValue == actualValue {
		return true
	}

	// Attempt comparison after type conversion
	if actualValue.Type().ConvertibleTo(expectedValue.Type()) && expectedValue == actualValue.Convert(expectedValue.Type()) {
		return true
	}

	// Last ditch effort
	if fmt.Sprintf("%#v", expected) == fmt.Sprintf("%#v", actual) {
		return true
	}

	return false

}

// ToExist tests that expectation is not equal to nil.
func (e *Expectation) ToExist() {
	e.To(
		"exist",
		nil,
		func(a, b interface{}) bool {

			success := true

			if a == nil {
				success = false
			} else {
				value := reflect.ValueOf(a)
				kind := value.Kind()
				if kind >= reflect.Chan && kind <= reflect.Slice && value.IsNil() {
					success = false
				}
			}

			return success
		},
	)
}

// ToNotExist tests that the expectation is equal to nil.
func (e *Expectation) ToNotExist() {
	e.To(
		"not exist",
		nil,
		func(a, b interface{}) bool {
			return isNil(a)
		},
	)
}

func isNil(object interface{}) bool {
	if object == nil {
		return true
	}

	value := reflect.ValueOf(object)
	kind := value.Kind()
	if kind >= reflect.Chan && kind <= reflect.Slice && value.IsNil() {
		return true
	}

	return false
}

// ToBeEmpty evaluates types for zero values and returns true if completely empty.
func (e *Expectation) ToBeEmpty() {
	e.To(
		"",
		nil,
		func(a, b interface{}) bool {
			return isEmpty(a)
		},
	)
}

// ToNotBeEmpty evaluates types for zero values and returns true if completely empty.
func (e *Expectation) ToNotBeEmpty() {
	e.To(
		"",
		nil,
		func(a, b interface{}) bool {
			return !isEmpty(a)
		},
	)
}

var zeros = []interface{}{
	int(0),
	int8(0),
	int16(0),
	int32(0),
	int64(0),
	uint(0),
	uint8(0),
	uint16(0),
	uint32(0),
	uint64(0),
	float32(0),
	float64(0),
}

// isEmpty gets whether the specified object is considered empty or not.
func isEmpty(object interface{}) bool {

	if object == nil {
		return true
	} else if object == "" {
		return true
	} else if object == false {
		return true
	}

	for _, v := range zeros {
		if object == v {
			return true
		}
	}

	objValue := reflect.ValueOf(object)

	switch objValue.Kind() {
	case reflect.Map:
		fallthrough
	case reflect.Slice, reflect.Chan:
		{
			return (objValue.Len() == 0)
		}
	case reflect.Ptr:
		{
			switch object.(type) {
			case *time.Time:
				return object.(*time.Time).IsZero()
			default:
				return false
			}
		}
	}
	return false
}

func (e *Expectation) notImplemented() {
	e.Output.PrintFeature()
	e.Output.PrintContext()
	e.Output.PrintWhen()
	e.Output.PrintTitleNotImplemented()
}
