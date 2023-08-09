package gocode

import (
	"errors"
)

type Code string

func (code Code) Error() string {
	return string(code)
}

const (
	DefaultCode Code = "0"
	SuccessCode Code = "1"
)

func Parse(err error) Code {
	if err == nil {
		return SuccessCode
	}

	var (
		code Code
		ok   bool
		e    = err
	)

	for {
		if e == nil {
			break
		}

		if code, ok = e.(Code); ok {
			break
		}

		e = errors.Unwrap(e)
	}

	if code == "" {
		code = DefaultCode
	}

	return code
}
