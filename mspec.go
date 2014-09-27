package mspec

import (
	"github.com/eduncan911/gomspec/colors"
	"strings"
)

// MSpec is global config for the package.
var MSpec *MSpecConfig

// MSpecConfig defines the configurations and registrations for package.
type MSpecConfig struct {
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
	MSpec = &MSpecConfig{}
	MSpec.Defaults()
}

// AssertionsFn will assign the assertions used for all tests.
// MyCustomAsserts must implement the gomspec.Assert interface.
//
//    MSpec.RegisterAssertions(func(s *Specification) Assert {
//        return &MyCustomAssertions{}
//    })
func (c *MSpecConfig) AssertionsFn(fn func(s *Specification) Assert) {
	c.assertFn = fn
}

// Defaults will reset all options back to their default configuration.
// Useful for custom colors in the middle of a specification.  Do note
// that this will also change the assertions package back to the default
// Testify module.
func (c *MSpecConfig) Defaults() {

	// setup a default configuration
	MSpec = &MSpecConfig{
		AnsiOfFeature:            strings.Join([]string{colors.White}, ""),
		AnsiOfGiven:              strings.Join([]string{colors.Grey}, ""),
		AnsiOfWhen:               strings.Join([]string{colors.LightGreen}, ""),
		AnsiOfThen:               strings.Join([]string{colors.Green}, ""),
		AnsiOfThenNotImplemented: strings.Join([]string{colors.LightYellow}, ""),
		AnsiOfThenWithError:      strings.Join([]string{colors.RegBg, colors.White, colors.Bold}, ""),
		AnsiOfCode:               strings.Join([]string{colors.DarkGrey}, ""),
		AnsiOfCodeError:          strings.Join([]string{colors.White, colors.Bold}, ""),
		AnsiOfExpectedError:      strings.Join([]string{colors.Red}, ""),
	}

	// register the default Assertions package
	MSpec.AssertionsFn(func(s *Specification) Assert {
		return newAssertions(s)
	})
}

func (c *MSpecConfig) resetLasts() {
	c.lastGiven = ""
	c.lastWhen = ""
	c.lastSpec = ""
}
