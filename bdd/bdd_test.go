package bdd

import (
	"testing"
)

// Some concepts defined:
// - Simple Call: sentence(event, func(it))
// - Multiple Call: sentence(event, func(it, args), like)

// Feature Simple or Multiple Given
// - As a developer,
// - I want to be able to test given with a simple call, or with a multiple call,
// - So that I have more testing power.
func Test_Simple_and_Multiple_Given(t *testing.T) {
	given, like, s := Sentences()

	given(t, "a TestSumOp ts with no handicap", func(when When) {
		ts := NewTestSumOp(0)

		when("ts.LastResultAsString() is called", func(it It) {
			val := ts.LastResultAsString()

			it("should return a empty string", func(assert Assert) {
				assert.Empty(val)
			})
		})

		when("ts.Sum(1, 2) is called", func(it It) {
			val := ts.Sum(1, 2)

			it("should return 3", func(assert Assert) {
				assert.Equal(val, 3)
			})

			it("should have ts.LastResultAsString() return '3'", func(assert Assert) {
				assert.Equal(ts.LastResultAsString(), "3")
			})
		})
	})

	given(t, "a TestSumOp ts with handicap %[1]v", func(when When, args ...interface{}) {
		ts := NewTestSumOp(args[0].(int))

		when("ts.LastResultAsString() is called", func(it It) {
			val := ts.LastResultAsString()

			it("should return a empty string", func(assert Assert) {
				assert.Empty(val)
			})
		})

		when("ts.Sum(%[2]v, %[3]v) is called", func(it It) {
			val := ts.Sum(args[1].(int), args[2].(int))

			it("should return %[4]v", func(assert Assert) {
				assert.Equal(val, args[3].(int))
			})

			it("should have ts.LastResultAsString() return '%[5]v'", func(assert Assert) {
				assert.Equal(ts.LastResultAsString(), args[4].(string))
			})
		})
	}, like(
		s(0, 1, 2, 3, "3"), s(1, 1, 2, 4, "4"), s(0, 2, 3, 5, "5"),
		s(-1, 2, 3, 4, "4"), s(0, -3, 2, -1, "-1"), s(2, -3, 2, 1, "1"),
		s(0, 2, 0, 2, "2"), s(-2, 2, 0, 0, "0"),
	))
}

// Feature Simple or Multiple When
// - As a developer,
// - I want to be able to test when with a simple call, or with a multiple call,
// - So that I have more testing power.
func Test_Simple_and_Multiple_When(t *testing.T) {
	given, like, s := Sentences()

	given(t, "a empty TestSumOp ts", func(when When) {
		var ts TestSumOp

		when("ts.LastResultAsString() is called", func(it It) {
			val := ts.LastResultAsString()

			it("should return a empty string", func(assert Assert) {
				assert.Empty(val)
			})
		})

		when("ts.Sum(%[1]v, %[2]v) is called", func(it It, args ...interface{}) {
			val := ts.Sum(args[0].(int), args[1].(int))

			it("should return %[3]v", func(assert Assert) {
				assert.Equal(val, args[2])
			})
			it("should have ts.LastResultAsString() return '%[4]v'", func(assert Assert) {
				assert.Equal(ts.LastResultAsString(), args[3])
			})
		}, like(
			s(1, 2, 3, "3"), s(2, 3, 5, "5"), s(-3, 2, -1, "-1"),
			s(2, 0, 2, "2"), s(12, 21, 33, "33"), s(-18, 5, -13, "-13"),
		))
	})
}

// Feature Simple and Multiple it
// - As a developer,
// - I want to be able to test it with a simple call, or with a multiple call,
// - So that I have more testing power.
func Test_Simple_and_Multiple_It(t *testing.T) {
	given, like, s := Sentences()

	given(t, "a empty TestSumOp ts", func(when When) {
		var ts TestSumOp

		when("val := ts.Sum(1, 2) is called", func(it It) {
			val := ts.Sum(1, 2)

			it("should have val equal to 3", func(assert Assert) {
				assert.Equal(val, 3)
			})
			it("should have ts.LastResultAsString() return '3'", func(assert Assert) {
				assert.Equal(ts.LastResultAsString(), "3")
			})
			it("should have TestInt(val).Sum(%[1]v) return %[2]v", func(assert Assert, args ...interface{}) {
				assert.Equal(TestInt(val).Sum(args[0].(int)), args[1].(int))
			}, like(
				s(0, 3), s(1, 4), s(2, 5), s(10, 13), s(3000, 3003),
				s(-1, 2), s(-3, 0), s(-10, -7), s(-3000, -2997),
			))
		})
	})
}

// Feature Mixed Multiple sentences
// - As a developer,
// - I want to be able to test with mixed multiple sentences, like given, when and it,
// - So that I have more testing power.
func Test_Mixed_Multiple_sentences(t *testing.T) {
	given, like, s := Sentences()

	given(t, "a empty TestSumOp ts with handicap %[1]v", func(when When, args ...interface{}) {
		h := args[0].(int)
		ts := NewTestSumOp(h)

		when("val := ts.Sum(%[1]v, %[2]v) is called", func(it It, args ...interface{}) {
			a := args[0].(int)
			b := args[1].(int)
			val := ts.Sum(a, b)

			it("should have TestInt(val).Sum(%[1]v) return %[2]v", func(assert Assert, args ...interface{}) {
				assert.Equal(TestInt(val).Sum(args[0].(int)), args[1].(int))
			}, like(
				s(1, 1+a+b+h), s(2, 2+a+b+h),
				s(3, 3+a+b+h), s(12, 12+a+b+h),
			))
		}, like(
			s(1, 2), s(2, 3), s(2, 0),
			s(1, -2), s(3, 0),
		))
	}, like(
		s(1), s(2), s(3), s(10),
		s(21), s(637), s(-1),
	))
}

// Feature Simple and Multiple NonImplemented sentences
// - As a developer,
// - I want to be able to test with simple and multiple NonImplemented sentences,
// - So I have more warnings to sentences not implemented.
func Test_Simple_and_Multiple_NonImplemented_sentences(t *testing.T) {
	given, like, s := Sentences()

	given(t, "a empty TestSumOp ts", func(when When) {
		when("ts.LastResultAsString() is called", func(it It) {
			it("should return a empty string")
		})
		when("ts.Sum(%[1]v, %[2]v) is called", like(
			s(1, 2),
			s(-1, 2),
			s(-1, 0),
			s(3, 0),
			s(13, 2),
		))
	})

	given(t, "a empty TestSumOp ts with handicap 3", func(when When) {

		when("val := ts.Sum(1, 2) is called", func(it It) {
			it("should have val equal to 3")
			it("should have ts.LastResultAsString() return '3'")
			it("should have TestInt(val).Sum(%[1]v) return %[2]v", like(
				s(0, 3), s(1, 4), s(2, 5), s(10, 13), s(3000, 3003),
				s(-1, 2), s(-3, 0), s(-10, -7), s(-3000, -2997),
			))
		})
	})
}
