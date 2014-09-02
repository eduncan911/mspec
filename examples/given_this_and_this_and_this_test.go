package examples

import (
	. "github.com/eduncan911/gomspec"
	"testing"
)

func Test_Given_this_and_this_and_this(t *testing.T) {

	// example of providing rich stories with multiple conditions.
	// this is done by using a \n new line character in the given
	// which is automatically padded for you.
	//
	// Given a dog that has been painted red
	// and the paint is washable
	// and no one has washed the dog yet
	// When the dog is washed
	// It should have the paint come off
	// It should smell like a clean dog
	//

	Given(t, "a dog that has been painted red\nand the paint is washable\nand no one has washed the dog yet", func(when When) {

		when("the dog is washed", func(it It) {

			it("should have the paint come off", func(expect Expect) {
				expect(true).ToEqual(true)
			})

			it("should smell like a clean dog", func(expect Expect) {
				expect(true).ToEqual(true)
			})

		})
	})

	/*	Outputs:

		Feature: Example of Given this And this And this

		  Given a dog that has been painted red
		  and the paint is washable
		  and no one has washed the dog yet

		    When the dog is washed
		    » It should have the paint come off
		    » It should smell like a clean dog
	*/
}
