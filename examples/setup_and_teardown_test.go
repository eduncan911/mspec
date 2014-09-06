package examples

import (
	. "github.com/eduncan911/gomspec"
	"testing"
)

/*
	There are two forms of setting up and teardown a context
	before exact specification is run:

		* On a Per Context basis, where the same instance of the
		  context is shared across all specs.

		* on a Per Spec basis, where the Setup and Teardown of
		  the context happens for each and every Spec that is run.

	MSpec supports the first one by default.

	Optionally, you can use the Setup(before, after) feature to specify
	a context to setup a context before each spec is run, as well a
	teardown method to run after each spec is run.
*/

func Test_Setup_Shared_Context(t *testing.T) {

	// this example shows a shared context
	//

	Given(t, "a dog that has been painted\nand the paint is washable", func(when When) {

		// the target we are testing (in TDD terms)
		//
		d := BirthDog()

		// setup the shared context that runs just once
		//
		p := &paint{
			color:      "red",
			iswashable: true,
		}
		d.Paint(p)

		when("washing the dog", func(it It) {

			err := d.Wash()

			it("should not have an error", func(assert Assert) {
				assert.NoError(err)
			})

			it("should back to normal color", func(assert Assert) {
				assert.Equal(normalColor, d.color)
			})

			it("should have timesWashed be only 1 time", func(assert Assert) {

				// if the context ran only once, then this should be 1
				assert.Equal(1, d.timesWashed)
			})

		})
	})

	/* Output:

	Feature: Shared Context
	  Given a dog that has been painted
	  and the paint is washable
	    When washing the dog
	    » It should not have an error
	    » It should back to normal color
	    » It should have timesWashed be only 1 time

	*/

}

func Test_Setup_NonShared_Context(t *testing.T) {

	// this example shows a non-shared context in that the context
	// will always be run before the test, and then the teardown
	// will deconstruct the context after the test.
	//

	// TODO implement
	Given(t, "a healthy dog", func(when When) {

		when("visiting the vet", func(it It) {

			it("should have its blood checked")

			it("should have its teeth checked")

		})

	})

}
