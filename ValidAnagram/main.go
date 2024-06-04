package main

import "fmt"

func isAnagram(firstString string, secoundString string) bool {
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
	firstAndSecoundComparison := isAnagram(firstWord, secoundWord)
	fmt.Printf("First word is %s and secound word is %s. They are anagrams? %v\n", firstWord, secoundWord, firstAndSecoundComparison)

	thirdWord := "rat"
	fourthWord := "car"
	thirdAndFourthComparison := isAnagram(thirdWord, fourthWord)
	fmt.Printf("Third word is %s and fourth word is %s. They are anagrams? %v\n", thirdWord, fourthWord, thirdAndFourthComparison)

}
