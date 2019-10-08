package utils

import "fmt"

func printNum(num int) {
	fmt.Println("Current Number:", num)
}

// Add ...
func Add(nums ...int) int {
	total := 0
	for _, v := range nums {
		printNum(v)
		total += v
	}
	return total

}
