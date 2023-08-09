package gocode

import (
	"errors"
)

type Code string

func (code Code) Error() string {
	return string(code)
}

var (
	defaultCode Code
	successCode Code
)

func SetDefaultCode(code Code) {
	defaultCode = code
}

func SetSuccessCode(code Code) {
	successCode = code
}

func Parse(err error) Code {
	if err == nil {
		return successCode
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
		code = defaultCode
	}

	return code
}
