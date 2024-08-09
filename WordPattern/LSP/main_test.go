package main

import "testing"

func TestWordPattern(t *testing.T) {
	type testCase struct {
		name      string
		pattern   string
		s         string
		expected  bool
		processor WordProcessor
	}
	tests := []testCase{
		{name: "the string follows the pattern", pattern: "abba", s: "dog cat cat dog", expected: true, processor: &DefaultWordProcessorAdapter{&DefaultWordProcessor{}}},
		{name: "the string does not follow the pattern", pattern: "abba", s: "dog cat cat fish", expected: false, processor: &DefaultWordProcessorAdapter{&DefaultWordProcessor{}}},
		{name: "empty string", pattern: "abba", s: "", expected: false, processor: &DefaultWordProcessorAdapter{&DefaultWordProcessor{}}},
		{name: "empty pattern", pattern: "", s: "dog cat cat dog", expected: false, processor: &DefaultWordProcessorAdapter{&DefaultWordProcessor{}}},
		{name: "adapter: empty delimiter", pattern: "abba", s: "dog cat cat dog", expected: true, processor: &DefaultWordProcessorAdapter{&DefaultWordProcessor{}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			processor := test.processor
			current := WordPattern(test.pattern, test.s, processor)
			if current != test.expected {
				t.Errorf("Test case '%s': expected %v, got %v", test.name, test.expected, current)
			}
		})
	}
}
