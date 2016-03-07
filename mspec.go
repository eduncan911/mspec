package mspec

import (
	"strings"

	"github.com/eduncan911/go-mspec/colors"
)

var config *Config

// Config defines the configuration used by the package.
type Config struct {
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
func SetConfig(c Config) {
	config = &c
}

// ResetConfig will reset all options back to their default configuration.
// Useful for custom colors in the middle of a specification.
func ResetConfig() {
	// setup a default configuration
	config = &Config{
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

func (c *Config) resetLasts() {
	c.lastGiven = ""
	c.lastWhen = ""
	c.lastSpec = ""
}
