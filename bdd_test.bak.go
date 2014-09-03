package gomspec

import (
	"testing"
)

func Test_Bdd_Specifications(t *testing.T) {

	Given(t, "a unique scenerio", func(when When) {

		when("an event occurs", func(it It) {
			it("should evaluate to 1", func(expect Expect) {
				expect(1).ToEqual(1)
			})

			it("should also evaluate to 3", func(expect Expect) {
				expect(3).ToEqual(3)
			})

			it("should perform another evaluation", func(expect Expect) {
				expect(4).ToNotEqual(5)
			})

			it("should not have this implemented", NA())

			it("should also perform another evaluation", func(expect Expect) {
				expect("hellow").ToNotEqual("world")
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

			it("should increment count to 1", setup(func(expect Expect) {
				expect(count).ToEqual(1)
			}))

			if count != 0 {
				t.Error("In BDD-specs, count should have been reset to zero by the teardown func")
			}
		})
	})

}
