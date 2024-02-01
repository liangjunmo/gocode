package gocode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseErrorIsNil(t *testing.T) {
	var err error
	code := Parse(err)
	require.Equal(t, SuccessCode, code)
}

func TestParseErrorIsNotCode(t *testing.T) {
	err := fmt.Errorf("err")
	code := Parse(err)
	require.Equal(t, DefaultCode, code)
}

func TestParseErrorIsWrappedButIsNotCode(t *testing.T) {
	err := fmt.Errorf("err")
	err = fmt.Errorf("%v: %w", err, fmt.Errorf("err"))
	code := Parse(err)
	require.Equal(t, DefaultCode, code)
}

func TestParseErrorIsCode(t *testing.T) {
	notFoundCode := Code("NotFound")
	err := fmt.Errorf("%w", notFoundCode)
	code := Parse(err)
	require.Equal(t, notFoundCode, code)
}

func TestParseErrorIsWrappedCode(t *testing.T) {
	notFoundCode := Code("NotFound")
	err := fmt.Errorf("err")
	err = fmt.Errorf("%v: %w", err, notFoundCode)
	code := Parse(err)
	require.Equal(t, notFoundCode, code)
}

func TestParseErrorIsMultipleWrappedCode(t *testing.T) {
	invalidRequestCode := Code("invalidRequest")
	notFoundCode := Code("NotFound")
	err := fmt.Errorf("err")
	err = fmt.Errorf("%v: %w", err, invalidRequestCode)
	for i := 0; i < 10; i++ {
		err = fmt.Errorf("%v: %w", err, Code(fmt.Sprintf("%d", i)))
	}
	err = fmt.Errorf("%v: %w", err, notFoundCode)
	code := Parse(err)
	require.Equal(t, notFoundCode, code)
	require.NotEqual(t, invalidRequestCode, code)
}
