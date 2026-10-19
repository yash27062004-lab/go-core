package main

import (
	"fmt"

	"github.com/yash27062004-lab/go-core/solid-principle/internal/checkout"
	"github.com/yash27062004-lab/go-core/solid-principle/internal/payment"
)

func main() {
	// POLYMORPHISM:
	// Different payment types share one common interface.
	methods := []checkout.PaymentMethod{
		payment.Card{LastFour: "4242"},
		payment.Wallet{UserID: "U-101"},
	}

	// Use the same checkout workflow for every payment method.
	for _, method := range methods {
		order := checkout.New(method)

		if err := order.PayOrder(1200); err != nil {
			fmt.Println("Payment failed:", err)
			continue
		}
	}

	// Demonstrate validation before payment processing.
	order := checkout.New(payment.Card{LastFour: "4242"})

	if err := order.PayOrder(0); err != nil {
		fmt.Println("Expected error:", err)
	}
}
