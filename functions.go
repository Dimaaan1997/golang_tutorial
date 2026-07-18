package main

import "fmt"

func print(numbers []*int) {
	for _, n := range numbers {
		fmt.Printf("%d ", *n)
	}

	fmt.Printf("\n")
}

func mult(numbers []*int, factor int) {
	extPointers := make(map[*int]any)

	for _, ptr := range numbers {
		if _, ok := extPointers[ptr]; !ok {
			extPointers[ptr] = struct{}{}
			*ptr *= factor

		}
	}
}
