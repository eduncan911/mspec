/*
Package gomspec is a testing package for the Go framework that extends Go's built-in testing package.  It is modeled after the BDD Feature Specification story workflow, such as:

  With Feature X
    Given a context
    When an event occurs
    Then it should do something

Currently it has an included `Expectation` struct that mimics basic assertion behaviors.  Future plans may allow for custom assertion packages (like testify).

Referencing / Importing

  import "github.com/eduncan911/gomspec"

Building

There is no executable.  Instead, this is a package to include in your
favorite projects for test.  If you wish to build for sanity checks (the
Go compiler will only check that a package builds if it does not have a
main() func), build it like any other project:

  go build

Testing

There is nothing like using a testing package to test itself.  There is
some nice rich information avaliable.

  go test

So, Why another BDD Framework?

When evaluating several BDD frameworks, [Pranavraja's Zen](https://github.com/pranavraja/zen)
package for Go came close; but, it was lacking the more "story" overview.
 There is something to be said for simple testing in Go (and simple coding);
 therefore, if you are the type to keep it simple, then you may want to use
 his framework as it is just the context and specs.

I forked his code and submitted a few bug tweaks.  But along the way, I
started to have grand visions of my soul mate [Machine.Specifications](https://github.com/machine/machine.specifications)
(aka MSpec for short) for BDD testing.  The ease of defining complete
stories right down to the scenarios without having to implement them
intrigued me in Go.  Could I bring those practices to Go, using a
bare-bones framework?  Ok, done.  And since it was so heavily inspired
by Aaron's MSpec project, I kept the name going here: `GoMspec`.

While keeping backwards compatibility with his existing Zen framework,
I defined several goals for this package:

* Had to stay simple with Give/When/Then definitions.  No complex coding.
* Keep the low syntax noise from the existing Zen package.
* I had to be able to code specs without implementation details.

No Implementation Details needed

That last goal above is key.  I tend to design my C# code using Machine.Specifications in true BDD-style by writing entire stories and grand specs.
In C# land, it's not unheard of me hitting 50 to 100 specs across a
single feature and a few different contexts in an hour or two, before
writing any code.  So with this framework, I came up with a simple method
name, `NA()`, to keep the syntax noise down.

Therefore, you are free to code specs with just a little syntax noise:

	// defining specs in Go, without imeplementing or stubbing code
	it("should do this", NA())
	it("should do that", NA())
	it("should not be red", NA())
	it("should not be from the year 8,000 BC", NA())

*/
package gomspec
