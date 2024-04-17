package handlers_test

import (
	"errors"
	"testing"

	"github.com/m3lnic/think-money-technical-test/internal/handlers"
	"github.com/stretchr/testify/assert"
)

var ErrMyFake error = errors.New("my fake error")

func TestErrorRes(t *testing.T) {
	newErr := handlers.NewErrorRes(ErrMyFake)
	assert.Equal(t, ErrMyFake.Error(), newErr.Message)
	assert.Equal(t, "{\"message\":\"my fake error\"}", newErr.ToString())
}
