package examples

import (
	. "github.com/eduncan911/go-mspec"
	"testing"
)

func Test_Specing_A_New_Feature(t *testing.T) {

	// MSpec was designed to let you capture your free formed thoughts
	// as you dream up the next big idea.  forget all about coding, how
	// you are going to implement it, etc: just think on how you will
	// design the feature.
	//
	// then document it, using MSpec.  you can quickly spec new features
	// with little syntax noise using this framework.
	//
	// below shows an example of a through process of developing an API with
	// a few known methods.  someone has told us that the API doesn't handle
	// errors, so for this "Feature" we are going to spec out the error
	// handling.
	//

	// Given a valid Api call, what shall we do?  not sure yet.
	//
	Given(t, "a valid Api")

	// Given an invalid Api call...
	//
	Given(t, "an invalid Api", func(when When) {

		// ...When GetUsers is called, we don't know what should happen yet.
		//
		when("GetUsers is called")

		// ...When GetStatus is called...
		//
		when("GetStatus is called", func(it It) {

			// ...It Should return an invalid status code
			it("should return an invalid status code")

			// ...It Should return an error message
			it("should return an error message")

			// ...It should return an 200 http status code
			it("should return an 200 http status code")

		})

	})

	/* Outputs:

	Feature: Specing A New Feature
	  Given a valid Api

	  Given an invalid Api
	    When GetUsers is called
	    When GetStatus is called
	    » It should return an invalid status code «-- NOT IMPLEMENTED
	    » It should return an error message «-- NOT IMPLEMENTED
	    » It should return an 200 http status code «-- NOT IMPLEMENTED
	*/
}

func Test_Application_Startup(t *testing.T) {

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
			it("should ping the master service")
			it("should ask the master for work to do")
		})
	})

	Given(t, "a ping request", func(when When) {
		when("the response is alive", func(it It) {
			it("should ask the master for work to do")
		})

		when("an error is returned", func(it It) {
			it("should panic")
			it("should exit")
		})
	})

	Given(t, "a request for work to do", func(when When) {
		when("the response has work to do", func(it It) {
			it("should create a NewClient()")
			it("should start the monitor's Ticker")
			it("should start the client's process")
		})

		when("no work is returned", func(it It) {
			it("should sleep for the defined polling interval")
		})

		when("an error is returned", func(it It) {
			it("should sleep for the defined polling interval")
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
