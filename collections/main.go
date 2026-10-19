package main

import "fmt"

/*
PATTERN: COLLECTION → LOOKUP → STATE UPDATE

Slice  = Keep an ordered collection.
Map    = Find data using a key.
Pointer = Modify the original value.
*/

type Product struct {
	Name  string
	Stock int
}

func main() {
	// SLICE: Preserve the order in which products are displayed.
	productOrder := []string{"KB-01", "MS-01"}

	// MAP: Associate each product ID with its product.
	// Pointers let us update the stored products directly.
	products := map[string]*Product{
		"KB-01": &Product{Name: "Keyboard", Stock: 10},
		"MS-01": &Product{Name: "Mouse", Stock: 20},
	}

	// LOOKUP: Find a product using its unique key.
	product, exists := products["KB-01"]
	if !exists {
		fmt.Println("Product not found")
		return
	}

	fmt.Printf("Before sale: %s, stock = %d\n", product.Name, product.Stock)

	// STATE UPDATE: Change the original product's stock.
	product.Stock -= 2

	// DISPLAY: Use the slice to maintain a predictable display order.
	fmt.Println("\nCurrent inventory:")

	for _, id := range productOrder {
		item := products[id]
		fmt.Printf("%s: %d units\n", item.Name, item.Stock)
	}
}
