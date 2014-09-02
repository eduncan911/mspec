# GoMspec

A BDD Feature Specifications testing package for Go(Lang)

Source Documentation available at [![GoDoc](https://godoc.org/github.com/eduncan911/gomspec?status.svg)](https://godoc.org/github.com/eduncan911/gomspec)

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
```

### Go Test It

You run the tests using Go's built-in testing framework.  

`GoMspec` is configured by default to output all stories to the console for easy visibility.

`$ go test`

Outputs:

```
Feature: Washing Dogs

  Given a dog that has been painted red
  and the paint is washable
  and no one has washed the dog yet

    When the dog is washed
    » It should have the paint come off
    » It should be a normal color
    » It should smell like a clean dog
```

The output specifies the feature and then the scenario you are testing.  There are multiple output settings that can be configured.

## Errors are well defined

Let's add a feature that has a spec that will blow up.

```go
package main

import (
    . "github.com/eduncan911/gomspec"
    "testing"
)

func Test_Creating_a_Client(t *testing.T) {

    Given(t, "a valid ProviderConfig", func(when When) {

        pc := &ProviderConfig{
            Name:          "Acme Corp",
            ShellScript:   "acmecorp-import.sh",
            RunValidation: true,
            RunMatching:   true,
            RunUpserts:    true,
        }

        when("calling NewClient() constructor", func(it It) {

            c, err := NewClient(*pc)

            it("should not return an error.", func(expect Expect) {
                expect(err).ToNotExist()
            })

            it("should return a valid client object.", func(expect Expect) {
                expect(c).ToExist()
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

(insert error pic with color output)

## Examples

Be sure to check out more examples in the examples/ folder.

```bash
$ cd examples/
$ go test
```

## Why another BDD Framework?

When evaluating several BDD frameworks, [Pranavraja's Zen](https://github.com/pranavraja/zen) package for Go came close; but, it was lacking the more "story" overview.  There is something to be said for simple testing in Go (and simple coding); therefore, if you are the type to keep it simple, then you may want to use his framework as it is just the context and specs.

I forked his code and submitted a few bug tweaks.  But along the way, I started to have grand visions of my soul mate [Machine.Specifications](https://github.com/machine/machine.specifications) (aka MSpec for short) for BDD testing.  The ease of defining complete stories right down to the scenarios without having to implement them intrigued me in Go.  Could I bring those practices to Go, using a bare-bones framework?  Ok, done.  And since it was so heavily inspired by Aaron's MSpec project, I kept the name going here: `GoMspec`.

While keeping backwards compatibility with his existing Zen framework, I defined several goals for this package:

* Had to stay simple with Give/When/Then definitions.  No complex coding.
* Keep the low syntax noise from the existing Zen package.
* I had to be able to code specs without implementation details.

### No Implementation Details needed

That last goal above is key.  I tend to design my C# code using Machine.Specifications in true BDD-style by writing entire stories and grand specs.  In C# land, it's not unheard of me hitting 50 to 100 specs across a single feature and a few different contexts in an hour or two, before writing any code.  So with this framework, I came up with a simple method name, `NA()`, to keep the syntax noise down.  

Therefore, you are free to code specs with just a little syntax noise:

```go
// defining specs in Go, without imeplementing or stubbing code
it("should do this", NA())
it("should do that", NA())
it("should not be red", NA())
it("should not be from the year 8,000 BC", NA())
```

I can live with that.  I think it is on par with Machine.Specifications' `Not Implemented` details:

```c#
// this is how you do it in C# with MSpec
It should_do_this;
It should_do_that;
it should_not_be_red;
it should_not_be_from_the_year_8000_BC;
```

Those underscores always bugged me.  So it's a trade off to have free-form quoted text verses defining a delegate with underscores that I always fat-fingering.

# Roadmap

* write blog post
* write wiki
* more examples as well as custom formatters/expectations
* `Setup()` examples
* Total tests passed, errored, skipped
* HTML output
* surpressing output (quiet)
* concurrent channel execution of `it`s

