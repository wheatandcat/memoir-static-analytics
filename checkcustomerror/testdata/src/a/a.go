package a

import "fmt"

func f() error {
	var gopher int

	if gopher == 0 {
		err := fmt.Errorf("Error: %s", gopher)
		return err // want "require customError wrap"
	}

	if gopher == 1 {
		err := fmt.Errorf("Error: %s", gopher)
		return customError(err) // OK
	}

	return nil // OK
}

func customError(err error) error {
	return err // nocheck:checkcustomerror
}
