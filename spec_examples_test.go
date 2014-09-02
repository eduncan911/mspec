package gomspec

import (
	"testing"
)

func Test_Example_of_Given_this_And_this_And_this(t *testing.T) {

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

func Test_Example_of_Multiple_Givens(t *testing.T) {

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

		when("when it walks into a bar", func(it It) {

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

	       When when it walks into a bar
	       » It should have the bartender ask, "Hey, why the long face?"
	*/
}

func Test_Example_of_Stubbing_an_entire_Feature(t *testing.T) {

	// a real-world example of a project I am working on where i need
	// a microservice to call out to a master api and ask for work to do.
	// this is what I practice: spec'ing out your application first given
	// as many scenarios as you can think of.  once you have tweaked them
	// to your liking and you think you covered all bases, then go write
	// some code!
	//
	// and btw, your specs are never perfect the first time around.  it's ok
	// to refine the scenarios and specs as you code for unknowns or "yeah,
	// I don't need that right now" (aka future work).
	//

	Given(t, "a valid set of command-line args", func(when When) {
		when("the main() func executes", func(it It) {
			it("should ping the master service", NA())
			it("should ask the master for work to do", NA())
		})
	})

	Given(t, "a ping request", func(when When) {
		when("the response is alive", func(it It) {
			it("should ask the master for work to do", NA())
		})

		when("an error is returned", func(it It) {
			it("should panic", NA())
			it("should exit", NA())
		})
	})

	Given(t, "a request for work to do", func(when When) {
		when("the response has work to do", func(it It) {
			it("should create a NewClient()", NA())
			it("should start the monitor's Ticker", NA())
			it("should start the client's process", NA())
		})

		when("no work is returned", func(it It) {
			it("should sleep for the defined polling interval", NA())
		})

		when("an error is returned", func(it It) {
			it("should sleep for the defined polling interval", NA())
		})
	})

	/*	Outputs:

		Feature: Example of Stubbing an entire Feature

		  Given a valid set of command-line args

		    When the main() func executes
		    » It should ping the master service «-- NOT IMPLEMENTED
		    » It should ask the master for work to do «-- NOT IMPLEMENTED

		  Given a ping request

		    When the response is alive
		    » It should ask the master for work to do «-- NOT IMPLEMENTED

		    When an error is returned
		    » It should panic «-- NOT IMPLEMENTED
		    » It should exit «-- NOT IMPLEMENTED

		  Given a request for work to do

		    When the response has work to do
		    » It should create a NewClient() «-- NOT IMPLEMENTED
		    » It should start the monitor's Ticker «-- NOT IMPLEMENTED
		    » It should start the client's process «-- NOT IMPLEMENTED

		    When no work is returned
		    » It should sleep for the defined polling interval «-- NOT IMPLEMENTED

		    When an error is returned
		    » It should sleep for the defined polling interval «-- NOT IMPLEMENTED
	*/
}
