package main

import "fmt"

func main() {
	fmt.Print(regressiva(5))
}

func regressiva(i int) int {

	if i == 0 {
		return 1
	}

	return i * regressiva(i-1)
}
