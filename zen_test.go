package gomspec

import (
	"testing"
)

// These specs are kept to show backwards compatibility
// with the original Zen project it was forked from.
//

func Test_Zen_compatibility_with_ExampleDesc(t *testing.T) {

	Desc(t, "testing Equality specs", func(it It) {
		it("should have an integer equal to itself", func(expect Expect) {
			expect(1).ToEqual(1)
		})
	})
}

func Test_Zen_compatibility_with_SetupAndTeardown(t *testing.T) {
	count := 0

	before := func() {
		count++
	}

	after := func() {
		count--
	}

	setup := Setup(before, after)

	Desc(t, "using Setup with specs", func(it It) {
		it("should execute before by incrementing count", setup(func(expect Expect) {
			expect(count).ToEqual(1)
		}))

		if count != 0 {
			t.Error("Count should have been reset to zero by the teardown func")
		}
	})
}
