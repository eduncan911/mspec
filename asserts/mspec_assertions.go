package asserts

import "fmt"

/*
	This is the glue used to bind Testify's assertions to the
	gomspec package as the default asserting package.

	This may be moved to gomspec's namespace.  The issue is that
	Assertions currently has a private field "t" that cannot be accessed
	outside of this page.

	t is used for running Errorf() to output the error to the *testing.T
	default implementation.

*/

type mspectTestingT struct {
}

var t = &mspectTestingT{}

func (m *mspectTestingT) Errorf(format string, args ...interface{}) {
	fmt.Println(format, args)
}

// NewAssertions constructs a wrapper around Testify's asserts.
func NewAssertions() *Assertions {
	return &Assertions{
		t: t,
	}
}
