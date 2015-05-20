# MSpec - Thou Shall Spec Features

`MSpec` is a BDD context/specification testing package for Go(Lang) with a strong emphases on spec'ing your feature(s) and scenarios first, before any code is written using as little syntax noise as possible.  This leaves you free to think of your project and features as a whole without the distraction of writing any code with the added benefit of having tests ready for your project.

[![GoDoc](https://godoc.org/github.com/eduncan911/go-mspec?status.svg)](https://godoc.org/github.com/eduncan911/go-mspec) holds 
the source documentation (where else?)

Features

* Uses natural language (Given/When/Then)
* Stubbing
* Human-readable outputs
* HTML output (coming soon)
* Use custom Assertions
* Configuration options
* Uses Testify's rich assertions
* Uses Go's built-in testing.T package

# Go Get It

Install it with one line of code:

`go get github.com/eduncan911/go-mspec`

There are no external dependencies and it is built against Go's internal packages.  The only dependency is that you have [GOPATH setup normaly](https://golang.org/doc/code.html).

# Go Stub a New Feature

Using Dan North's original BDD definitions, you spec code using the Given/When/Then storyline similar to:

```
Feature X
    Given a context
    When an event occurs
    Then it should do something
```

You represent these thoughts in your tests like this:

```go
// api_test.go
package main

import (
    . "github.com/eduncan911/go-mspec"
    "testing"
)

func Test_API_Contract(t *testing.T) {

    Given(t, "a valid Api")

    Given(t, "an invalid Api", func(when When) {

        when("GetStatus is called", func(it It) {

            it("should return an invalid status code")

            it("should return an error message")

            it("should return an 200 http status code")

        })

        when("GetUsers is called")
    })
}
```

Note that `Given`, `when` and `it` all have optional variadic parameters.  This allows you to spec things out with as little or as far as you want.  

This compiles and allows you to `go test` it immediately:

```bash
    $ go test
      Feature: API Contract
        Given a valid Api
         
        Given an invalid Api
          When GetStatus is called
          » It should return an invalid status code «-- NOT IMPLEMENTED
          » It should return an error message «-- NOT IMPLEMENTED
          » It should return an 200 http status code «-- NOT IMPLEMENTED
            
          When GetUsers is called

```

It is not uncommon to go back and tweak your stories over time as you talk with your domain experts, modifying exactly the scenarios and specifications that should happen.

Print it out and stick it on your office door for everyone to see what you are working on.





`MSpec` is a testing package for the Go framework that extends Go's built-in testing package.  It is modeled after the BDD Feature Specification story workflow such as:

```
Feature X
    Given a context
    When an event occurs
    Then it should do something
```



### Go Spec It

Pay attention to the function name as it is used as part of the output.

```go
// dogs_test.go
//
package dogs

import (
    . "github.com/eduncan911/go-mspec"
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

            it("should have the paint come off", func(assert Assert) {
                assert.Nil(d.paint)
            })

            it("should be a normal color", func(assert Assert) {
                assert.Equal(d.color, normalColor)
            })

            it("should smell like a clean dog", func(assert Assert) {
                assert.True(d.washed)
            })
        })
    })
}
```

### Go Test It

You run the tests using Go's built-in testing framework.  

`MSpec` is configured by default to output all stories to the console for easy visibility.

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

The output specifies the feature and then the scenario you are testing.  There are multiple output settings that can be configured as well.

## Errors are well defined

Let's add a feature that has a spec that will blow up.

```go
package main

import (
    . "github.com/eduncan911/go-mspec"
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

            it("should not return an error.", func(assert Assert) {
                assert.NoError(err)
            })

            it("should return a valid client object.", func(assert Assert) {
                assert.NotNil(c)
            })

            // this will blow up!
            //
            it("should generate a big error", func(assert Assert) {
                assert.True(false)
            })
        })
    })
}
```

Outputs:

(insert error pic with color output)

## Examples

Be sure to check out more advanced examples in the examples/ folder including how to spec code without writing any implementation details.

```bash
~/go/src/github.com/eduncan911/go-mspec$ cd examples/
~/go/src/github.com/eduncan911/go-mspec/examples$ go test
```

Or just open the files and take a look.  That's the most important part anyways.

## Why another BDD Framework?

When evaluating several BDD frameworks, [Pranavraja's Zen](https://github.com/pranavraja/zen) package for Go came close - really close; but, it was lacking the more "story" overview I've been accustomed to over the years with [Machine.Specifications](https://github.com/machine/machine.specifications) in C# (.NET land).  

Do note that there is something to be said for simple testing in Go (and simple coding); therefore, if you are the type to keep it short and sweet and just code, then you may want to use Pranavraja's framework as it is just the context (Desc) and specs writing.

I forked his code and submitted a few bug tweaks at first.  But along the way, I started to have grand visions of my soul mate [Machine.Specifications](https://github.com/machine/machine.specifications) (which is called MSpec for short) for BDD testing.  The ease of defining complete stories right down to the scenarios without having to implement them intrigued me in C#.  It freed me from worrying about implementation details and just focus on the feature I was writing: What did it need to do?  What context was I given to start with? What should it do?

So while using Pranavraja's Zen framework, I kept asking myself: Could I bring those MSpec practices to Go, using a bare-bones framework?  Ok, done.  And since it was so heavily inspired by Aaron's MSpec project, I kept the name going here: `MSpec`.

While keeping backwards compatibility with his existing Zen framework, I defined several goals for this package:

* Had to stay simple with Give/When/Then definitions.  No complex coding.
* Keep the low syntax noise from the existing Zen package.
* I had to be able to write features, scenarios and specs with no implementation details needed.

### No Implementation Details Needed

That last goal above is key and I think is what speaks truly about what BDD is: focus on the story, feature and/or context you are designing - focus on the Behavior!  I tended to design my C# code using Machine.Specifications in this BDD-style by writing entire stories and grand specs up front - designing the system I was building, or the feature I was extending.  In C# land, it's not unheard of me hitting 50 to 100 specs across a single feature and a few different contexts in an hour or two, before writing any code.  Which at that point, I had everything planned out pretty much the way it should behave.  

So with this framework, I came up with:




# Roadmap

* write blog post
* more examples as well as custom formatters/expectations
* `Setup()` examples
* Total tests passed, errored, skipped
* HTML output
* surpressing output (quiet)
* concurrent channel execution of `it`s

NOTE: If you are looking for the Zen version that remains compatible with Pranavraja's Zen](https://github.com/pranavraja/zen) 
version, you will want to refer to the [specific tag v0.1](https://github.com/eduncan911/go-mspec/tree/v0.1).
