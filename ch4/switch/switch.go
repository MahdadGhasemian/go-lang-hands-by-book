package main

import "fmt"

func simple_switch() {
	words := []string{"a", "cow", "smile", "gopher",
		"octopus", "anthropologist"}
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word!")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "is exactly the right length:", wordLen)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is a long word!")
		}
	}
}

func blank_switch() {
	words := []string{"a", "cow", "smile", "gopher"}
	for _, word := range words {
		switch wordLen := len(word); {
		case wordLen <= 4:
			fmt.Println(word, "is a short word!")
		case wordLen == 5:
			fmt.Println(word, "is exactly the right length:", wordLen)
		case wordLen >= 6 && wordLen <= 9:
			fmt.Println(word, "is a medium length word!")
		default:
			fmt.Println(word, "is a long word!")
		}
	}
}

func main() {
	simple_switch()
	blank_switch()
}
