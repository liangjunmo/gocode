package gocode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/liangjunmo/gocode"
)

func TestCode(t *testing.T) {
	var err error
	code := gocode.Parse(err) // err is nil
	assert.Equal(t, gocode.SuccessCode, code)

	err = fmt.Errorf("error")
	code = gocode.Parse(err) // cannot parse err
	assert.Equal(t, gocode.DefaultCode, code)

	var notFoundCode gocode.Code = "NotFound"
	err = fmt.Errorf("%s: %w", "some message", notFoundCode)
	code = gocode.Parse(err) // can parse err
	assert.Equal(t, notFoundCode, code)
}
