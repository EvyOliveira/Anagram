package main

import "testing"

func TestFindMajorityElement(t *testing.T) {
	tests := []struct {
		name             string
		input            []int
		expectedMajority int
		isMajority       bool
	}{
		{name: "array with majority element", input: []int{3, 2, 3}, expectedMajority: 3, isMajority: true},
		{name: "array with single element", input: []int{1}, expectedMajority: 1, isMajority: true},
	}

	checker := &SimpleMajorityChecker{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, isMajority := findMajorityElement(tt.input, checker)
			if result != tt.expectedMajority || isMajority != tt.isMajority {
				t.Errorf("Expected majority element: %d, isMajority: %t, got: %d, isMajority: %t", tt.expectedMajority, tt.isMajority, result, isMajority)
			}
		})
	}
}
