package main

import (
	"fmt"

	stderrors "errors"
	"github.com/pkg/errors"
)

func findSomething() error {
	return errors.Errorf("something not found")
}

func main() {
	err := findSomething()
	// can print error stack
	exitf("Error1: %+v", err)

	fmt.Println()
	fmt.Println()
	// cannot print error stack
	err = ErrNotFound(err)
	exitf("Error2: %+v", err)
}

func exitf(format string, args ...interface{}) {
	fmt.Printf(format, args...)

	//os.Exit(1)
}

type errNotFound struct{ error }

func ErrNotFound(err error) error {
	if err == nil || IsErrNotFound(err) {
		return err
	}

	return errNotFound{err}
}

func IsErrNotFound(err error) bool {
	return stderrors.As(err, &errNotFound{})
}
