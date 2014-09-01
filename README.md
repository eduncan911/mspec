# GoMspec

## BDD Feature Specifications testing for Go(Lang)

Blog Post(link)
Documentation / Wiki(link)

`GoMspec` is a testing package for the Go framework that extends Go's built-in testing package.  It is modeled after the BDD Feature Specification story workflow such as:

```
Feature X
    Given a context
    When an event occurs
    Then it should do something
```

Currently it has an included `Expectation` struct that mimics basic assertion behaviors.  Future plans may allow for custom assertion packages (like testify).

### Go Get It

`go get github.com/eduncan911/gomspec`

There are no external dependencies and is built against Go's internal packages: it's simple and lightweight.

### Go Spec It

Pay attention to the function name as it is used as part of the output.

```go
package main

import (
    . "github.com/eduncan911/gomspec"
    "testing"
)

func TestNewClient(t *testing.T) {

    Given(t, "a valid ProviderConfig", func(when When) {

        pc := &ProviderConfig{
            Name:          "Acme Corp",
            ShellScript:   "acmecorp-import.sh",
            RunValidation: true,
            RunMatching:   true,
            RunUpserts:    true,
        }

        when("calling NewClient", func(it It) {

            c, err := NewClient(*pc)

            it("should not return an error.", func(expect Expect) {
                expect(err).ToNotExist()
            })

            it("should return a valid client object.", func(expect Expect) {
                expect(c).ToExist()
            })

            it("should have cancelToken as false.", func(expect Expect) {
                expect(c.cancelToken).ToEqual(false)
            })

             // an example of more work to do
            it("should ask master for more work", NA())

            it("should have a valid pointer to a time.Ticker.", func(expect Expect) {
                expect(c.stateUpdateTicker).ToExist()
            })
        })
    })
}
```

### Go Test It

You run the tests using Go's built-in testing framework.  

`GoMspec` is configured by default to output all stories to the console for easy visibility.

`$ go test`

Outputs:

![Go BDD Test Output](http://i.imgur.com/MRJvVTc.png)

The output specifies the feature and then the scenario you are testing.  There are multiple output settings that can be configured.

## Errors are well defined

Let's add a 6th new spec that will blow up.

```go
package main

import (
    . "github.com/eduncan911/gomspec"
    "testing"
)

func TestNewClient(t *testing.T) {

    Given(t, "a valid ProviderConfig", func(when When) {

        pc := &ProviderConfig{
            Name:          "Acme Corp",
            ShellScript:   "acmecorp-import.sh",
            RunValidation: true,
            RunMatching:   true,
            RunUpserts:    true,
        }

        when("calling NewClient", func(it It) {

            c, err := NewClient(*pc)

            it("should not return an error.", func(expect Expect) {
                expect(err).ToNotExist()
            })

            it("should return a valid client object.", func(expect Expect) {
                expect(c).ToExist()
            })

            it("should have cancelToken as false.", func(expect Expect) {
                expect(c.cancelToken).ToEqual(false)
            })

            it("should ask master for more work", NA())

            it("should have a valid pointer to a time.Ticker.", func(expect Expect) {
                expect(c.stateUpdateTicker).ToExist()
            })

            // this will blow up!
            //
            it("should generate a big error", func(expect Expect) {
                expect(true).ToEqual(false)
            })
        })
    })
}
```

Outputs:

![Go BDD Tests Output](http://i.imgur.com/qshhxYp.png)

## Why another BDD Framework?

When evaluating several BDD frameworks, [Pranavraja's Zen](https://github.com/pranavraja/zen) package for Go came close; but, it was lacking the more "story" overview.  There is something to be said for simple testing in Go (and simple coding); therefore, if you are the type to keep it simple, then you may want to use his framework as it is just the context and specs.

I forked his code and submitted a few bug tweaks.  But along the way, I started to have grand visions of my soul mate [Machine.Specifications](https://github.com/machine/machine.specifications) for BDD testing.  The ease of defining complete stories right down to the scenarios without having to implement them intrigued me in Go.  Could I bring those practices to Go, using a bare-bones framework?  Ok, done.

While keeping backwards compatibility with his existing Zen framework, I defined several goals for this package:

* Had to stay simple with Give/When/Then definitions.  No complex coding.
* Keep the low syntax noise from the existing Zen package.
* I had to be able to code specs without implementation details.

### No Implementation Details needed

That last goal above is key.  I tend to design my C# code using Machine.Specifications in true BDD-style by writing entire stories and grand specs.  In C# land, it's not unheard of me hitting 50 to 100 specs across a single feature and a few different contexts in an hour or two, before writing any code.  So with this framework, I came up with a simple method name, `NA()`, to keep the syntax noise down.  

Therefore, you are free to code specs with just a little syntax noise:

```go
it("should do this", NA())
it("should do that", NA())
it("should not be red", NA())
it("should not be from the year 8,000 BC", NA())
```

I can live with that.  I think it is on par with Machine.Specifications not implemented details:

```c#
It should_do_this;
It should_do_that;
it should_not_be_red;
it should_not_be_from_the_year_8000_BC;
```

Those underscores always bugged me.  So it's a trade off to have free-form quoted text verses defining a delegate.

# IMCOMPLETE DOC: MORE TO COME

I just wanted to get this pushed tonight so I could get back to work.  This readme, as well as a blog post, is coming.

Todo

* write blog post
* write docs
* more examples as well as custom formatters/expectations
* `Setup()` examples
* `func Test_Name_With_Underscores(t *testing)` examples.
* Total tests passed, errored, skipped

Roadmap
* HTML output
* concurrent channel execution of `it`s

