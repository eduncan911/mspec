package mspec

import (
	"testing"
)

func Test_MSpec_Instances(t *testing.T) {

	Given(t, "an mspec instance", func(when When) {

		f := "feature value"

		mspecTest := &MSpecConfig{
			lastFeature: f,
			lastGiven:   "context value",
			lastWhen:    "when value",
			lastSpec:    "title value",
		}

		when("calling reset()", func(it It) {

			mspecTest.resetLasts()

			it("should not reset lastFeature as that is used globally", func(assert Assert) {
				assert.NotEmpty(mspecTest.lastFeature)
			})

			it("should keep the lastFeature value", func(assert Assert) {
				assert.Equal(mspecTest.lastFeature, f)
			})

			it("should set lastGiven to zero value", func(assert Assert) {
				assert.Empty(mspecTest.lastGiven)
			})

			it("should set lastWhen to zero value", func(assert Assert) {
				assert.Empty(mspecTest.lastWhen)
			})

			it("should set lastSpec to zero value", func(assert Assert) {
				assert.Empty(mspecTest.lastSpec)
			})
		})
	})
}
