package main

import (
	"reflect"
	"testing"
)

func TestCounter(t *testing.T) {
	input := "Apple banana apple orange! BANANA apple aPPLee."
	expected := map[string]int{
		"apple":  3,
		"banana": 2,
		"orange": 1,
		"applee": 1,
	}

	result := counter(input)

	// Use reflect.DeepEqual to compare maps
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("counter(%q) = %v; expected %v", input, result, expected)
	}
}
