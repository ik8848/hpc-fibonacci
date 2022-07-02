package main

import (
	"errors"
	"fmt"
	"math/big"
	"os"
)

func main() {

	input := new(big.Int)
	fmt.Print("Enter n'th term to calculate fibonacci for: ")
	_, err := fmt.Scanln(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error processing input: ", err)
		os.Exit(1)
	}
	val, err := fibonacci(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error calculating fibonacci: ", err)
		os.Exit(1)
	}
	fmt.Println(val)

}

func fibonacci(input *big.Int) (res *big.Int, err error) {
	first, second := big.NewInt(0), big.NewInt(1)

	if input.Cmp(big.NewInt(0)) < 0 {
		return nil, errors.New("invalid input number")
	}
	// if input is 0, return 0
	if input.Cmp(big.NewInt(0)) == 0 {
		return first, nil
	}
	// if input is 1 or 2, return 1
	if input.Cmp(big.NewInt(1)) == 0 || input.Cmp(big.NewInt(2)) == 0 {
		return second, nil
	}
	for i := 2; input.Cmp(big.NewInt(int64(i))) >= 0; i++ {
		sum := big.NewInt(0)
		sum.Add(first, second)
		first = second
		second = sum
	}
	return second, nil
}
