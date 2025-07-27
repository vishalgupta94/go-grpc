package main

import "fmt"

func printArray(arr []int) {
	for i := range arr {
		fmt.Printf("%d ", (arr)[i])
	}
}

func main() {
	sl := []int{1, 2, 34, 45}
	printArray(sl)
	fmt.Println("Golang")
}
