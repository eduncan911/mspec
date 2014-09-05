package gomspec

import (
	"fmt"
	"runtime"
	"strings"
	"testing"
)

// Specification holds the state of the context for a specific specification.
type Specification struct {
	T                       *testing.T
	Feature                 string
	Given                   string
	When                    string
	Spec                    string
	AssertFn                func(Assert)
	AssertionFailed         bool
	AssertionFailedMessages []string
}

func (spec *Specification) run() {

	// print our story before any assertion output (if any)
	spec.PrintFeature()
	spec.PrintContext()
	spec.PrintWhen()

	// execute the Assertion
	spec.AssertFn(MSpec.assertFn(spec))

	// if there was no error (which handles its own printing),
	// print the spec here.
	if !spec.AssertionFailed {
		spec.PrintSpec()
	}
}

// Given defines the Feature's specific context to be spec'd out.
func Given(t *testing.T, given string, whenFn func(When)) {

	whenFn(func(when string, itFn func(It)) {
		itFn(func(it string, assertFn func(Assert)) {

			spec := &Specification{
				T:               t,
				Feature:         featureDesc(6),
				Given:           given,
				When:            when,
				Spec:            it,
				AssertFn:        assertFn,
				AssertionFailed: false,
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

// TODO verify works
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
		// TODO implement
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
