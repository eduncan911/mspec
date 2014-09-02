package examples

import (
	. "github.com/eduncan911/gomspec"
	"testing"
)

func Test_Stubbing_a_Feature(t *testing.T) {

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
