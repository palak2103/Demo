package main

import "testing"

func test_add_5_2(t *testing.T) {
	t.Logf("hello")
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
}
func test_add_2_0(t *testing.T) {
	t.Logf("hello")
	// define variables for test / Preparation
	a := 2
	b := 0
	wanted := 2

	// call function to test
	got := Add(a, b)

	// verify result
	if got != wanted {
		t.Errorf("Add(%d,%d) = %d, wanted %d", a, b, got, wanted)
	}
}
func test_add_0_0(t *testing.T) {
	t.Logf("hello")
	// define variables for test / Preparation
	a := 0
	b := 0
	wanted := 0

	// call function to test
	got := Add(a, b)

	// verify result
	if got != wanted {
		t.Errorf("Add(%d,%d) = %d, wanted %d", a, b, got, wanted)
	}
}
