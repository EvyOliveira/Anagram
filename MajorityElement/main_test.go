package main

import (
	"errors"
	"testing"
)

func TestFindMajorityElement(t *testing.T) {
	tests := []struct {
		name     string
		input    []int64
		expected int64
	}{
		{name: "slice with majority element", input: []int64{3, 2, 3}, expected: 3},
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

func TestIsAnEmptySlice(t *testing.T) {
	tests := []struct {
		name          string
		input         []int64
		expectedValue int64
		expectedError error
	}{
		{name: "empty slice", input: []int64{}, expectedValue: 0, expectedError: errors.New("slice is empty")},
		{name: "non-empty slice", input: []int64{1, 2, 3}, expectedValue: 1, expectedError: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := isAnEmptySlice(tt.input)
			if result != tt.expectedValue || !compareErrors(t, err, tt.expectedError) {
				t.Errorf("Expected value: %d, error: %v, got: %d, error: %v", tt.expectedValue, tt.expectedError, result, err)
			}
		})
	}
}

func compareErrors(t *testing.T, err1, err2 error) bool {
	if err1 == nil && err2 == nil {
		return true
	} else if err1 == nil || err2 == nil {
		return false
	}
	return err1.Error() == err2.Error()
}
