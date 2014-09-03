package examples

import (
	. "github.com/eduncan911/gomspec"
	"testing"
)

func Test_Multiple_Givens(t *testing.T) {

	// example of grouping different given contexts in the same feature.
	// this is especially useful for spec'ing out entire features without
	// switching mental contexts by staying in the same func definition.
	//

	Given(t, "no dogs available", func(when When) {

		when("creating a new dog", func(it It) {

			d := BirthDog()

			it("should be a normal color", func(expect Expect) {
				expect(d.color).ToEqual(normalColor)
			})

			it("should not need washing", func(expect Expect) {
				expect(d.washed).ToEqual(false)
			})
		})
	})

	Given(t, "an unpainted dog", func(when When) {

		d := BirthDog()
		colorToPaint := "green"

		when("painting the dog a permanent green", func(it It) {

			d.Paint(&paint{
				color:      colorToPaint,
				iswashable: false,
			})

			it("should have paint on it", func(expect Expect) {
				expect(d.paint).ToExist()
			})

			it("should be the color green", func(expect Expect) {
				expect(d.color).ToEqual(colorToPaint)
			})

			it("should not be washed", func(expect Expect) {
				expect(d.washed).ToEqual(false)
			})
		})
	})

	Given(t, "a painted dog", func(when When) {

		d := BirthDog()
		d.Paint(&paint{
			color:      "red",
			iswashable: true,
		})

		when("washing the dog", func(it It) {

			d.Wash()

			it("should have the paint come off", func(expect Expect) {
				expect(d.paint).ToNotExist()
			})

			it("should be a normal color", func(expect Expect) {
				expect(d.color).ToEqual(normalColor)
			})

			it("should smell like a clean dog", func(expect Expect) {
				expect(d.washed).ToEqual(true)
			})
		})
	})

	/* Outputs:

	Feature: Multiple Givens

	  Given no dogs available

	    When creating a new dog
	    » It should be a normal color
	    » It should not need washing

	  Given an unpainted dog

	    When painting the dog a permanent green
	    » It should should have paint
	    » It should be the color green
	    » It should not be washed

	  Given a painted dog

	    When washing the dog
	    » It should have the paint come off
	    » It should be a normal color
	    » It should smell like a clean dog
	*/
}
