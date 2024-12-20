package main

import (
	"fmt"
)

func main() {
	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	chars := []rune(input)

	// Convert runes to a slice of strings
	var stringSlice []string
	for _, char := range chars {
		stringSlice = append(stringSlice, string(char))
	}

	fmt.Println(stringSlice)
}
