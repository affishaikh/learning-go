package main

import "fmt"

type Note struct {
	amount   int
	quantity int
}

func withdraw(amount int) {
	result := []Note{}
	availableAmount := [4]Note{
		{2000, 100},
		{500, 100},
		{100, 100},
		{50, 100},
	}
	var remainingAmount = amount

	for i := 0; i < 4 && remainingAmount > 100; i++ {
		currentNote := availableAmount[i].amount
		result = append(result, Note{currentNote, remainingAmount / currentNote})
		remainingAmount = remainingAmount % currentNote
	}

	fmt.Println(result)
}

func main() {

	withdraw(45001)
}
