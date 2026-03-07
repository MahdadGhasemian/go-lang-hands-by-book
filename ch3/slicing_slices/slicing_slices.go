package main

import "fmt"

func main() {
	x := []string{"a", "b", "c", "d"}
	v := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:]

	fmt.Println("x:", x)
	fmt.Println("v:", v)
	fmt.Println("z:", z)
	fmt.Println("d:", d)
	fmt.Println("e:", e)

	x[1] = "v"
	v[0] = "x"
	z[1] = "y"

	fmt.Println("x:", x)
	fmt.Println("v:", v)
	fmt.Println("z:", z)
	fmt.Println("d:", d)
	fmt.Println("e:", e)
}
