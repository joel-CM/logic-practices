package main

func MoveZeros(arr []int) []int {
	zeros := []int{}
	others := []int{}
	for _, num := range arr {
		if num == 0 {
			zeros = append(zeros, num)
		} else {
			others = append(others, num)
		}
	}
	return append(others, zeros...)
}
