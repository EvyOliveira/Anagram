package main

import "testing"

func TestAnagramChecker_IsAnagram(t *testing.T) {
	tests := []struct {
		name          string
		firstString   string
		secoundString string
		expected      bool
	}{
		{name: "it's an anagram", firstString: "anagram", secoundString: "nagaram", expected: true},
		{name: "it's not an anagram", firstString: "rat", secoundString: "car", expected: false},
		{name: "empty string", firstString: "", secoundString: "test", expected: false},
		{name: "both empty strings", firstString: "", secoundString: "", expected: false},
		{name: "different case", firstString: "Race", secoundString: "care", expected: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			anagramChecker := AnagramChecker{}
			current := anagramChecker.IsAnagram(test.firstString, test.secoundString)
			if current != test.expected {
				t.Errorf("Test case '%s': expected %v, got %v", test.name, test.expected, current)
			}
		})
	}
}
