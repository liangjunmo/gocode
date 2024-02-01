package gocode

import (
	"errors"
)

type Code string

func (code Code) Error() string {
	return string(code)
}

func (code Code) String() string {
	return string(code)
}

const (
	DefaultCode Code = "default"
	SuccessCode Code = "success"
)

func Parse(err error) Code {
	if err == nil {
		return SuccessCode
	}

	code := DefaultCode

	for {
		if err == nil {
			break
		}

		if errors.As(err, &code) {
			break
		}

		err = errors.Unwrap(err)
	}

	return code
}
