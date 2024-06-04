package main

import (
	"fmt"
	"strings"
)

type Anagram interface {
	IsAnagram(string, string) bool
}

type AnagramChecker struct{}

func (a *AnagramChecker) IsAnagram(firstString, secoundString string) bool {
	preparedFirst, preparedSecound, err := prepareStrings(firstString, secoundString)
	if err != nil {
		return false
	}

	charactersMap := make(map[rune]int)
	for _, character := range preparedFirst {
		charactersMap[character]++
	}

	for _, character := range preparedSecound {
		if count, ok := charactersMap[character]; ok {
			count--
			if count == 0 {
				delete(charactersMap, character)
			} else {
				charactersMap[character] = count
			}
		} else {
			return false
		}
	}
	return len(charactersMap) == 0
}

func prepareStrings(firstString, secoundString string) (string, string, error) {
	if firstString == "" || secoundString == "" {
		return "", "", fmt.Errorf("empty strings are not allowed")
	}
	return strings.ToLower(firstString), strings.ToLower(secoundString), nil
}

func main() {
	anagramChecker := AnagramChecker{}

	firstString := "anagram"
	secoundString := "nagaram"
	firstAndSecoundComparison := anagramChecker.IsAnagram(firstString, secoundString)
	fmt.Printf("First word is %s and secound word is %s. They are anagrams? %v\n", firstString, secoundString, firstAndSecoundComparison)

	thirdWord := "rat"
	fourthWord := "car"
	thirdAndFourthComparison := anagramChecker.IsAnagram(thirdWord, fourthWord)
	fmt.Printf("Third word is %s and fourth word is %s. They are anagrams? %v\n", thirdWord, fourthWord, thirdAndFourthComparison)

	emptyString := ""
	fifthWord := "test"
	emptyStringAndFifthWordComparison := anagramChecker.IsAnagram(emptyString, fifthWord)
	fmt.Printf("Empty string is %s and fifth word is %s. They are anagrams? %v\n", emptyString, fifthWord, emptyStringAndFifthWordComparison)

	sixthWord := "iRaceMa"
	seventhWord := "aMerIca"
	sixthWordAndseventhWordComparison := anagramChecker.IsAnagram(sixthWord, seventhWord)
	fmt.Printf("Sixth word is %s and seventh word is %s. They are anagrams? %v\n", sixthWord, seventhWord, sixthWordAndseventhWordComparison)
}
