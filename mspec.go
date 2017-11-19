package mspec

import (
	"strings"

	"github.com/eduncan911/go-mspec/colors"
)

var config *MSpecConfig

// MSpecConfig defines the configuration used by the package.
type MSpecConfig struct {
	output outputType

	AnsiOfFeature            string
	AnsiOfGiven              string
	AnsiOfWhen               string
	AnsiOfThen               string
	AnsiOfThenNotImplemented string
	AnsiOfThenWithError      string
	AnsiOfCode               string
	AnsiOfCodeError          string
	AnsiOfExpectedError      string

	assertFn func(*Specification) Assert

	lastFeature string
	lastGiven   string
	lastWhen    string
	lastSpec    string
}

func init() {
	ResetConfig()

	// set to verbose output by default
	SetVerbose()

	// register the default Assertions package
	AssertionsFn(func(s *Specification) Assert {
		return newAssertions(s)
	})
}

// AssertionsFn will assign the assertions used for all tests.
// The specified struct must implement the mspec.Assert interface.
//
//    mspec.AssertionsFn(func(s *Specification) Assert {
//	    return &MyCustomAssertions{}
//    })
func AssertionsFn(fn func(s *Specification) Assert) {
	config.assertFn = fn
}

// SetConfig takes a Config instance and will be used for all tests
// until ResetConfig() is called.
//
//    mspec.SetConfig(Config{
//      AnsiOfFeature: "",	// remove color coding for Feature
//    })
//
func SetConfig(c MSpecConfig) {
	config = &c
}

// ResetConfig will reset all options back to their default configuration.
// Useful for custom colors in the middle of a specification.
func ResetConfig() {
	// setup a default configuration
	config = &MSpecConfig{
		AnsiOfFeature:            strings.Join([]string{colors.White}, ""),
		AnsiOfGiven:              strings.Join([]string{colors.Grey}, ""),
		AnsiOfWhen:               strings.Join([]string{colors.LightGreen}, ""),
		AnsiOfThen:               strings.Join([]string{colors.Green}, ""),
		AnsiOfThenNotImplemented: strings.Join([]string{colors.LightYellow}, ""),
		AnsiOfThenWithError:      strings.Join([]string{colors.RegBg, colors.White, colors.Bold}, ""),
		AnsiOfCode:               strings.Join([]string{colors.Grey}, ""),
		AnsiOfCodeError:          strings.Join([]string{colors.White, colors.Bold}, ""),
		AnsiOfExpectedError:      strings.Join([]string{colors.Red}, ""),
	}
}

// SetVerbose is used to set the output to Stdout (default).
// Do not use this at this time.  The package API
// will most likely change.
func SetVerbose() {
	config.output = outputStdout
}

// SetSilent is used to make all output silent.
// Do not use this at this time.  The package API
// will most likely change.
func SetSilent() {
	config.output = outputNone
}

type outputType int

const (
	outputNone outputType = 1 << iota
	outputStdout
	outputStderr
	outputHTML
)

func (c *MSpecConfig) resetLasts() {
	c.lastGiven = ""
	c.lastWhen = ""
	c.lastSpec = ""
}
