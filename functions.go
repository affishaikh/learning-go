package main

import "fmt"

func average(nums ...int) int {
	total := 0

	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}

	return total/len(nums)
}

// Multiple returns
func multiplicationAddition(nums ...int) (int, int) {
	multiplication := 1
	sum := 0

	for i := 0; i < len(nums); i++ {
		multiplication *= nums[i]
		sum += nums[i]
	}

	return multiplication, sum
}

// Naked returns
func sum(nums ...int) (sum int) {
	sum = 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	return
}

func main() {
	fmt.Println(average(1,2,3,4,5))

	fmt.Println(multiplicationAddition(1,2,3,4,5))

	fmt.Println(sum(1,2,3,4,5))
}
