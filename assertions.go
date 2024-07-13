package assertions

import (
	"fmt"
	"os"
	"testing"
)

func Assert(t testing.TB, condition bool, description string) {
	if condition {
		return
	}
	fmt.Fprintln(os.Stderr, getExpectedGivenMessage(description, true, false))
	t.FailNow()
}

func AssertEquals(t testing.TB, expected any, given any, description string) {
	if expected == given {
		return
	}
	fmt.Fprintln(os.Stderr, getExpectedGivenMessage(description, expected, given))
	t.FailNow()
}
