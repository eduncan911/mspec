package mspec

import (
	"fmt"
	"runtime"
	"strings"
	"testing"
)

func (spec *Specification) run() {

	// execute the Assertion
	spec.AssertFn(config.assertFn(spec))

	// if there was no error (which handles its own printing),
	// print the spec here.
	if spec.notImplemented {
		spec.PrintSpecNotImplemented()
	} else if !spec.AssertionFailed {
		spec.PrintSpec()
	}
}

// Given defines the Feature's specific context to be spec'd out.
func Given(t *testing.T, given string, when ...func(When)) {

	// setup the spec that we will be using
	spec := &Specification{
		T:       t,
		Feature: featureDesc(2),
		Given:   given,
	}
	spec.PrintFeature()
	spec.PrintContext()

	for _, whenFn := range when {
		whenFn(func(when string, its ...func(It)) {

			spec.When = when
			spec.PrintWhen()

			for _, itFn := range its {
				itFn(func(it string, assertFns ...func(Assert)) {

					spec.Spec = it
					// Spec output is handled in the spec.run() below

					if len(assertFns) > 0 {
						// having at least 1 assert means we are implemented
						for _, assertFn := range assertFns {
							spec.AssertFn = assertFn
							spec.notImplemented = false
						}
					} else {
						spec.AssertFn = notImplemented()
						spec.notImplemented = true
					}

					// run() handles contextual printing and some delegation
					// to the Assert's implemention for error handling
					spec.run()
				})
			}
		})
	}

	// reset to default
	config.resetLasts()

	fmt.Println()
}

// When defines the action or event when Given a specific context.
type When func(when string, it ...func(It))

// It defines the specification of When something happens.
type It func(title string, assert ...func(Assert))

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

// notImplemented is used to mark a specification that needs coding out.
var notImplemented = func() func(Assert) {
	return func(assert Assert) {
		// nothing to do here
	}
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
