package main

import (
	"strconv"
)

/*
		   Digital root is the recursive sum of all the digits in a number.

		   Given n, take the sum of the digits of n. If that value has more than one digit, continue reducing in this way until a single-digit number is produced. The input will be a non-negative integer.

		   Examples:
		   	16  -->  1 + 6 = 7
	   		942  -->  9 + 4 + 2 = 15  -->  1 + 5 = 6
			132189  -->  1 + 3 + 2 + 1 + 8 + 9 = 24  -->  2 + 4 = 6
			493193  -->  4 + 9 + 3 + 1 + 9 + 3 = 29  -->  2 + 9 = 11  -->  1 + 1 = 2

*/

func DigitalRoot(n int) int {
	string_from_n := strconv.Itoa(n)

	if len(string_from_n) == 1 {
		return n
	}

	sum := 0
	for _, s := range string_from_n {
		int_from_s, _ := strconv.Atoi(string(s))
		sum += int_from_s
	}

	if len(strconv.Itoa(sum)) > 1 {
		return DigitalRoot(sum)
	}

	return sum
}
