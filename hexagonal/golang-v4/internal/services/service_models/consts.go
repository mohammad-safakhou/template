package service_models

import (
	"errors"
	"fmt"
)

const (
	Tenant = "RighTel"
)

var (
	ErrRecordNotFound = errors.New("record not found")

	ErrInternalProblem = errors.New("internal error")
)

func NewErrInternalProblem(why string) error {
	return fmt.Errorf("%w: %s", ErrInternalProblem, why)
}
