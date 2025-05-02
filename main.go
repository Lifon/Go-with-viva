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
	if isLeapYear() {
		fmt.Println("Leap Year")
	} else {
		fmt.Println("Not a Leap Year")
	}

}

func isLeapYear() bool {
	var year int
	fmt.Println("Enter a year:")
	fmt.Scanln(&year)
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				return true
			}
			return false
		}
		return true
	}
	return false
}
