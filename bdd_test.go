package gomspec

import (
	"testing"
)

func Test_Bdd_Specifications(t *testing.T) {

	Given(t, "a unique scenerio", func(when When) {

		when("an event occurs", func(it It) {
			it("should evaluate 1s are equal", func(assert Assert) {
				assert.Equal(1, 1)
			})

			it("should also evaluate 3 and 4 are not equal", func(assert Assert) {
				assert.NotEqual(3, 4)
			})

			it("should not have this implemented")

			it("should perform another evaluation", func(assert Assert) {
				assert.Contains("shoppy", "opp")
			})

			it("should error here", func(assert Assert) {
				assert.True(false)
			})

			it("should also perform another evaluation", func(assert Assert) {
				assert.NotEqual("hello", "world")
			})
		})
	})

	Given(t, "a scenario that needs a setup and teardown", func(when When) {

		count := 0

		when("using the Setup() extension", func(it It) {

			before := func() {
				count++
			}

			after := func() {
				count--
			}

			setup := Setup(before, after)

			it("should increment count to 1", setup(func(assert Assert) {
				assert.Equal(1, count)
			}))

			it("should decrement comment during teardown back to 0", func(assert Assert) {
				assert.Equal(0, count)
			})
		})
	})

}
