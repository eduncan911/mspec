package test

import "github.com/eduncan911/mspec/bdd"

//"testing"

// func Example_BDD_Specifications(t *testing.T) {

// 	Given(t, "a unique scenerio", func(when When) {

// 		when("an event occurs", func(it It) {
// 			it("should evaluate 1s are equal", func(assert bdd.Assert) {
// 				assert.Equal(1, 1)
// 			})

// 			it("should also evaluate 3 and 4 are not equal", func(assert bdd.Assert) {
// 				assert.NotEqual(3, 4)
// 			})

// 			it("should not have this implemented")

// 			it("should perform another evaluation", func(assert bdd.Assert) {
// 				assert.Contains("shoppy", "opp")
// 			})

// 			it("should error here", func(assert bdd.Assert) {
// 				assert.True(false)
// 			})

// 			it("should also perform another evaluation", func(assert bdd.Assert) {
// 				assert.NotEqual("hello", "world")
// 			})
// 		})
// 	})

// 	Given(t, "a scenario that needs a setup and teardown", func(when When) {

// 		count := 0

// 		when("using the Setup() extension", func(it It) {

// 			before := func() {
// 				count++
// 			}

// 			after := func() {
// 				count--
// 			}

// 			setup := Setup(before, after)

// 			it("should increment count to 1", setup(func(assert bdd.Assert) {
// 				assert.Equal(1, count)
// 			}))

// 			it("should decrement comment during teardown back to 0", func(assert bdd.Assert) {
// 				assert.Equal(0, count)
// 			})
// 		})
// 	})

// }

func Example() {

	given, _, _ := bdd.Sentences()

	//var t *testing.T
	// you can quickly spec new features with little syntax noise
	//

	// GIVEN a valid Api, what shall we do?  not sure yet.
	//
	given(nil, "a valid Api")

	// GIVEN an invalid Api...
	//
	given(nil, "an invalid Api", func(when bdd.When) {

		// ...WHEN GetUsers is called, we don't know what SHOULD happen yet.
		//
		when("GetUsers is called")

		// ...WHEN GetStatus is called...
		//
		when("GetStatus is called", func(it bdd.It) {

			// ...IT SHOULD return an invalid status code
			it("should return an invalid status code")

			// ...IT SHOULD return an error message
			it("should return an error message")

		})

	})
}
