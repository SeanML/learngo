package main

import "fmt"

func genNumbers() []int {
	numbers := []int{}
	for i := 0; i <= 10; i++ {
		numbers = append(numbers, i)
	}

	return numbers
}

func printEvenOrOdd(n []int) {
	for _, v := range n {
		if v%2 == 0 {
			fmt.Println(v, "is even")
		} else {
			fmt.Println(v, "is odd")
		}
	}
}
