package a_test

import (
	"errors"
	"testing"
)

func e() (string, error) {
	return "test", errors.New("test") // OK testファイルは除外する
}

func Test_F(t *testing.T) {
	e()
}
