package main

import (
	"fmt"
	"sort"
)

func bigSorting(unsorted []string) []string {
	sort.Slice(unsorted, func(i, j int) bool {
		if len(unsorted[i]) != len(unsorted[j]) {
			return len(unsorted[i]) < len(unsorted[j])
		} else {
			return unsorted[i] < unsorted[j]
		}
	})

	return unsorted
}

func main() {
	fmt.Println(bigSorting([]string{"1",
		"2",
		"100",
		"12303479849857341718340192371",
		"3084193741082937",
		"3084193741082938",
		"111",
		"200",
	}))
}
