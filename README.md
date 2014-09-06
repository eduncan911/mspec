2014/09/06 - I've completed a large refactor; but, I haven't updated this README / documentation updates yet.  I've released this to `master` for dog-fooding to any quirks that may show up.

## GoMSpec

A BDD Feature Specifications testing package for Go(Lang) with a strong emphases on spec'ing your feature(s) and scenarios before any code is written.  This leaves you free to think of your project and features as a whole without the distraction of writing code.

Source Documentation available at [![GoDoc](https://godoc.org/github.com/eduncan911/gomspec?status.svg)](https://godoc.org/github.com/eduncan911/gomspec)

`GoMSpec` is a testing package for the Go framework that extends Go's built-in testing package.  It is modeled after the BDD Feature Specification story workflow such as:

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

`GoMSpec` is configured by default to output all stories to the console for easy visibility.

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
~/go/src/github.com/eduncan911/gomspec$ cd examples/
~/go/src/github.com/eduncan911/gomspec/examples$ go test
```

Or just open the files and take a look.  That's the most important part anyways.

## Why another BDD Framework?

When evaluating several BDD frameworks, [Pranavraja's Zen](https://github.com/pranavraja/zen) package for Go came close - really close; but, it was lacking the more "story" overview I've been accustomed to over the years with [Machine.Specifications](https://github.com/machine/machine.specifications) in C# (.NET land).  

Do note that there is something to be said for simple testing in Go (and simple coding); therefore, if you are the type to keep it short and sweet and just code, then you may want to use Pranavraja's framework as it is just the context (Desc) and specs writing.

I forked his code and submitted a few bug tweaks at first.  But along the way, I started to have grand visions of my soul mate [Machine.Specifications](https://github.com/machine/machine.specifications) (which is called MSpec for short) for BDD testing.  The ease of defining complete stories right down to the scenarios without having to implement them intrigued me in C#.  It freed me from worrying about implementation details and just focus on the feature I was writing: What did it need to do?  What context was I given to start with? What should it do?

So while using Pranavraja's Zen framework, I kept asking myself: Could I bring those MSpec practices to Go, using a bare-bones framework?  Ok, done.  And since it was so heavily inspired by Aaron's MSpec project, I kept the name going here: `GoMSpec`.

While keeping backwards compatibility with his existing Zen framework, I defined several goals for this package:

* Had to stay simple with Give/When/Then definitions.  No complex coding.
* Keep the low syntax noise from the existing Zen package.
* I had to be able to write features, scenarios and specs with no implementation details needed.

### No Implementation Details Needed

That last goal above is key and I think is what speaks truly about what BDD is: focus on the story, feature and/or context you are designing - focus on the Behavior!  I tended to design my C# code using Machine.Specifications in this BDD-style by writing entire stories and grand specs up front - designing the system I was building, or the feature I was extending.  In C# land, it's not unheard of me hitting 50 to 100 specs across a single feature and a few different contexts in an hour or two, before writing any code.  Which at that point, I had everything planned out pretty much the way it should behave.  

So with this framework, I came up with:

```go
// defining specs in Go, without implementing or stubbing code
//
package examples

import (
    . "github.com/eduncan911/gomspec"
    "testing"
)

func Test_Specing_A_New_Feature(t *testing.T) {

    // you can quickly spec new features with little syntax noise
    //

    // GIVEN a valid Api, what shall we do?  not sure yet.
    //
    Given(t, "a valid Api")

    // GIVEN an invalid Api...
    //
    Given(t, "an invalid Api", func(when When) {

        // ...WHEN GetUsers is called, we don't know what SHOULD happen yet.
        //
        when("GetUsers is called")

        // ...WHEN GetStatus is called...
        //
        when("GetStatus is called", func(it It) {

            // ...IT SHOULD return an invalid status code
            it("should return an invalid status code")

            // ...IT SHOULD return an error message
            it("should return an error message")

        })
    })
}
```

Note that `Given`, `when` and `it` all have optional variadic parameters.  THis allows you to spec things out as far as you want.

Outputs:

```
    Feature: Specing A New Feature
      Given a valid Api

      Given an invalid Api
        When GetUsers is called
        When GetStatus is called
        » It should return an invalid status code «-- NOT IMPLEMENTED
        » It should return an error message «-- NOT IMPLEMENTED
        » It should return an 200 http status code «-- NOT IMPLEMENTED

```



# Roadmap

* write blog post
* write wiki
* more examples as well as custom formatters/expectations
* `Setup()` examples
* Total tests passed, errored, skipped
* HTML output
* surpressing output (quiet)
* concurrent channel execution of `it`s

NOTE: If you are looking for the Zen version that remains compatible with Pranavraja's Zen](https://github.com/pranavraja/zen) version, you will want to refer to the [specific tag v0.1](https://github.com/eduncan911/gomspec/tree/v0.1).
