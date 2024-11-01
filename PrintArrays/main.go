package main

import "fmt"

func main() {
	var arr = []int{21, 30, 62, 5, 12, 51, 96} // aqui eu declaro o meu array com 6 posições
	printArray(arr)
}

func printArray(arr []int) { // função recebe um array de inteiros com tamanho indefinido
	for _, value := range arr { // se 0 for menor que o tamanho do meu array, percorra sobre i até ele chegar no ultimo index da lista
		fmt.Printf("%d\t", value) // printando o array
	}
}
