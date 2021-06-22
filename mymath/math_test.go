package mymath

import "testing"

func TestAddAllPositive(t *testing.T) {
	returned := Add(10, 20)
	if returned != 30 {
		t.Error("wrong result")
	}
}

func TestAddAllNegative(t *testing.T) {
	returned := Add(-10, -20)
	if returned != -30 {
		t.Error("wrong result")
	}
}
