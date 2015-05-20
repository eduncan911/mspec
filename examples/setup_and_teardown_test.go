package examples

import (
	. "github.com/eduncan911/go-mspec"
	"testing"
)

/*
	There are two forms to setup and teardown a context for each
	specification:

		* on a Per Context basis, where the same instance of the
		  context is shared across all specs and assumes each specification
		  will not mutex state.

		* on a Per Spec basis, where the Setup and Teardown of
		  the context happens for each specification - mutating state before
		  and after each spec is run.

	MSpec supports the first one by default.

	Optionally, you can use the Setup(before, after) feature to specify
	a context to setup a context before each spec is run, as well a
	teardown method to run after each spec is run.
*/

func Test_Setup_Shared_Context(t *testing.T) {

	// this example shows the default shared context does not mutate.
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

	// this example shows a mutable context in that the context
	// will always be run before the test, and then the teardown
	// will deconstruct the context after the test.
	//
	// these types of tests are useful for setting up database connections
	// or some other external dependencies that must be satisfied before
	// the tests can run.
	//

	// TODO implement
	Given(t, "a healthy dog after 1 year since last checkup", func(when When) {

		d := BirthDog() // dog hasn't taken any steps yet

		before := func() {
			d.steps++ // dog must take at least 1 step before each spec
		}

		after := func() {
			d.steps++ // dog takes another step after each spec
		}

		setup := Setup(before, after)

		when("visiting the vet", func(it It) {

			d.VisitVet()

			it("should have taken 1 step", setup(func(assert Assert) {
				// first spec makes the dog take 1 step in setup()
				assert.Equal(1, d.steps)
			}))

			it("should have taken 3 steps by now", setup(func(assert Assert) {
				// before++ and after++ and before++ == 3
				assert.Equal(3, d.steps)
			}))

			it("should have taken 4 steps total and no more setups", func(assert Assert) {
				// because we are not using setup() here, the state
				// of the dog's steps does not mutate.
				assert.Equal(4, d.steps)
			})
		})
	})

	/* Outputs:

	Feature: Setup NonShared Context
	  Given a healthy dog after 1 year since last checkup
	    When visiting the vet
	    » It should have taken 1 step
	    » It should have taken 3 steps by now
	    » It should have taken 4 steps total and no more setups

	*/

}
