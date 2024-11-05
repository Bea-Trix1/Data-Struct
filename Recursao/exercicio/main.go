package main

import "fmt"

func Fibonacci(num int) int {
	if num < 2 {
		return 1
	}

	return Fibonacci(num-1) + Fibonacci(num-2)
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Print(Fibonacci(i), " ")
	}

}
