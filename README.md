# MSpec - Thou Shalt Spec Features

`MSpec` is a BDD context/specification testing package for Go(Lang) with a strong emphases on spec'ing your feature(s) and scenarios first, before any code is written using as little syntax noise as possible.  This leaves you free to think of your project and features as a whole without the distraction of writing any code with the added benefit of having tests ready for your project.

[![GoDoc](https://godoc.org/github.com/eduncan911/go-mspec?status.svg)](https://godoc.org/github.com/eduncan911/go-mspec) [![Build Status](https://travis-ci.org/eduncan911/go-mspec.svg?branch=master)](https://travis-ci.org/eduncan911/go-mspec) [![Go Report Card](https://goreportcard.com/badge/github.com/eduncan911/go-mspec)](https://goreportcard.com/report/github.com/eduncan911/go-mspec)

## Features

* Uses natural language (Given/When/Then)
* Stubbing (write specs with no code)
* Human-readable outputs
* HTML output, e.g. for C.I servers (coming soon...)
* Override and use your own custom Assertions
* Configuration options
* Uses Testify's rich assertions by default
* Uses Go's built-in testing.T package (no dependencies)

## API Specification

2016-02-20

The API as of tag v0.4 has been finalized; though, it could change depending on feedback or bug fixes.

There are additional features to add that will only expand the API.  

We do not expect any more breaking changes as of v0.4.

# Go Get It

Install it with one line of code:

`go get -v -u github.com/eduncan911/go-mspec`

There are no external dependencies and it is built against Go's internal packages.  The only dependency is that you have [GOPATH setup normaly](https://golang.org/doc/code.html).

# Go Use It

## Stubbing a new Feature

Using [Dan North's original BDD definitions](http://dannorth.net/introducing-bdd/), you spec code using the Given/When/Then storyline similar to:

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

Note that `Given`, `when` and `it` all have optional variadic parameters.  This allows you to spec things out as little or as far as you want.  

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


## Implement a Specification

Let's write a full specification with real code.

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

Now you can run the tests using Go's built-in testing framework.  

`$ go test`

This outputs:

```
$ go test
Feature: Washing Dogs

  Given a dog that has been painted red
  and the paint is washable
  and no one has washed the dog yet
    When the dog is washed
    » It should have the paint come off
    » It should be a normal color
    » It should smell like a clean dog
```

The output specifies the feature and then the scenario you are testing.  

There are multiple output settings that can be configured. `MSpec` is 
configured by default to output stdout for easy visibility.  An HTML runner will be 
included (shortly); or, you can implement your own custom output (e.g. json post to 
a C.I. build server).

# Errors are well defined

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

This outputs:

![mspec error example](http://i.imgur.com/iuVlElc.png)

The error message outputs:

* What should have happened.
* The test file and line number that failed.
* A snippet of code around that line number.

The default coloring also makes it standout amongst other tests that passed.

## More Examples

Be sure to check out more advanced examples in the examples/ folder including how to spec code without writing any implementation details.

```bash
$ cd $GOPATH/src/github.com/eduncan911/go-mspec/examples/
$ go test
```

Or just open the files and take a look.  That's the most important part anyways.

# Why another BDD Framework?

When evaluating several BDD frameworks, [Pranavraja's Zen](https://github.com/pranavraja/zen) package for Go came close - really close; but, it was lacking the more "story" overview I've been accustomed to over the years with [Machine.Specifications](https://github.com/machine/machine.specifications) in C# (.NET land).  

Do note that there is something to be said for simple testing in Go (and simple coding); therefore, if you are the type to keep it short and sweet and just code, then you may want to use Pranavraja's framework as it is just the context (Desc) and specs writing.

I forked his code and submitted a few bug tweaks at first.  But along the way, I started to have grand visions of my soul mate [Machine.Specifications](https://github.com/machine/machine.specifications) (which is called MSpec for short) for BDD testing.  The ease of defining complete stories right down to the scenarios without having to implement them intrigued me in C#.  It freed me from worrying about implementation details and just focus on the feature I was writing: What did it need to do?  What context was I given to start with? What should it do?

So while using Pranavraja's Zen framework, I kept asking myself: Could I bring those MSpec practices to Go, using a bare-bones framework?  Ok, done.  And since it was so heavily inspired by Aaron's MSpec project, I kept the name going here: `MSpec`.

# Roadmap

* write blog post
* more examples as well as custom formatters/expectations
* `SetConfig()` examples
* Total tests passed, errored, skipped
* HTML output
* surpressing output (quiet)
* concurrent channel execution of `it`s
* custom outputs

