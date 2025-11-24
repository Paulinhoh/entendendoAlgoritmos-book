// busca binaria  O(log(n))
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Binary shearch...")

	// lista deve estar ordenada
	list := []int{2, 3, 22, 5, 7, 10, 13, 15, 19, 23, 30}
	sort.Ints(list)

	position := binarySearch(list, 22)

	fmt.Println(position)
}

func binarySearch(list []int, target int) int {
	start := 0
	end := len(list) - 1

	for start <= end {
		middle := (end + start) / 2
		attempt := list[middle]

		if attempt == target {
			return middle
		}
		if attempt > target {
			end = middle - 1
		}
		if attempt < target {
			start = middle + 1
		}
	}

	return -1
}
