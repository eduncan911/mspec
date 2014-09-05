package gomspec

import (
	"fmt"
	"github.com/eduncan911/gomspec/asserts"
	"strings"
)

/*
	This is the glue used to bind Testify's assertions to the
	gomspec package as the default asserting package.

	mspecTestingT implements the asserts.TestingT interface
	with Errorf(...) func.

	It is used to print the specification portion to the output
	when an error occurs.  Also, it sets a flag that is used by
	the bdd framework to know that an error has been printed and
	therefore do not print a normal specification title.

	Errors are only handled this way under one condition: that
	is that Errorf() be executed by the Assertion package.  Else,
	we do not get the flag to know that an error has been found.

	The current Testify assert package fires Errorf() on every
	Fail(), which all asserts will fire when an error occurs.  So,
	we just wrap that below.

*/

type mspectTestingT struct {
	spec *Specification
}

func (m *mspectTestingT) Errorf(format string, args ...interface{}) {
	// because we control the output of specification, we
	// need to store these details in a state for later use in
	// the bdd framework.

	// string foo
	err := fmt.Sprintf(format, args...)
	err = strings.Replace(err, "\r", "", -1)
	err = strings.Replace(err, "        ", "\t\t\t", -1)

	// parse out Testify's location info by removing the first line
	lines := strings.Split(err, "\n")
	out := ""
	for i := range lines {
		if strings.Contains(lines[i], "Location:") {
			continue
		}
		if lines[i] == "" {
			continue
		}
		if out == "" {
			out = lines[i]
		} else {
			out = strings.Join([]string{out, "\n", lines[i]}, "")
		}
	}

	// if len(lines) >= 4 {
	// 	out = strings.Join([]string{lines[1], "\n", lines[2]}, "")
	// } else if len(lines) == 3 {
	// 	out = lines[1]
	// }

	m.spec.PrintTitleWithError()
	m.spec.PrintError(out)
	m.spec.AssertionFailed = true
}

// NewAssertions constructs a wrapper around Testify's asserts.
func newAssertions(s *Specification) *asserts.Assertions {
	return asserts.New(
		&mspectTestingT{
			spec: s,
		})
}
