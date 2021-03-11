package main

import "testing"

func TestSum(t *testing.T) {

	result := sum(2, 3)

	if result != 5 {
		t.Error("The result must be 5")
	}
}
func TestSub(t *testing.T) {

	result := sub(4, 2)

	if result != 2 {
		t.Error("The result must be 2")
	}
}
func TestTimes(t *testing.T) {

	result := times(4, 2)

	if result != 8 {
		t.Error("The result must be 8")
	}
}
func TestSumX(t *testing.T) {

	result := sumX(4, 2)

	if result != 10 {
		t.Error("The result must be 10")
	}
}
