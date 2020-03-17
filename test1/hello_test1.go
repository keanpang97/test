package test1

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

//TestSum is to test sum function
func TestSum(t *testing.T) {

	numberReturn := function.Sum(3, 3)

	if numberReturn != 6 {
		t.Errorf("IntMin(3, 3) = %d; 6 is the correct answer", numberReturn)
	}
}
