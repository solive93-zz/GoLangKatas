package main

import "strings"

func Convert(word string) string {
	array := toArray(strings.ToLower(word))
	for i := 0; i < len(array); i++ {
		if isVowel(array[i]) {
			array[i] = "mumy"
		}
	}
	return toString(array)
}

func toArray(word string) []string {
	return strings.Split(word,"")
}

func toString(array []string) string {
	return strings.Join(array, "")
}

func isVowel(char string) bool {
	vowels := [5]string{"a", "e", "i", "o", "u"}
	vowelTable := make(map[string]bool)
	for _, vowel := range vowels {
		vowelTable[vowel] = true
	}
	return vowelTable[char]
}

