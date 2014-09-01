package gomspec

import (
	"fmt"
	"io/ioutil"
	"path"
	"runtime"
	"strings"
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
	content  string
	filename string
	next     string
	number   int
	prev     string
}

var lastFeature string
var lastContext string
var lastWhen string
var lastTitle string

func (spec *specification) PrintFeature() {
	if lastFeature == spec.Feature {
		return
	}
	fmt.Printf("%s%s%s\n", mspec.AnsiOfFeature, spec.Feature, reset)
	lastFeature = spec.Feature
}

func (spec *specification) PrintContext() {
	if lastContext == spec.Context {
		return
	}
	fmt.Printf("%s  Given %s%s\n", mspec.AnsiOfGiven, spec.Context, reset)
	lastContext = spec.Context
}

func (spec *specification) PrintWhen() {
	if lastWhen == spec.When {
		return
	}
	fmt.Printf("\n%s    When %s%s\n", mspec.AnsiOfWhen, spec.When, reset)
	lastWhen = spec.When
}

func (spec *specification) PrintTitle() {
	if lastTitle == spec.Title {
		return
	}
	fmt.Printf("%s    » It %s %s\n", mspec.AnsiOfThen, spec.Title, reset)
	lastTitle = spec.Title
}

func (spec *specification) PrintTitleWithError() {
	if lastTitle == spec.Title {
		return
	}
	fmt.Printf("%s    » It %s %s\n", mspec.AnsiOfThenWithError, spec.Title, reset)
	lastTitle = spec.Title
}

func (spec *specification) PrintTitleNotImplemented() {
	if lastTitle == spec.Title {
		return
	}
	fmt.Printf("%s    » It %s «-- NOT IMPLEMENTED%s\n", mspec.AnsiOfThenNotImplemented, spec.Title, reset)
	lastTitle = spec.Title
}

func (spec *specification) PrintError(message string) {
	spec.PrintTitle()
	failingLine, err := getFailingLine()

	if err != nil {
		return
	}

	fmt.Printf("%s      %s%s\n", mspec.AnsiOfExpectedError, message, reset)
	fmt.Printf("%s      %s:%d%s\n", mspec.AnsiOfCode, path.Base(failingLine.filename), failingLine.number, reset)
	spec.PrintFailingLine(&failingLine)
	spec.T.Fail()
}

func (spec *specification) PrintFailingLine(failingLine *failingLine) {
	fmt.Printf("%s        %d. %s%s\n", mspec.AnsiOfCode, failingLine.number-1, failingLine.prev, reset)
	fmt.Printf("%s        %d. %s %s\n", mspec.AnsiOfCodeError, failingLine.number, failingLine.content, reset)
	fmt.Printf("%s        %d. %s%s\n", mspec.AnsiOfCode, failingLine.number+1, failingLine.next, reset)
	fmt.Println()
}

func getFailingLine() (failingLine, error) {
	_, filename, ln, _ := runtime.Caller(3)
	// TODO: this is really hacky, need to find a way of not using magic numbers for runtime.Caller
	// If we are not in a test file, we must still be inside this package,
	// so we need to go up one more stack frame to get to the test file
	if !strings.HasSuffix(filename, "_test.go") {
		_, filename, ln, _ = runtime.Caller(4)
	}

	bf, err := ioutil.ReadFile(filename)

	if err != nil {
		return failingLine{}, fmt.Errorf("Failed to open %s", filename)
	}

	lines := strings.Split(string(bf), "\n")[ln-2 : ln+2]

	return failingLine{
		softTabs(lines[1]),
		filename,
		softTabs(lines[2]),
		int(ln),
		softTabs(lines[0]),
	}, nil

}

func softTabs(text string) string {
	return strings.Replace(text, "\t", "  ", -1)
}
