package main

import "testing"

func TestAnagramChecker_IsAnagram(t *testing.T) {
	checker := &AnagramChecker{}

	tests := []struct {
		name          string
		firstString   string
		secoundString string
		expected      bool
	}{
		{name: "it's an anagram", firstString: "anagram", secoundString: "nagaram", expected: true},
		{name: "it's not an anagram", firstString: "rat", secoundString: "car", expected: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			current := checker.IsAnagram(test.firstString, test.secoundString)
			if current != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, current)
			}
		})
	}
}
