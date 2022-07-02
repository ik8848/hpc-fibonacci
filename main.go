package main

import (
	"fmt"
	"os"
)

func main() {

	var input int
	fmt.Print("Enter n'th term to calculate fibonacci for: ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(fibonacci(input))
}

func fibonacci(input int) (res int) {
	first, second := 1, 1

	for i := 2; i < input; i++ {
		res = first + second
		first = second
		second = res
	}
	return res
}
