package main

import "testing"

func TestFindMajorityElement(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{name: "array with majority element", input: []int{3, 2, 3}, expected: 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMajorityElement(tt.input)
			if result != tt.expected {
				t.Errorf("expected majority element: %d, got: %d", tt.expected, result)
			}
		})
	}
}
