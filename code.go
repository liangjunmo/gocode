package gocode

import (
	"errors"
)

type Code string

func (code Code) Error() string {
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
	e := err

	for {
		if e == nil {
			break
		}

		if c, ok := e.(Code); ok {
			code = c
			break
		}

		e = errors.Unwrap(e)
	}

	return code
}
