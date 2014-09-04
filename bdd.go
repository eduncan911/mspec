package gomspec

import (
	"fmt"
	"runtime"
	"strings"
	"testing"
)

// Specification holds the state of the context for a specific specification.
type Specification struct {
	T       *testing.T
	Feature string
	Given   string
	When    string
	Spec    string
	Fn      func(Assert)
}

func (spec *Specification) run() {

	// maybe print everything from here now?
	spec.PrintFeature()
	spec.PrintContext()
	spec.PrintWhen()

	// TODO need a way to test for the assert failing here,
	// then control which spec result we print.
	spec.PrintSpec()

	spec.Fn(MSpec.assertFn(spec))
}

// Given defines the Feature's specific context to be spec'd out.
func Given(t *testing.T, given string, whenFn func(When)) {

	whenFn(func(when string, itFn func(It)) {
		itFn(func(it string, assertFn func(Assert)) {

			spec := &Specification{
				t,
				featureDesc(6),
				given,
				when,
				it,
				assertFn,
			}
			spec.run()

		})
	})

	// reset to default
	MSpec.resetLasts()
}

// When defines the action or event when Given a specific context.
type When func(when string, fn func(It))

// It defines the specification of when something happens.
type It func(title string, fn func(Assert))

// Setup is used to define before/after (setup/teardown) functions.
func Setup(before, after func()) func(fn func(Assert)) func(Assert) {
	return func(fn func(Assert)) func(Assert) {
		before()
		return func(assert Assert) {
			fn(assert)
			after()
		}
	}
}

// NotImplemented is used to mark a specification that needs coding out.
func NotImplemented() func(Assert) {
	return func(assert Assert) {
		//expect(nil).notImplemented()
	}
}

// NA is shorthand for the NotImplemented() function.
func NA() func(Assert) {
	return NotImplemented()
}

var featureDesc = func(callerDepth int) string {
	pc, _, _, _ := runtime.Caller(callerDepth)
	m := fmt.Sprintf("%s", runtime.FuncForPC(pc).Name())
	i := strings.LastIndex(m, ".")
	m = m[i+1 : len(m)]
	m = strings.Replace(m, "Test_", "", 1)
	m = strings.Replace(m, "Test", "", 1)
	return strings.Replace(m, "_", " ", -1)
}
