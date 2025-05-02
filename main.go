package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	nums := []int{2, 3, 4}
	sum := 0

	for _, num := range nums {
		sum += num
	}
	fmt.Println("Sum:", sum)
	fmt.Println("length of sum: ", len(nums))
}
