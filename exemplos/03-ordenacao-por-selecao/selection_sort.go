// busca por seleção O(n²)
package main

import "fmt"

func main() {
	data := []int{64, 25, 12, 22, 11}
	fmt.Println("Slice desordenado:", data)

	selectionSort(data)
	fmt.Println("\nSlice ordenado:", data)
}

func selectionSort(array []int) {
	tamArray := len(array)
	for i := 0; i < tamArray; i++ {
		menorIndice := i
		for j := i + 1; j < tamArray; j++ {
			if array[menorIndice] > array[j] {
				menorIndice = j
			}
		}
		array[menorIndice], array[i] = array[i], array[menorIndice]
	}
}
