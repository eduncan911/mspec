package mspec

import (
	"fmt"
	"io/ioutil"
	"path"
	"runtime"
	"strings"
	"testing"

	"github.com/eduncan911/go-mspec/colors"
)

type formatter interface {
	PrintFeature()
	PrintContext()
	PrintWhen()
	PrintTitle()
	PrintTitleNotImplemented()
	PrintTitleWithError()
	PrintError(string)
}

type failingLine struct {
	prev     string
	content  string
	next     string
	filename string
	number   int
	lines    []string
}

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

	NotImplemented bool
}

func (spec *Specification) PrintFeature() {
	if config.lastFeature == spec.Feature {
		return
	}
	if config.Output != OutputNone {
		fmt.Printf("%sFeature: %s%s\n", config.AnsiOfFeature, spec.Feature, colors.Reset)
	}
	config.lastFeature = spec.Feature
}

func (spec *Specification) PrintContext() {
	if config.lastGiven == spec.Given {
		return
	}
	if config.Output != OutputNone {
		fmt.Printf("%s  Given %s%s\n", config.AnsiOfGiven, padLf(spec.Given, 2), colors.Reset)
	}
	config.lastGiven = spec.Given
}

func (spec *Specification) PrintWhen() {
	if config.lastWhen == spec.When {
		return
	}
	if config.Output != OutputNone {
		fmt.Printf("%s    When %s%s\n", config.AnsiOfWhen, spec.When, colors.Reset)
	}
	config.lastWhen = spec.When
}

func (spec *Specification) PrintSpec() {
	if config.Output != OutputNone {
		fmt.Printf("%s    » It %s %s\n", config.AnsiOfThen, spec.Spec, colors.Reset)
	}
	config.lastSpec = spec.Spec
}

func (spec *Specification) PrintSpecWithError() {
	if config.lastSpec == spec.Spec {
		return
	}
	if config.Output != OutputNone {
		fmt.Printf("%s    » It %s %s\n", config.AnsiOfThenWithError, spec.Spec, colors.Reset)
	}
	config.lastSpec = spec.Spec
}

func (spec *Specification) PrintSpecNotImplemented() {
	if config.Output != OutputNone {
		fmt.Printf("%s    » It %s «-- NOT IMPLEMENTED%s\n", config.AnsiOfThenNotImplemented, spec.Spec, colors.Reset)
	}
	config.lastSpec = spec.Spec
}

func (spec *Specification) PrintError(message string) {
	failingLine, err := getFailingLine()

	if err != nil {
		return
	}
	if config.Output != OutputNone {
		fmt.Printf("%s%s%s\n", config.AnsiOfExpectedError, message, colors.Reset)
		fmt.Printf("%s        in %s:%d%s\n", config.AnsiOfCode, path.Base(failingLine.filename), failingLine.number, colors.Reset)
		fmt.Printf("%s        ---------\n", config.AnsiOfCode)
		fmt.Printf("%s        %d. %s%s\n", config.AnsiOfCode, failingLine.number-1, softTabs(failingLine.prev), colors.Reset)
		fmt.Printf("%s        %d. %s %s\n", config.AnsiOfCodeError, failingLine.number, failingLine.content, colors.Reset)
		fmt.Printf("%s        %d. %s%s\n", config.AnsiOfCode, failingLine.number+1, softTabs(failingLine.next), colors.Reset)
		fmt.Println()
		spec.T.Fail()
		fmt.Println()
	}
}

// Run handles contextual printing and some delegation
// to the Assert's implementation for error handling
func (spec *Specification) Run() {

	// execute the Assertion
	spec.AssertFn(config.assertFn(spec))

	// if there was no error (which handles its own printing),
	// print the spec here.
	if spec.NotImplemented {
		spec.PrintSpecNotImplemented()
	} else if !spec.AssertionFailed {
		spec.PrintSpec()
	}
}

func getFailingLine() (failingLine, error) {

	// this entire func is now a hack because of where it is being called,
	// which is now one caller higher.  previously it was being called in the
	// Expect struct which had the right caller info.  but now, it is being
	// called after the Assertion has been executed to print details to the
	// string.

	_, filename, ln, _ := runtime.Caller(5)

	// this is really hacky, need to find a way of not using magic numbers for runtime.Caller
	// If we are not in a test file, we must still be inside this package,
	// so we need to go up one more stack frame to get to the test file
	if !strings.HasSuffix(filename, "_test.go") {
		_, filename, ln, _ = runtime.Caller(6)
	}

	bf, err := ioutil.ReadFile(filename)

	if err != nil {
		return failingLine{}, fmt.Errorf("Failed to open %s", filename)
	}

	lines := strings.Split(string(bf), "\n")[ln-2 : ln+2]

	return failingLine{
		prev:     softTabs(lines[0]),
		content:  softTabs(lines[1]),
		next:     softTabs(lines[2]),
		filename: filename,
		number:   ln,
	}, nil

}

func softTabs(text string) string {
	return strings.Replace(text, "\t", "  ", -1)
}

func padLf(strToPad string, padding int) string {
	pad := func() string {
		s := "\n"
		for i := 0; i < padding; i++ {
			s = strings.Join([]string{s, " "}, "")
		}
		return s
	}
	return strings.Replace(
		strToPad,
		"\n",
		pad(),
		-1,
	)
}
