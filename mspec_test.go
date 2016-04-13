package mspec

import (
	"math"
	"testing"
)

func Test_MSpec_Instances(t *testing.T) {

	SetSilent()

	Given(t, "an mspec instance", func(when When) {

		f := "feature value"

		c := &MSpecConfig{
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

func BenchmarkGivenStub(b *testing.B) {
	SetSilent()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := &testing.T{}
		Given(t, "a single given")
	}
}

func BenchmarkWhenStub(b *testing.B) {
	SetSilent()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := &testing.T{}
		Given(t, "a single given", func(when When) {
			when("a single when")
		})
	}
}

func BenchmarkThenStub(b *testing.B) {
	SetSilent()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := &testing.T{}
		Given(t, "a single given", func(when When) {
			when("a single when", func(it It) {
				it("should have a single then")
			})
		})
	}
}

func BenchmarkError(b *testing.B) {
	SetSilent()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := &testing.T{}
		Given(t, "a context to fail", func(when When) {
			when("prepping to call the thing to fail", func(it It) {
				it("should fail", func(assert Assert) {
					assert.True(false)
				})
			})
		})
	}
}

func BenchmarkSimpleMspec(b *testing.B) {
	SetSilent()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := &testing.T{}
		Given(t, "xyz", func(when When) {
			ii := int8(10)
			when("123", func(it It) {
				it("should be this", func(assert Assert) {
					if !assert.Equal(10, int(ii)) {
						b.Fail()
					}
				})
			})
		})
	}
}

func BenchmarkSimpleTest(b *testing.B) {
	SetSilent()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := &testing.T{} // get accurate GC numbers
		if t == nil {
			b.Fail()
		}
		ii := int8(10)
		if 10 != int(ii) {
			b.Fail()
		}
	}
}

func BenchmarkComplexMspec(b *testing.B) {
	SetSilent()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := &testing.T{}
		Given(t, "a struct we create often", func(when When) {

			x := struct {
				Value string
				Log   float64
			}{}

			when("we do something more complex", func(it It) {

				x.Value = "a string to be set"
				x.Log = math.Log(20)
				ii := int8(10)

				it("should have x.Value be what we expect", func(assert Assert) {
					if !assert.Equal("a string to be set", x.Value) {
						b.Fail()
					}
				})

				it("should have x.Log be what we expect", func(assert Assert) {
					if !assert.Equal(2.995732273553991, x.Log) {
						b.Fail()
					}
				})

				it("should be true", func(assert Assert) {
					if !assert.Equal(10, int(ii)) {
						b.Fail()
					}
				})
			})

			when("we do something else", func(it It) {

				x.Value = "another string to be set"
				x.Log = math.Log(15)

				it("should have x.Value be what we expect", func(assert Assert) {
					if !assert.Equal("another string to be set", x.Value) {
						b.Fail()
					}
				})

				it("should have x.Log be what we expect", func(assert Assert) {
					if !assert.Equal(2.70805020110221, x.Log) {
						b.Fail()
					}
				})

				it("should be true", func(assert Assert) {
					if !assert.Equal(math.Log2E, 1/math.Ln2) {
						b.Fail()
					}
				})
			})
		})
	}
}

func BenchmarkComplexTest(b *testing.B) {
	SetSilent()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := &testing.T{} // get accurate GC numbers
		if t == nil {
			b.Fail()
		}
		x := struct {
			Value string
			Log   float64
		}{}

		x.Value = "a string to be set"
		x.Log = math.Log(20)
		ii := int8(10)

		if x.Value != "a string to be set" {
			b.Fail()
		}

		if x.Log != 2.995732273553991 {
			b.Fail()
		}

		if 10 != int(ii) {
			b.Fail()
		}

		x.Value = "another string to be set"
		x.Log = math.Log(15)

		if x.Value != "another string to be set" {
			b.Fail()
		}

		if x.Log != 2.70805020110221 {
			b.Fail()
		}

		if math.Log2E != 1/math.Ln2 {
			b.Fail()
		}
	}
}
