package gomspec

import (
	"testing"
)

type output struct {
	featurePrinted bool
	contextPrinted bool
	whenPrinted    bool
	titlePrinted   bool
	errors         string
}

func (out *output) PrintFeature() {
	out.featurePrinted = true
}

func (out *output) PrintContext() {
	out.contextPrinted = true
}

func (out *output) PrintWhen() {
	out.whenPrinted = true
}

func (out *output) PrintTitle() {
	out.titlePrinted = true
}

func (out *output) PrintTitleNotImplemented() {
	out.titlePrinted = true
}

func (out *output) PrintTitleWithError() {
	out.titlePrinted = true
}

func (out *output) PrintError(err string) {
	out.errors += err
}

// Prints title, even when test passes
func TestMatcherPrintsTitle(t *testing.T) {
	out := new(output)
	e := &Expectation{1, out}
	e.ToEqual(1)
	if titlePrinted := out.titlePrinted; !titlePrinted {
		t.Errorf("should have printed title, but didn't")
	}
}

func Test_Custom_Matchers(t *testing.T) {

	Given(t, "a need to test a number that differs by one from another number", func(when When) {

		// Target: an instance of &expectation
		out := new(output)
		e := &Expectation{4, out}

		// a custom matcher
		differsByOne := func(a, b interface{}) bool {
			diff := a.(int) - b.(int)
			return diff == 1 || diff == -1
		}

		when("4 differs by one from 5", func(it It) {

			// invoke the custom matcher
			e.To("differ by one from", 5, differsByOne)

			it("should not have any errs in the output", func(expect Expect) {
				err := out.errors
				expect(err).ToEqual("")
			})
		})

		when("4 does not differ by one from 10", func(it It) {

			// invoke the custom matcher
			e.To("differ by one from", 10, differsByOne)

			it("should have a specific error message", func(expect Expect) {
				err := out.errors
				expect(err).ToEqual("Expected `4` to differ by one from `10`")
			})
		})
	})

	Given(t, "a need to test a number that divisible by a number", func(when When) {

		// Target: an instance of &expectation
		out := new(output)
		e := &Expectation{9, out}

		// a custom matcher
		divisibleBy := func(a, b interface{}) bool {
			return a.(int)%b.(int) == 0
		}

		when("9 is divisible by 3", func(it It) {

			// invoke the custom matcher
			e.To("be divisible by", 3, divisibleBy)

			it("should not have any errs in the output", func(expect Expect) {
				err := out.errors
				expect(err).ToEqual("")
			})
		})

		when("9 is not divisible by 2", func(it It) {

			// invoke the custom matcher
			e.To("be divisible by", 2, divisibleBy)

			it("should have a specific error message", func(expect Expect) {
				err := out.errors
				expect(err).ToEqual("Expected `9` to be divisible by `2`")
			})
		})
	})
}

type object struct {
	ref   *string
	ref2  *string
	value string
}

func Test_Expectations(t *testing.T) {

	Given(t, "a struct value", func(when When) {

		value := "some text"
		o := &object{
			ref:   &value,
			value: value,
		}

		when("evaluating expect(o).ToNotBeNil()", func(it It) {
			it("should throw an error", NA())
			it("should say that ToNotBeNil() only works on pointers", NA())
		})

		when("evaluating expect(o2).ToBeNil()", func(it It) {
			it("should throw an error", NA())
			it("should say that ToBeNil() only works on pointers", NA())
		})

		when("evaluating expect(o2).ToEqual(o)", func(it It) {

			var o2 = o

			it("should be equal", func(expect Expect) {
				expect(o2).ToEqual(o)
			})
		})

		when("evaluating expect(o2).ToNotEqual(o)", func(it It) {

			var o2 = &object{
				ref:   &value,
				value: value,
			}

			it("should not be equal", func(expect Expect) {
				expect(o2).ToNotEqual(o)
			})
		})
	})

	Given(t, "a pointer of a value", func(when When) {

		value := "some text"
		o := &object{
			ref:   &value,
			value: value,
		}

		when("evaluating expect(o.ref).ToNotBeNil()", func(it It) {
			it("should not be nil", func(expect Expect) {
				expect(o.ref).ToNotBeNil()
			})
		})

		when("evaluating expect(o.ref2).ToBeNil()", func(it It) {
			it("should be nil", func(expect Expect) {
				expect(o.ref2).ToBeNil()
			})
		})

		when("evaluating expect(o.ref).ToEqual(o.ref)", func(it It) {
			it("should be equal", func(expect Expect) {
				expect(o.ref).ToEqual(o.ref)
			})
		})

		when("evaluating expect(o.ref).ToNotEqual(o.ref2)", func(it It) {
			it("should not be equal", func(expect Expect) {
				expect(o.ref).ToNotEqual(o.ref2)
			})
		})
	})

	Given(t, "an int value", func(when When) {

		i := 5

		when("evaluating expect(i).ToNotBeNil()", func(it It) {
			it("should throw an error", NA())
			it("should say that ToNotBeNil() only works on pointers", NA())
		})

		when("evaluating expect(i).ToBeNil()", func(it It) {
			it("should throw an error", NA())
			it("should say that ToBeNil() only works on pointers", NA())
		})

		when("evaluating expect(5).ToEqual(5)", func(it It) {
			it("should be equal", func(expect Expect) {
				expect(i).ToEqual(5)
			})
		})

		when("evaluating expect(5).ToNotEqual(6)", func(it It) {
			it("should not be equal", func(expect Expect) {
				expect(i).ToNotEqual(6)
			})
		})

	})

	Given(t, "a string value", func(when When) {

		s := "hello"

		when("evaluating expect(s).ToNotBeNil()", func(it It) {
			it("should throw an error", NA())
			it("should say that ToNotBeNil() only works on pointers", NA())
		})

		when("evaluating expect(s).ToBeNil()", func(it It) {
			it("should throw an error", NA())
			it("should say that ToBeNil() only works on pointers", NA())
		})

		when("evaluating expect(s).ToEqual(s)", func(it It) {
			it("should be equal", func(expect Expect) {
				expect(s).ToEqual("hello")
			})
		})

		when("evaluating expect(s).ToNotEqual(\"world\")", func(it It) {
			it("should not be equal", func(expect Expect) {
				expect(s).ToNotEqual("world")
			})
		})
	})
}
