package util

import (
	"fmt"
	"strings"
)

// FailIf panics if it encounters an error.
func FailIf(err error, msg ...string) {
	if err != nil {
		if len(msg) > 0 {
			err = fmt.Errorf("%s: %w", strings.Join(msg, " "), err)
		}
		panic(err)
	}
}

// Assert panics if b is false.
func Assert(actual, expected any) {
	if actual != expected {
		panic(fmt.Sprintf("Assertion failed, expected %v but got %v", expected, actual))
	}
	fmt.Print(" ✅ ")
}
