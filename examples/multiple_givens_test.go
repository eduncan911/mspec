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

	Given(t, "a rabbi, a priest, and a Lutheran minister", func(when When) {

		when("they walk into a bar", func(it It) {

			it("should have the bartender ask, \"Is this some kind of a joke?\"", func(expect Expect) {
				expect(true).ToEqual(true)
			})

			it("should have at least 1 laughing audience member", func(expect Expect) {
				expect(true).ToEqual(true)
			})
		})
	})

	Given(t, "a Screwdriver walks into a bar", func(when When) {

		when("the bartender says, \"Hey, we have a drink named after you!\"", func(it It) {

			it("should have the Screwdriver saying, \"You have a drink named Murray?\"", func(expect Expect) {
				expect(true).ToEqual(true)
			})
		})
	})

	Given(t, "a Horse", func(when When) {

		when("it walks into a bar", func(it It) {

			it("should have the bartender ask, \"Hey, why the long face?\"", func(expect Expect) {
				expect(true).ToEqual(true)
			})
		})
	})

	/* Outputs:

	   Feature: Example of Multiple Givens

	     Given a rabbi, a priest, and a Lutheran minister

	       When they walk into a bar
	       » It should have the bartender ask, "Is this some kind of a joke?"
	       » It should have at least 1 laughing audience member

	     Given a Screwdriver walks into a bar

	       When the bartender says, "Hey, we have a drink named after you!"
	       » It should have the Screwdriver saying, "You have a drink named Murray?"

	     Given a Horse

	       When it walks into a bar
	       » It should have the bartender ask, "Hey, why the long face?"
	*/
}
