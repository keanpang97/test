package test2

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"function"
)

//TestHello is to test hello function
func TestHello(t *testing.T) {

	stringReturn := function.Hello("abc")

	assert.Equal(t, stringReturn, "Hello abc", "function hello is not functioning well.")
}
