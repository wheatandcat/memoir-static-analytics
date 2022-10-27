package a

import (
	"errors"
	"fmt"
)

func f() error {
	var gopher int

	if gopher == 0 {
		err := fmt.Errorf("Error: %s", gopher)
		return err // want "require customError wrap"
	}

	if gopher == 1 {
		err := fmt.Errorf("Error: %s", gopher)
		return CustomError(err) // OK
	}

	return nil // OK
}

func CustomError(err error) error {
	return err // nocheck:checkcustomerror
}

func a() (string, error) {
	err := errors.New("test")
	return "test", err // want "require customError wrap"
}

func b() (string, error) {
	err := errors.New("test")
	return "test", CustomError(err) // OK
}

func c() (string, error) {
	return "test", errors.New("test") // want "require customError wrap"
}

func d() (string, error) {
	return "test", CustomError(errors.New("test")) // OK
}
