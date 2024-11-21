package errors

import (
	"errors"
	"fmt"
)

func Wrap(err error, msg string) error {
	return fmt.Errorf("%s: %w", msg, err)
}

func Is(err error, target error) bool {
	return errors.Is(err, target)
}
