package examples

import (
	. "github.com/eduncan911/gomspec"
	"github.com/eduncan911/gomspec/colors"
	"testing"
)

/*
	MSpec supports a number of configuration options.

		* Setting the color output to the console.
		* Changing the Assertion package.
		* Quiet option (surpress console output) <- NOT IMPLEMENTED
		* HTML output <- NOT IMPLEMENTED
		* Concurrent execution of specs (default is off) <- NOT IMPLEMENTED

	You can change these by accessing the global variable MSpec.

*/

func Test_Changing_Console_Colors(t *testing.T) {

	// you can use any color library that writes color codes to the console.
	// included is a library i wrote under gomspec/colors, as you can note
	// in the import statement above.
	//
	MSpec.AnsiOfFeature = colors.LightMagenta
	MSpec.AnsiOfGiven = colors.LightMagenta
	MSpec.AnsiOfWhen = colors.LightMagenta
	MSpec.AnsiOfThen = colors.LightMagenta

	Given(t, "that we want the output in all purple", func(when When) {

		when("the configuration is changed", func(it It) {

			it("should have everything be purple", func(assert Assert) {})

		})
	})
}

func Test_Console_Default_Colors(t *testing.T) {

	// resets to default
	MSpec.Defaults()

	Given(t, "that we want to reset everything back to defaults", func(when When) {

		when("the Defaults configuration is called", func(it It) {

			it("should be back to default colors", func(assert Assert) {})
		})
	})
}
