package main

import (
	"fmt"
	"strings"
)

func WordPattern(pattern, s string) bool {
	words := strings.Fields(s)
	if len(words) != len(pattern) {
		return false
	}
	patternToWord := make(map[rune]string)
	wordToPattern := make(map[string]rune)
	for i, char := range pattern {
		word := words[i]
		if wordPattern, ok := patternToWord[char]; ok {
			if wordPattern != word {
				return false
			}
		} else {
			patternToWord[char] = word
		}
		if charPattern, ok := wordToPattern[word]; ok {
			if charPattern != char {
				return false
			}
		} else {
			wordToPattern[word] = char
		}
	}
	return true
}

func main() {
	s := "dog cat cat dog"
	pattern := "abba"
	if WordPattern(pattern, s) {
		fmt.Println("the string " + s + " follows the pattern " + pattern + ".")
	} else {
		fmt.Println("the string " + s + " does not follow the pattern " + pattern + ".")
	}
}
