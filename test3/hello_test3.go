package test3

import (
	"testing"
	"function"
)

//TestSum is to test sum function
func TestSum(t *testing.T) {

	numberReturn := function.Sum(3, 3)

	if numberReturn != 6 {
		t.Errorf("IntMin(3, 3) = %d; 6 is the correct answer", numberReturn)
	}
}