package mspec

import "time"

// Assert is an interface used by each Specification that can used to
// enforce a rule.
//
// Custom assertion packages may be defined by calling mspec.AssertionsFn(...).
// An internal assert/forward_assertions.go currently implements the default instance at runtime.
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

	// TODO Implement Condition()
	// Uses a Comparison to assert a complex condition.
	///Condition(comp Comparison, msgAndArgs ...interface{}) bool

	// TODO Implement Panics()
	// Panics asserts that the code inside the specified PanicTestFunc panics.
	//
	//   assert.Panics(func(){
	//     GoCrazy()
	//   }, "Calling GoCrazy() should panic")
	//
	// Returns whether the assertion was successful (true) or not (false).
	//Panics(f PanicTestFunc, msgAndArgs ...interface{}) bool

	// TODO Implement NotPanics()
	// NotPanics asserts that the code inside the specified PanicTestFunc does NOT panic.
	//
	//   assert.NotPanics(func(){
	//     RemainCalm()
	//   }, "Calling RemainCalm() should NOT panic")
	//
	// Returns whether the assertion was successful (true) or not (false).
	//NotPanics(f PanicTestFunc, msgAndArgs ...interface{}) bool

	// WithinDuration asserts that the two times are within duration delta of each other.
	//
	//   assert.WithinDuration(time.Now(), time.Now(), 10*time.Second, "The difference should not be more than 10s")
	//
	// Returns whether the assertion was successful (true) or not (false).
	WithinDuration(expected, actual time.Time, delta time.Duration, msgAndArgs ...interface{}) bool

	// TODO Implement InDelta()
	// InDelta asserts that the two numerals are within delta of each other.
	//
	// 	 assert.InDelta(t, math.Pi, (22 / 7.0), 0.01)
	//
	// Returns whether the assertion was successful (true) or not (false).
	//InDelta(t TestingT, expected, actual interface{}, delta float64, msgAndArgs ...interface{}) bool

	// TODO Implement InEpsilon()
	// InEpsilon asserts that expected and actual have a relative error less than epsilon
	//
	// Returns whether the assertion was successful (true) or not (false).
	//InEpsilon(t TestingT, expected, actual interface{}, epsilon float64, msgAndArgs ...interface{}) bool

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
