package main

import "fmt"

type Anagram interface {
	IsAnagram(string, string) bool
}

type AnagramChecker struct{}

func (a *AnagramChecker) IsAnagram(firstString, secoundString string) bool {
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
	anagramChecker := AnagramChecker{}

	firstString := "anagram"
	secoundString := "nagaram"
	firstAndSecoundComparison := anagramChecker.IsAnagram(firstString, secoundString)
	fmt.Printf("First word is %s and secound word is %s. They are anagrams? %v\n", firstString, secoundString, firstAndSecoundComparison)

	thirdWord := "rat"
	fourthWord := "car"
	thirdAndFourthComparison := anagramChecker.IsAnagram(thirdWord, fourthWord)
	fmt.Printf("Third word is %s and fourth word is %s. They are anagrams? %v\n", thirdWord, fourthWord, thirdAndFourthComparison)
}
