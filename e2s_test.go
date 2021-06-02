package e2s

import (
	"errors"
	"testing"
)

func TestErrFatal(t *testing.T) {
	IfErrFatal(nil)
}

func TestErrWarning(t *testing.T) {
	IfErrWarning(nil)
	IfErrWarning(errors.New("Test"))
}

func TestWarning(t *testing.T) {
	Warning(errors.New("Test"))
}
