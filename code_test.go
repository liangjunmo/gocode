package gocode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/liangjunmo/gocode"
)

func TestCode(t *testing.T) {
	var (
		defaultCode  gocode.Code = "unknown"
		successCode  gocode.Code = "success"
		notFoundCode gocode.Code = "not_found"
	)

	gocode.SetDefaultCode(defaultCode)
	gocode.SetSuccessCode(successCode)

	var err error

	code := gocode.Parse(err)
	assert.Equal(t, successCode, code)

	err = fmt.Errorf("%w: %s", notFoundCode, "uid 1 not found")

	code = gocode.Parse(err)
	assert.Equal(t, notFoundCode, code)
}
