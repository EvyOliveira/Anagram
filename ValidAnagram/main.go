package main

import (
	"fmt"
	"strings"
)

func IsAnagram(firstString, secoundString string) bool {
	if firstString == "" || secoundString == "" {
		return false
	}

	firstString = strings.ToLower(firstString)
	secoundString = strings.ToLower(secoundString)

	charactersMap := make(map[rune]int)
	for _, character := range firstString {
		charactersMap[character]++
	}

	for _, character := range secoundString {
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

func main() {
	firstWord := "anagram"
	secoundWord := "nagaram"
	firstAndSecoundComparison := IsAnagram(firstWord, secoundWord)
	fmt.Printf("First word is %s and secound word is %s. They are anagrams? %v\n", firstWord, secoundWord, firstAndSecoundComparison)

	thirdWord := "rat"
	fourthWord := "car"
	thirdAndFourthComparison := IsAnagram(thirdWord, fourthWord)
	fmt.Printf("Third word is %s and fourth word is %s. They are anagrams? %v\n", firstWord, secoundWord, thirdAndFourthComparison)

	emptyString := ""
	fifthWord := "test"
	emptyStringAndFifthWordComparison := IsAnagram(emptyString, fifthWord)
	fmt.Printf("Empty string is %s and fifth word is %s. They are anagrams? %v\n", emptyString, fifthWord, emptyStringAndFifthWordComparison)

	sixthWord := "iRaceMa"
	seventhWord := "aMerIca"
	sixthWordAndseventhWordComparison := IsAnagram(sixthWord, seventhWord)
	fmt.Printf("Sixth word is %s and seventh word is %s. They are anagrams? %v\n", sixthWord, seventhWord, sixthWordAndseventhWordComparison)
}
