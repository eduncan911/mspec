
# mspec
    import "github.com/eduncan911/gomspec"

Package mspec is a BDD Feature Specifications testing package for Go(Lang) with a strong emphases on spec'ing your feature(s) and scenarios first before any code is written.  This leaves you free to think of your project and features as a whole without the distraction of writing code.

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

### Testing
There is nothing like using a testing package to test itself.  There is
some nice rich information avaliable.


	go test

## Examples

Be sure to check out more examples in the examples/ folder.


	$ cd examples/
	/examples$ go test

Or just open the files and take a look.  That's the most important part anyways.

### Why Another BDD Framework
When evaluating several BDD frameworks, [Pranavraja's Zen](a href="https://github.com/pranavraja/zen">https://github.com/pranavraja/zen</a>) package for Go came close - really close; but, it was lacking the more "story" overview I've been accustomed to over the years with [Machine.Specifications](a href="https://github.com/machine/machine.specifications">https://github.com/machine/machine.specifications</a>) in C# (.NET land).

Do note that there is something to be said for simple testing in Go (and simple coding); therefore, if you are the type to keep it short and sweet and just code, then you may want to use Pranavraja's framework as it is just the context (Desc) and specs writing.

I forked his code and submitted a few bug tweaks at first.  But along the way, I started to have grand visions of my soul mate [Machine.Specifications](a href="https://github.com/machine/machine.specifications">https://github.com/machine/machine.specifications</a>) (which is called MSpec for short) for BDD testing.  The ease of defining complete stories right down to the scenarios without having to implement them intrigued me in C#.  It freed me from worrying about implementation details and just focus on the feature I was writing: What did it need to do?  What context was I given to start with? What should it do?

So while using Pranavraja's Zen framework, I kept asking myself: Could I bring those MSpec practices to Go, using a bare-bones framework?  Ok, done.  And since it was so heavily inspired by Aaron's MSpec project, I kept the name going here: `GoMSpec`.

While keeping backwards compatibility with his existing Zen framework, I defined several goals for this package:

* Had to stay simple with Give/When/Then definitions.  No complex coding.
* Keep the low syntax noise from the existing Zen package.
* I had to be able to write features, scenarios and specs with no implementation details needed.

### No Implementation Details needed
That last goal above is key and I think is what speaks truly about what BDD is: focus on the story, feature and/or context you are designing - focus on the Behavior!  I tended to design my C# code using Machine.Specifications in this BDD-style by writing entire stories and grand specs up front - designing the system I was building, or the feature I was extending.  In C# land, it's not unheard of me hitting 50 to 100 specs across a single feature and a few different contexts in an hour or two, before writing any code.  Which at that point, I had everything planned out pretty much the way it should behave.

So with this framework, I came up with a simple method name, `NA()`, to keep the syntax noise down.

Therefore, you are free to code specs with just a little syntax noise:


	// defining specs in Go, without imeplementing or stubbing code
	it("should do this", NA())
	it("should do that", NA())
	it("should not be red", NA())
	it("should not be from the year 8,000 BC", NA())






## func AssertionsFn
``` go
func AssertionsFn(fn func(s *Specification) Assert)
```
AssertionsFn will assign the assertions used for all tests.
The specified struct must implement the mspec.Assert interface.


	   mspec.AssertionsFn(func(s *Specification) Assert {
		    return &MyCustomAssertions{}
	   })


## func Given
``` go
func Given(t *testing.T, given string, when ...func(When))
```
Given defines the Feature's specific context to be spec'd out.


## func ResetConfig
``` go
func ResetConfig()
```
ResetConfig will reset all options back to their default configuration.
Useful for custom colors in the middle of a specification.


## func Setup
``` go
func Setup(before, after func()) func(fn func(Assert)) func(Assert)
```
Setup is used to define before/after (setup/teardown) functions.



## type Assert
``` go
type Assert interface {

    // Implements asserts that an object is implemented by the specified interface.
    //
    //    assert.Implements((*MyInterface)(nil), new(MyObject), "MyObject")
    Implements(interfaceObject interface{}, object interface{}, msgAndArgs ...interface{}) bool

    // IsType asserts that the specified objects are of the same type.
    IsType(expectedType interface{}, object interface{}, msgAndArgs ...interface{}) bool

    // Equal asserts that two objects are equal.
    //
    //    assert.Equal(123, 123, "123 and 123 should be equal")
    //
    // Returns whether the assertion was successful (true) or not (false).
    Equal(expected, actual interface{}, msgAndArgs ...interface{}) bool

    // Exactly asserts that two objects are equal is value and type.
    //
    //    assert.Exactly(int32(123), int64(123), "123 and 123 should NOT be equal")
    //
    // Returns whether the assertion was successful (true) or not (false).
    Exactly(expected, actual interface{}, msgAndArgs ...interface{}) bool

    // NotNil asserts that the specified object is not nil.
    //
    //    assert.NotNil(err, "err should be something")
    //
    // Returns whether the assertion was successful (true) or not (false).
    NotNil(object interface{}, msgAndArgs ...interface{}) bool

    // Nil asserts that the specified object is nil.
    //
    //    assert.Nil(err, "err should be nothing")
    //
    // Returns whether the assertion was successful (true) or not (false).
    Nil(object interface{}, msgAndArgs ...interface{}) bool

    // Empty asserts that the specified object is empty.  I.e. nil, "", false, 0 or a
    // slice with len == 0.
    //
    // assert.Empty(obj)
    //
    // Returns whether the assertion was successful (true) or not (false).
    Empty(object interface{}, msgAndArgs ...interface{}) bool

    // slice with len == 0.
    //
    // if assert.NotEmpty(obj) {
    //   assert.Equal("two", obj[1])
    // }
    //
    // Returns whether the assertion was successful (true) or not (false).
    NotEmpty(object interface{}, msgAndArgs ...interface{}) bool

    // Len asserts that the specified object has specific length.
    // Len also fails if the object has a type that len() not accept.
    //
    //    assert.Len(mySlice, 3, "The size of slice is not 3")
    //
    // Returns whether the assertion was successful (true) or not (false).
    Len(object interface{}, length int, msgAndArgs ...interface{}) bool

    // True asserts that the specified value is true.
    //
    //    assert.True(myBool, "myBool should be true")
    //
    // Returns whether the assertion was successful (true) or not (false).
    True(value bool, msgAndArgs ...interface{}) bool

    // False asserts that the specified value is true.
    //
    //    assert.False(myBool, "myBool should be false")
    //
    // Returns whether the assertion was successful (true) or not (false).
    False(value bool, msgAndArgs ...interface{}) bool

    // NotEqual asserts that the specified values are NOT equal.
    //
    //    assert.NotEqual(obj1, obj2, "two objects shouldn't be equal")
    //
    // Returns whether the assertion was successful (true) or not (false).
    NotEqual(expected, actual interface{}, msgAndArgs ...interface{}) bool

    // Contains asserts that the specified string contains the specified substring.
    //
    //    assert.Contains("Hello World", "World", "But 'Hello World' does contain 'World'")
    //
    // Returns whether the assertion was successful (true) or not (false).
    Contains(s, contains string, msgAndArgs ...interface{}) bool

    // NotContains asserts that the specified string does NOT contain the specified substring.
    //
    //    assert.NotContains("Hello World", "Earth", "But 'Hello World' does NOT contain 'Earth'")
    //
    // Returns whether the assertion was successful (true) or not (false).
    NotContains(s, contains string, msgAndArgs ...interface{}) bool

    // WithinDuration asserts that the two times are within duration delta of each other.
    //
    //   assert.WithinDuration(time.Now(), time.Now(), 10*time.Second, "The difference should not be more than 10s")
    //
    // Returns whether the assertion was successful (true) or not (false).
    WithinDuration(expected, actual time.Time, delta time.Duration, msgAndArgs ...interface{}) bool

    // NoError asserts that a function returned no error (i.e. `nil`).
    //
    //   actualObj, err := SomeFunction()
    //   if assert.NoError(err) {
    //	   assert.Equal(actualObj, expectedObj)
    //   }
    //
    // Returns whether the assertion was successful (true) or not (false).
    NoError(theError error, msgAndArgs ...interface{}) bool

    // Error asserts that a function returned an error (i.e. not `nil`).
    //
    //   actualObj, err := SomeFunction()
    //   if assert.Error(err, "An error was expected") {
    //	   assert.Equal(err, expectedError)
    //   }
    //
    // Returns whether the assertion was successful (true) or not (false).
    Error(theError error, msgAndArgs ...interface{}) bool

    // EqualError asserts that a function returned an error (i.e. not `nil`)
    // and that it is equal to the provided error.
    //
    //   actualObj, err := SomeFunction()
    //   if assert.Error(err, "An error was expected") {
    //	   assert.Equal(err, expectedError)
    //   }
    //
    // Returns whether the assertion was successful (true) or not (false).
    EqualError(theError error, errString string, msgAndArgs ...interface{}) bool
}
```
Assert is an interface used by each Specification that can used to
enforce a rule.

Custom assertion packages may be defined by calling mspec.AssertionsFn(...).
An internal assert/forward_assertions.go currently implements the default instance at runtime.











## type It
``` go
type It func(title string, assert ...func(Assert))
```
It defines the specification of When something happens.











## type MSpecConfig
``` go
type MSpecConfig struct {
    AnsiOfFeature            string
    AnsiOfGiven              string
    AnsiOfWhen               string
    AnsiOfThen               string
    AnsiOfThenNotImplemented string
    AnsiOfThenWithError      string
    AnsiOfCode               string
    AnsiOfCodeError          string
    AnsiOfExpectedError      string
    // contains filtered or unexported fields
}
```
MSpecConfig defines the configuration used by the package.











## type Specification
``` go
type Specification struct {
    T                       *testing.T
    Feature                 string
    Given                   string
    When                    string
    Spec                    string
    AssertFn                func(Assert)
    AssertionFailed         bool
    AssertionFailedMessages []string
    // contains filtered or unexported fields
}
```
Specification holds the state of the context for a specific specification.











### func (\*Specification) PrintContext
``` go
func (spec *Specification) PrintContext()
```


### func (\*Specification) PrintError
``` go
func (spec *Specification) PrintError(message string)
```


### func (\*Specification) PrintFeature
``` go
func (spec *Specification) PrintFeature()
```


### func (\*Specification) PrintSpec
``` go
func (spec *Specification) PrintSpec()
```


### func (\*Specification) PrintSpecNotImplemented
``` go
func (spec *Specification) PrintSpecNotImplemented()
```


### func (\*Specification) PrintSpecWithError
``` go
func (spec *Specification) PrintSpecWithError()
```


### func (\*Specification) PrintWhen
``` go
func (spec *Specification) PrintWhen()
```


## type When
``` go
type When func(when string, it ...func(It))
```
When defines the action or event when Given a specific context.

















- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)