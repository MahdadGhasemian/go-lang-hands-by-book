package main

import "fmt"

func for_range() {
	evenVals := []int{2, 4, 6, 8, 10}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}
}

func for_range2() {
	evenVals := []int{2, 4, 6, 8, 10}
	for _, v := range evenVals {
		fmt.Println(v)
	}
}

func for_range3() {
	uniqueNames := map[string]bool{
		"Fred":  true,
		"Raul":  true,
		"Wilma": false,
	}
	for k := range uniqueNames {
		fmt.Println(k)
	}
}

func map_iteration_over_varies() {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	for i := 0; i < 3; i++ {
		fmt.Println("Loop", i)
		for k, v := range m {
			fmt.Println(k, v)
		}
	}
}

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 5; {
		fmt.Println(i)
		if i%2 == 0 {
			i++
		} else {
			i += 2
		}
	}

	for i := 0; i < 5; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}
		if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}
		fmt.Println(i)
	}

	for_range()

	for_range2()

	for_range3()

	map_iteration_over_varies()
}
