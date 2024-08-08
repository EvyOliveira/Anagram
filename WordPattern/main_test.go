package main

import "testing"

func TestWordPattern(t *testing.T) {
	tests := []struct {
		name     string
		pattern  string
		s        string
		expected bool
	}{
		{name: "the string follows the pattern", pattern: "abba", s: "dog cat cat dog", expected: true},
		{name: "the string does not follow the pattern", pattern: "abba", s: "dog cat cat fish", expected: false},
		{name: "empty string", pattern: "abba", s: "", expected: false},
		{name: "empty pattern", pattern: "", s: "dog cat cat dog", expected: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			current := WordPattern(test.pattern, test.s)
			if current != test.expected {
				t.Errorf("Test case '%s': expected %v, got %v", test.name, test.expected, current)
			}
		})
	}
}
