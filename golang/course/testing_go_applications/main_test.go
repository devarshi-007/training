package main_test

import (
	"testing"
)

// testings , testing/quick, testing/iotest, net/http/httptest

func TestAddition(t *testing.T) {
	got := 2 + 2
	expected := 4
	if got != expected {
		t.Errorf("ecpected: %v got: %v\n", expected, got)
	}
}

func TestSubtractions(t *testing.T) {
	got := 2 - 2.0
	expected := 0.0

	if got != expected {
		t.Errorf("expected: %v got: %v\n", expected, got)
	}
}
