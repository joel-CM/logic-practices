package main

import (
	"sort"
)

func FindShort(s string) int {
	lengths := []int{}

	word := ""
	for _, e := range s {
		letter := string(e)
		if letter != " " {
			word += letter
		} else {
			lengths = append(lengths, len(word))
			word = ""
		}
	}
	lengths = append(lengths, len(word))

	sort.Ints(lengths) // sort lengths and return the first element
	return lengths[0]
}
