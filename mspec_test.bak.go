package gomspec

import (
	"testing"
)

func Test_Mspec_Instances(t *testing.T) {

	Given(t, "an mspec instance", func(when When) {

		f := "feature value"

		mspecTest := &Mspec{
			lastFeature: f,
			lastContext: "context value",
			lastWhen:    "when value",
			lastTitle:   "title value",
		}

		when("calling reset()", func(it It) {

			mspecTest.resetLasts()

			it("should not reset lastFeature as that is used globally", func(expect Expect) {
				expect(mspecTest.lastFeature).ToNotEqual("")
			})

			it("should keep the lastFeature value", func(expect Expect) {
				expect(mspecTest.lastFeature).ToEqual(f)
			})

			it("should set lastContext to zero value", func(expect Expect) {
				expect(mspecTest.lastContext).ToEqual("")
			})

			it("should set lastWhen to zero value", func(expect Expect) {
				expect(mspecTest.lastWhen).ToEqual("")
			})

			it("should set lastTitle", func(expect Expect) {
				expect(mspecTest.lastTitle).ToEqual("")
			})
		})
	})
}
