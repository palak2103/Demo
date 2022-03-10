package main

import "testing"

func test_add_5_2(t *testing.T) {
	// define variables for test / Preparation
	a := 5
	b := 2
	wanted := 7

	// call function to test
	got := Add(a, b)

	// verify result
	if got != wanted {
		t.Errorf("Add(%d,%d) = %d, wanted %d", a, b, got, wanted)
	}

	t.Logf("hello")

}
