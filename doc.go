/*Package mspec is a BDD context/specification testing package for Go(Lang) with a strong emphases on spec'ing your feature(s) and scenarios first, before any code is written using as little syntax noise as possible.  This leaves you free to think of your project and features as a whole without the distraction of writing any code with the added benefit of having tests ready for your project.

[![GoDoc](https://godoc.org/github.com/eduncan911/mspec?status.svg)](https://godoc.org/github.com/eduncan911/mspec) holds the source documentation (where else?)

Features

* Uses natural language (Given/When/Then)
* Stubbing
* Human-readable outputs
* HTML output (coming soon)
* Use custom Assertions
* Configuration options
* Uses Testify's rich assertions
* Uses Go's built-in testing.T package

Go Get It

Install it with one line of code:

	go get github.com/eduncan911/gomspec

There are no external dependencies and it is built against Go's internal packages.  The only dependency is that you have [GOPATH setup normaly](https://golang.org/doc/code.html).

Go Stub Something

Create a new file to hold your specs.

	pico sample_test.go

Using Dan North's original BDD definitions, you spec code using the Given/When/Then storyline similar to:

	Feature X
		Given a context
		When an event occurs
		Then it should do something

But this is just a static example.  Let's take a real example from one of my projects:

	Feature Public API Contracts
		Given an invalid Api
		When GetStatus is called
		Then it should return an invalid status code
		Then it should return an error message
		Then it should return an 200 htp status code.

You represent these thoughts in code like this:

	package main

	import (
	    . "github.com/eduncan911/gomspec"
	    "testing"
	)

	func Test_Public_API_Contract(t *testing.T) {

	    Given(t, "an invalid Api", func(when When) {

	        when("GetStatus is called", func(it It) {

	            it("should return an invalid status code")

	            it("should return an error message")

	            it("should return an 200 http status code")

	        })

			// we can start designing our package immediately by specifying more contexts
			// and specs.  there is no fuddling with real code that gets us off topic.
	        when("GetUsers is called")
	    })
	}

Note that `Given`, `when` and `it` all have optional variadic parameters.  This allows you to spec things out with as little or as far as you want.

That's it.  Now run it:

    $ go test
    Feature: API Contract
      Given an invalid Api
        When GetStatus is called
        » It should return an invalid status code «-- NOT IMPLEMENTED
        » It should return an error message «-- NOT IMPLEMENTED
        » It should return an 200 http status code «-- NOT IMPLEMENTED

        When GetUsers is called

	PASS
	ok  	github.com/eduncan911/gomspec	0.001s

Print it out and stick it on your office door for everyone to see what you are working on.

This is actually colored output in Terminal:



It is not uncommon to go back and tweak your stories over time as you talk with your domain experts, modifying exactly the scenarios and specifications that should happen.
























`GoMSpec` is a testing package for the Go framework that extends Go's built-in testing package.  It is modeled after the BDD Feature Specification story workflow such as:

  With Feature X
    Given a context
    When an event occurs
    Then it should do something

Currently it has an included `Expectation` struct that mimics basic assertion behaviors.  Future plans may allow for custom assertion packages (like testify).

Getting it

  go get github.com/eduncan911/gomspec

Importing it

  import . "github.com/eduncan911/gomspec"

Writing Specs

	// dogs_test.go
	//
	package dogs

	import (
	    . "github.com/eduncan911/gomspec"
	    "testing"
	)

	func Test_Washing_Dogs(t *testing.T) {

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
	}

Testing it

  go test

Which outputs the following:

	Feature: Washing Dogs

	  Given a dog that has been painted red
	  and the paint is washable
	  and no one has washed the dog yet

	    When the dog is washed
	    » It should have the paint come off
	    » It should be a normal color
	    » It should smell like a clean dog

	PASS
	ok  	github.com/eduncan911/gomspec/examples	0.007s

Nice eh?

Testing

There is nothing like using a testing package to test itself.  There is
some nice rich information avaliable.

  go test

## Examples

Be sure to check out more examples in the examples/ folder.

  $ cd examples/
  /examples$ go test

Or just open the files and take a look.  That's the most important part anyways.

Why Another BDD Framework

When evaluating several BDD frameworks, [Pranavraja's Zen](https://github.com/pranavraja/zen) package for Go came close - really close; but, it was lacking the more "story" overview I've been accustomed to over the years with [Machine.Specifications](https://github.com/machine/machine.specifications) in C# (.NET land).

Do note that there is something to be said for simple testing in Go (and simple coding); therefore, if you are the type to keep it short and sweet and just code, then you may want to use Pranavraja's framework as it is just the context (Desc) and specs writing.

I forked his code and submitted a few bug tweaks at first.  But along the way, I started to have grand visions of my soul mate [Machine.Specifications](https://github.com/machine/machine.specifications) (which is called MSpec for short) for BDD testing.  The ease of defining complete stories right down to the scenarios without having to implement them intrigued me in C#.  It freed me from worrying about implementation details and just focus on the feature I was writing: What did it need to do?  What context was I given to start with? What should it do?

So while using Pranavraja's Zen framework, I kept asking myself: Could I bring those MSpec practices to Go, using a bare-bones framework?  Ok, done.  And since it was so heavily inspired by Aaron's MSpec project, I kept the name going here: `GoMSpec`.

While keeping backwards compatibility with his existing Zen framework, I defined several goals for this package:

* Had to stay simple with Give/When/Then definitions.  No complex coding.
* Keep the low syntax noise from the existing Zen package.
* I had to be able to write features, scenarios and specs with no implementation details needed.

No Implementation Details needed

That last goal above is key and I think is what speaks truly about what BDD is: focus on the story, feature and/or context you are designing - focus on the Behavior!  I tended to design my C# code using Machine.Specifications in this BDD-style by writing entire stories and grand specs up front - designing the system I was building, or the feature I was extending.  In C# land, it's not unheard of me hitting 50 to 100 specs across a single feature and a few different contexts in an hour or two, before writing any code.  Which at that point, I had everything planned out pretty much the way it should behave.

So with this framework, I came up with a simple method name, `NA()`, to keep the syntax noise down.

Therefore, you are free to code specs with just a little syntax noise:

	// defining specs in Go, without imeplementing or stubbing code
	it("should do this", NA())
	it("should do that", NA())
	it("should not be red", NA())
	it("should not be from the year 8,000 BC", NA())

*/
package mspec
