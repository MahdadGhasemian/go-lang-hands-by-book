package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	y := make([]int, 4)
	num := copy(y, x)
	fmt.Println(y, num)

	z := make([]int, 2)
	copy(z, x[1:3])
	fmt.Println(z)
}
