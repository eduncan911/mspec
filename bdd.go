package gomspec

import (
	"fmt"
	"runtime"
	"strings"
	"testing"
)

type specification struct {
	T       *testing.T
	Feature string
	Context string
	When    string
	Title   string
	Fn      func(Expect)
}

func (spec *specification) run() {
	spec.Fn(func(val interface{}) *expectation {
		return &expectation{val, spec}
	})
}

// Given defines the Feature's specific context to be spec'd out.
func Given(t *testing.T, context string, scenerioWrapper func(When)) {

	scenerioWrapper(func(when string, testWrapper func(It)) {
		testWrapper(func(it string, fn func(Expect)) {
			spec := &specification{
				t,
				featureDesc(6),
				context,
				when,
				it,
				fn,
			}
			spec.run()
		})
	})

	// reset to default
	mspec.resetLasts()
}

// When defines the action or event when Given a specific context.
type When func(when string, fn func(It))

// It defines the specification of when something happens.
type It func(title string, fn func(Expect))

// Setup is used to define before/after (setup/teardown) functions.
func Setup(before, after func()) func(fn func(Expect)) func(Expect) {
	return func(fn func(Expect)) func(Expect) {
		before()
		return func(expect Expect) {
			fn(expect)
			after()
		}
	}
}

// NotImplemented is used to mark a specification that needs coding out.
func NotImplemented() func(Expect) {
	return func(expect Expect) { expect(nil).notImplemented() }
}

// NA is shorthand for the NotImplemented() function.
func NA() func(Expect) {
	return NotImplemented()
}

// Desc is legacy support for existing Zen users.
func Desc(t *testing.T, desc string, wrapper func(It)) {
	wrapper(func(it string, fn func(Expect)) {
		spec := &specification{
			t,
			featureDesc(4),
			"",
			desc,
			it,
			fn,
		}
		spec.run()
	})
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
