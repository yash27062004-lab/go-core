package main

import "fmt"

// PROCESS: Calculate the total of all product prices.
func calculateTotal(prices []int) int {
	total := 0

	// Process every price and accumulate the result.
	for _, price := range prices {
		total += price
	}

	return total
}

func main() {
	// INPUT: Store the prices of multiple products.
	prices := []int{250, 499, 150, 300}

	// PROCESS: Call the function to calculate the total.
	total := calculateTotal(prices)

	// OUTPUT: Display the calculated result.
	fmt.Printf("Total price: $%d\n", total)

	// DECISION: Choose behavior based on the total.
	if total >= 1000 {
		fmt.Println("Category: High-value order")
	} else {
		fmt.Println("Category: Regular order")
	}
}
