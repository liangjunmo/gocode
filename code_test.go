package gocode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/liangjunmo/gocode"
)

func TestCode(t *testing.T) {
	var err error
	code := gocode.Parse(err)
	assert.Equal(t, gocode.SuccessCode, code)

	err = fmt.Errorf("error")
	code = gocode.Parse(err)
	assert.Equal(t, gocode.DefaultCode, code)

	var notFoundCode gocode.Code = "not_found"
	err = fmt.Errorf("%w: %s", notFoundCode, "uid 1 not found")
	code = gocode.Parse(err)
	assert.Equal(t, notFoundCode, code)
}
