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
	// It should be a normal color
	// It should smell like a clean dog
	//

	Given(t, "a dog that has been painted red\nand the paint is washable\nand no one has washed the dog yet", func(when When) {

		d := BirthDog()
		d.Paint(&paint{
			color:      "red",
			iswashable: true,
		})

		when("the dog is washed", func(it It) {

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

	/*	Outputs:

		Feature: Given this and this and this

		  Given a dog that has been painted red
		  and the paint is washable
		  and no one has washed the dog yet

		    When the dog is washed
		    » It should have the paint come off
		    » It should be a normal color
		    » It should smell like a clean dog
	*/
}
