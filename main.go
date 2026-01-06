package main

import (
	"fmt"
	"test_git/division"
	"test_git/multiplication"
	"test_git/subtraction"
	"test_git/sum"
)

func main() {
	fmt.Println(sum.Sum(2, 2))
	fmt.Println(subtraction.Subtraction(2, 2))
	fmt.Println(multiplication.Multiplication(2, 2))
	num, err := division.Division(10, 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(num)
	num, err = division.Division(10, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(num)
}
