package main

import "fmt"

func sum(a, b, c int) int {
	sum := a + b + c
	avg := sum / 3
	return avg
}

func main() {
	fmt.Println(sum(5, 7, 10))
}
