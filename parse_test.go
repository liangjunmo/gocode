package gocode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseErrorIsNil(t *testing.T) {
	require.Equal(t, SuccessCode, Parse(nil))
}

func TestParseErrorIsNotNilButIsNotCode(t *testing.T) {
	require.Equal(t, DefaultCode, Parse(fmt.Errorf("err")))
}

func TestParseErrorIsWrappedButIsNotCode(t *testing.T) {
	err := fmt.Errorf("%v: %w", fmt.Errorf("err"), fmt.Errorf("err"))
	require.Equal(t, DefaultCode, Parse(err))
}

func TestParseErrorIsCode(t *testing.T) {
	notFoundCode := Code("NotFound")
	err := fmt.Errorf("%w", notFoundCode)
	require.Equal(t, notFoundCode, Parse(err))
}

func TestParseErrorIsWrappedCode(t *testing.T) {
	notFoundCode := Code("NotFound")
	err := fmt.Errorf("%v: %w", fmt.Errorf("err"), notFoundCode)
	require.Equal(t, notFoundCode, Parse(err))
}

func TestParseErrorIsMultipleWrappedCode(t *testing.T) {
	invalidRequestCode := Code("invalidRequest")
	notFoundCode := Code("NotFound")

	err := fmt.Errorf("%v: %w", fmt.Errorf("err"), invalidRequestCode)

	for i := 0; i < 10; i++ {
		err = fmt.Errorf("%v: %w", err, Code(fmt.Sprintf("%d", i)))
	}

	err = fmt.Errorf("%v: %w", err, notFoundCode)

	code := Parse(err)
	require.Equal(t, notFoundCode, code)
	require.NotEqual(t, invalidRequestCode, code)
}
