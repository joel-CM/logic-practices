package main

import "strings"

func DuplicateEncode(word string) string {
	duplicates := map[string]int{}
	result := ""

	// find repeated substrings
	for _, s := range word {
		duplicates[strings.ToLower(string(s))] += 1
	}

	//search in "duplicates" if the substring is repeated
	for _, s := range word {
		if duplicates[strings.ToLower(string(s))] > 1 {
			result += ")"
		} else {
			result += "("
		}
	}

	return result
}
