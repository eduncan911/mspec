package bdd

import (
	"fmt"
	"regexp"
	"runtime"
	"strings"
	"testing"

	"github.com/ddspog/mspec"
)

// Given defines the Feature's specific context to be spec'd out.
func Given(t *testing.T, given string, args ...interface{}) {
	stcs, givenLike := parseArgs(emptyS, args)
	whenFns := parseWhenFn(stcs)

	for _, givenS := range givenLike {
		// setup the spec that we will be using
		spec := &mspec.Specification{
			T:       t,
			Feature: featureDesc(2),
			Given:   fmtWithArgs(given, givenS),
		}
		spec.PrintFeature()
		spec.PrintContext()

		for _, whenFn := range whenFns {
			whenFn(func(when string, args ...interface{}) {
				stcs, whenLike := parseArgs(givenS, args)
				itFns := parseItFn(stcs)

				for _, whenS := range whenLike {
					spec.When = fmtWithArgs(when, whenS)
					spec.PrintWhen()

					for _, itFn := range itFns {
						itFn(func(it string, args ...interface{}) {
							stcs, itLike := parseArgs(whenS, args)
							assertFns := parseAssertFn(stcs)

							for _, itS := range itLike {
								spec.Spec = fmtWithArgs(it, itS)
								// Spec output is handled in the spec.Run() below

								if len(assertFns) > 0 {
									// Having at least 1 assert means we are implemented
									for _, assertFn := range assertFns {
										spec.AssertFn = func(a mspec.Assert) {
											assertFn(a, itS...)
										}

										spec.NotImplemented = false
									}
								} else {
									spec.AssertFn = notImplemented()
									spec.NotImplemented = true
								}

								// Run() handles contextual printing and some delegation
								// to the Assert's implementation for error handling
								spec.Run()
							}

						}, whenS...)
					}
				}
			}, givenS...)
		}

		// reset to default
		mspec.Config().ResetLasts()

		if mspec.Config().Output != mspec.OutputNone {
			fmt.Println()
		}
	}
}

// When defines the action or event when Given a specific context.
type When func(when string, args ...interface{})

// It defines the specification of When something happens.
type It func(title string, args ...interface{})

// Assert defines the action of asserting things during test.
type Assert interface {
	mspec.Assert
}

// S defines a set of arguments, to run on Given, When or It sentences.
type S []interface{}

// NewS return a new set of arguments, given on the function.
func NewS(args ...interface{}) (s S) {
	s = args
	return
}

var emptyS = []interface{}{}

// Like defines a set of environments to be run on a sentence like
// Given, When and It. It receives a list of sets of arguments, and
// those arguments will be used to conduct table-driven tests using
// this BDD framework.
func Like(sets ...S) []S {
	return sets
}

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

// Sentences return the functions Given, Like and NewS with new names,
// for convenience of developer.
func Sentences() (given func(*testing.T, string, ...interface{}), like func(...S) []S, s func(...interface{}) S) {
	given = Given
	like = Like
	s = NewS
	return
}

func fmtWithArgs(s string, args []interface{}) (f string) {
	if ok, _ := regexp.MatchString(".*\\%\\[[0-9]*\\](v|s).*", s); ok {
		f = fmt.Sprintf(s, args...)
	} else {
		f = s
	}
	return
}

func parseArgs(initSet S, args []interface{}) (stc []interface{}, like []S) {
	like = []S{initSet}

	switch len(args) {
	case 0:
		stc = []interface{}{}
	case 1:
		switch args[0].(type) {
		case []S:
			stc = []interface{}{}
			like = args[0].([]S)
		default:
			stc = []interface{}{args[0]}
		}
	default:
		n := len(args)
		stc = args[:n-1]
		like = args[n-1].([]S)
	}

	return
}

func parseWhenFn(args []interface{}) (whenFns []func(When, ...interface{})) {
	whenFns = make([]func(When, ...interface{}), len(args))
	for i := range args {
		switch v := args[i].(type) {
		case func(When):
			whenFns[i] = func(wh When, args ...interface{}) {
				v(wh)
			}
		default:
			whenFns[i] = v.(func(When, ...interface{}))
		}
	}
	return
}

func parseItFn(args []interface{}) (itFns []func(It, ...interface{})) {
	itFns = make([]func(It, ...interface{}), len(args))
	for i := range args {
		switch v := args[i].(type) {
		case func(It):
			itFns[i] = func(it It, args ...interface{}) {
				v(it)
			}
		default:
			itFns[i] = v.(func(It, ...interface{}))
		}
	}
	return
}

func parseAssertFn(args []interface{}) (assertFns []func(Assert, ...interface{})) {
	assertFns = make([]func(Assert, ...interface{}), len(args))
	for i := range args {
		switch v := args[i].(type) {
		case func(Assert):
			assertFns[i] = func(as Assert, args ...interface{}) {
				v(as)
			}
		default:
			assertFns[i] = v.(func(Assert, ...interface{}))
		}
	}
	return
}

// notImplemented is used to mark a specification that needs coding out.
var notImplemented = func() func(mspec.Assert) {
	return func(assert mspec.Assert) {
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
