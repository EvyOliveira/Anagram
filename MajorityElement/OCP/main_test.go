package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMajorityElement(t *testing.T) {
	tests := []struct {
		name                 string
		input                []int64
		expectedMajority     int64
		expectedError        bool
		expectedErrorMessage string
	}{
		{name: "array with majority element", input: []int64{3, 2, 3}, expectedMajority: 3, expectedError: false},
		{name: "array with single element", input: []int64{1}, expectedMajority: 1, expectedError: false},
		{name: "empty slice", input: []int64{}, expectedMajority: 0, expectedError: true, expectedErrorMessage: "slice is empty"},
	}

	checker := &MajorityElement{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := checker.FindMajorityElement(tt.input)
			if tt.expectedError {
				assert.EqualError(t, err, tt.expectedErrorMessage)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedMajority, result)
			}
		})
	}
}
