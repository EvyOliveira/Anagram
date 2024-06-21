package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMajorityElement(t *testing.T) {
	tests := []struct {
		name             string
		input            []int64
		candidate        int64
		expectedMajority bool
	}{
		{name: "array with majority element", input: []int64{3, 2, 3}, candidate: 3, expectedMajority: true},
		{name: "array with single element", input: []int64{1}, candidate: 1, expectedMajority: true},
		{name: "empty slice", input: []int64{}, candidate: 1, expectedMajority: false},
	}

	//checker := &SimpleMajorityChecker{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := checker.CheckMajority(tt.input, tt.candidate)
			assert.Equal(t, tt.expectedMajority, actual)
		})
	}
}
