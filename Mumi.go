package main

import "strings"

func Convert(word string) string {
	array := toArray(strings.ToLower(word))
	if !has_more_than_30_per_cent_of_vowels(array){
		return word
	}
	for i := 0; i < len(array); i++ {
		if isVowel(array[i]) {
			array[i] = "mumy"
		}
	}
	return toString(array)
}

func has_more_than_30_per_cent_of_vowels(array []string) bool {
	thirty_per_cent := (len(array) * 30)/100
	vowelCount := 0
	for i := 0; i < len(array); i++ {
		if isVowel(array[i]){
			vowelCount++
		}
	}
	return thirty_per_cent < vowelCount
}

func toFloat(i int) float64 {
	return float64(i)
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

