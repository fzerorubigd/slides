package main

import "fmt"

// Simple START OMIT
func minMax(items []int) (min int, max int) {
	if len(items) == 0 {
		panic("invalid usage")
	}

	min = items[0]
	max = items[0]
	for _, i := range items {
		if max < i {
			max = i
		}
		if min > i {
			min = i
		}
	}

	return
}

// Simple END OMIT

func main() {
	fmt.Println(minMax([]int{10, 20, 30}))
}
