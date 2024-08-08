package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

type WordProcessor interface {
	Split(s string) []string
}

type DefaultWordProcessor struct{}

func (d *DefaultWordProcessor) Split(s string) []string {
	return strings.Fields(s)
}

type CustomDelimiterWordProcessor struct {
	delimiter string
}

func (c *CustomDelimiterWordProcessor) Split(s string) []string {
	return strings.Split(s, c.delimiter)
}

type RegexWordProcessor struct {
	regex *regexp.Regexp
}

func (r *RegexWordProcessor) Split(s string) []string {
	return r.regex.FindAllString(s, -1)
}

type AdvancedWordProcessor struct{}

func (a *AdvancedWordProcessor) Split(s string) []string {
	var words []string
	var word []rune
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			word = append(word, r)
		} else if len(word) > 0 {
			words = append(words, string(word))
			word = nil
		}
	}
	if len(word) > 0 {
		words = append(words, string(word))
	}
	return words
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
	pattern := "aaaa"
	defaultProcessor := &DefaultWordProcessor{}

	if WordPattern(pattern, s, defaultProcessor) {
		fmt.Println("the string " + s + " follows the pattern " + pattern + ".")
	} else {
		fmt.Println("the string " + s + " does not follow the pattern " + pattern + ".")
	}

}
