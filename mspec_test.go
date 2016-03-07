package mspec

import (
	"testing"
)

func Test_MSpec_Instances(t *testing.T) {

	Given(t, "an mspec instance", func(when When) {

		f := "feature value"

		c := &Config{
			lastFeature: f,
			lastGiven:   "context value",
			lastWhen:    "when value",
			lastSpec:    "title value",
		}

		when("calling reset()", func(it It) {

			c.resetLasts()

			it("should not reset lastFeature as that is used globally", func(assert Assert) {
				assert.NotEmpty(c.lastFeature)
			})

			it("should keep the lastFeature value", func(assert Assert) {
				assert.Equal(c.lastFeature, f)
			})

			it("should set lastGiven to zero value", func(assert Assert) {
				assert.Empty(c.lastGiven)
			})

			it("should set lastWhen to zero value", func(assert Assert) {
				assert.Empty(c.lastWhen)
			})

			it("should set lastSpec to zero value", func(assert Assert) {
				assert.Empty(c.lastSpec)
			})
		})
	})
}
