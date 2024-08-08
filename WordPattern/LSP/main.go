package main

import (
	"fmt"
	"strings"
)

type WordProcessor interface {
	Split(s string) []string
}

type DefaultWordProcessor struct{}

func (d *DefaultWordProcessor) Split(s string) []string {
	return strings.Fields(s)
}

func WordPattern(pattern, s string, wordProcessor WordProcessor) bool {
	words := wordProcessor.Split(s)
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
	wordProcessor := DefaultWordProcessor{}

	if WordPattern(pattern, s, &wordProcessor) {
		fmt.Println("the string " + s + " follows the pattern " + pattern + ".")
	} else {
		fmt.Println("the string " + s + " does not follow the pattern " + pattern + ".")
	}

}
