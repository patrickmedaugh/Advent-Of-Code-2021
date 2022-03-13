package main

import "testing"

func TestMain(t *testing.T) {
	inputs := [6]string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}
	result := CalcPositions(inputs)
	if result != 150 {
		t.Errorf("Result was not correct, got: %d, want: %d", result, 150)
	}
}
