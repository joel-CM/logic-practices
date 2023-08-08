package main

/*
The idea is to have two pointers and for one to go through the string
until it finds an empty character or until it reaches the end,
then grab the string that is between the two pointers

	and reverse it, and then only add the number of spaces starting
	 from the string. advancing pointer
*/
func ReverseWords(str string) string {
	r := ""
	i := 0 // index
	j := 0 //pointer 1
	k := 0 // pointer 2
	for len(str) > i {
		letter := string(str[k]) // current letter
		// pointer 2 moves forward until it finds an empty character
		// or until it reaches the last character
		if letter == " " || len(str)-1 == k {
			// if the current letter is empty or is the last letter
			if len(str)-1 == k {
				// is the last letter
				r += reverseString(str[j:])
				break
			} else {
				r += reverseString(str[j:k])
			}
			j = k // start searching from position k
			whiteSpacesAdded := addWhiteSpaces(str, &r, k)
			k += whiteSpacesAdded
			j += whiteSpacesAdded
		} else {
			// if it is an empty character
			// or the last character, we keep moving until we find one
			k++
		}

		i++
	}

	return r
}

// returns a new string inverted
func reverseString(str string) string {
	r := ""
	for i := len(str) - 1; i >= 0; i-- {
		r += string(str[i])
	}
	return r
}

// returns how many whitespaces there are from the first one found
//
//	until a non-whitespace character is found
func addWhiteSpaces(from string, to *string, start int) (whiteSpaces int) {
	whiteSpaces = 0
	index := start
	for string(from[index]) == " " && len(from)-1 != index {
		*to += " "
		whiteSpaces++
		index++
	}
	return whiteSpaces
}
