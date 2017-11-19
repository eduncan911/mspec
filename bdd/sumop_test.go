package bdd

import (
	"fmt"
)

// TestSumOper is a test interface to suite test a struct.
type TestSumOper interface {
	Sum(int, int) int
	LastResultAsString() string
}

// TestInt is a named type after int for test purposes.
type TestInt int

// Sum is a function to sum a TestInt to int of TestInt types.
func (t TestInt) Sum(a interface{}) (n TestInt) {
	switch val := a.(type) {
	case int:
		n = t + TestInt(val)
	case TestInt:
		n = t + val
	default:
		n = TestInt(0)
	}
	return
}

// TestSumOper is a test structure to suite test a struct.
type TestSumOp struct {
	lastResultAsString string
	Handicap           TestInt
}

// NewTestSumOp creates a new TestSumOp with a handicap h.
func NewTestSumOp(h int) (t *TestSumOp) {
	t = &TestSumOp{
		Handicap: TestInt(h),
	}
	return
}

// Sum is a test method to suite test a struct.
func (s *TestSumOp) Sum(a, b int) (x int) {
	x = int(TestInt(a).Sum(b).Sum(s.Handicap))
	s.lastResultAsString = fmt.Sprintf("%v", x)
	return
}

// LastResultAsString is a test method to suite test a struct.
func (s *TestSumOp) LastResultAsString() (r string) {
	r = s.lastResultAsString
	return
}
