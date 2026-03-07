package main

import "fmt"

func main() {
	totalWins := map[string]int{}
	totalWins["Alice"] = 10
	totalWins["Bob"] = 15
	fmt.Println((totalWins["Bob"]))
	fmt.Println(totalWins["Kittens"])
	totalWins["Kittens"]++
	totalWins["Lions"] = 3
	fmt.Println(totalWins["Kittens"])
	fmt.Println(totalWins["Lions"])

	fmt.Println(totalWins)

	m := map[string]int{"Alice": 10, "Bob": 15}
	v, ok := m["Alice"]
	fmt.Println(v, ok)
	v, ok = m["Kittens"]
	fmt.Println(v, ok)

	delete(totalWins, "Kittens")
	fmt.Println(totalWins["Kittens"])

	clear(totalWins)
	fmt.Println(totalWins)
}
