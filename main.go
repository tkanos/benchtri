package main

import "fmt"

func main() {
	fmt.Println(Iterative(5))
	fmt.Println(Recursive(5))
	fmt.Println(TailRecursive(5, 0))
}

func Iterative(x int) int {
	total := 0
	for i := x; i > 0; i-- {
		total = total + i
	}

	return total
}

func Recursive(x int) int {
	if x == 1 {
		return x
	}

	return x + Recursive(x-1)
}

func TailRecursive(x int, total int) int {
	if x != 0 {
		return TailRecursive(x-1, total+x)
	}

	return total
}
