package testdemo

import "testing"

func TestPass(t *testing.T) {
	expected := 3
	if observed := Adder(1, 2); observed != expected {
		t.Fatalf("Adder(1, 2) = %v, want %v", observed, expected)
	}
}

func TestSkip(t *testing.T) {
	t.Skip("Skipping test")
}

func TestFail(t *testing.T) {
	expected := 3
	if observed := Adder(2, 2); observed != expected {
		t.Fatalf("Adder(2, 2) = %v, want %v", observed, expected)
	}
}
