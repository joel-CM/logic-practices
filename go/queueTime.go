package main

import "sort"

func QueueTime(customers []int, n int) int {
	tills := make([]int, n) // initialize an array with length n

	for _, customerTime := range customers {
		sort.Ints(tills)         // order the tills so that they remain from least to greatest time to finish
		tills[0] += customerTime // We add the client to the till that will finish faster
	}

	sort.Ints(tills)

	return tills[len(tills)-1]
}
