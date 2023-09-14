package main

func Multiple3And5(number int) int {
	multiples3And5 := []int{}
	for i := 1; i < number; i++ {
		if i%3 == 0 && i%5 == 0 {
			multiples3And5 = append(multiples3And5, i)
			continue
		}
		if i%3 == 0 {
			multiples3And5 = append(multiples3And5, i)
		}
		if i%5 == 0 {
			multiples3And5 = append(multiples3And5, i)
		}
	}
	return sum(multiples3And5)
}

func sum(list []int) int {
	res := 0
	for _, num := range list {
		res += num
	}
	return res
}
