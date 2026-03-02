package main

import "fmt"

func main() {
	var x int = 10
	var y float32 = 30.2
	var sum1 float32 = float32(x) + y
	var sum2 int = x + int(y)
	fmt.Println(sum1)
	fmt.Println(sum2)
}
