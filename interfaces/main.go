package main

import (
	"errors"
	"fmt"
)

/*
PATTERN: ONE INTERFACE → MULTIPLE IMPLEMENTATIONS

Define what a component must do, without tying the caller
to a specific implementation.
*/

// INTERFACE: Every notifier must provide a Send method.
type Notifier interface {
	Send(message string) error
}

// IMPLEMENTATION 1: Send notifications through email.
type EmailNotifier struct {
	Address string
}

func (e EmailNotifier) Send(message string) error {
	if message == "" {
		return errors.New("message cannot be empty")
	}

	fmt.Printf("Email to %s: %s\n", e.Address, message)
	return nil
}

// IMPLEMENTATION 2: Send notifications through SMS.
type SMSNotifier struct {
	Number string
}

func (s SMSNotifier) Send(message string) error {
	if message == "" {
		return errors.New("message cannot be empty")
	}

	fmt.Printf("SMS to %s: %s\n", s.Number, message)
	return nil
}

// POLYMORPHISM: Accept any type that satisfies Notifier.
func sendNotification(notifier Notifier, message string) error {
	return notifier.Send(message)
}

func main() {
	// Different implementations can be used through the same interface.
	notifiers := []Notifier{
		EmailNotifier{Address: "admin@example.com"},
		SMSNotifier{Number: "+91-9876543210"},
	}

	for _, notifier := range notifiers {
		// ERROR HANDLING: Check and handle returned errors.
		if err := sendNotification(notifier, "Deployment completed"); err != nil {
			fmt.Println("Notification failed:", err)
			continue
		}
	}

	// Demonstrate how an implementation reports invalid input.
	if err := sendNotification(notifiers[0], ""); err != nil {
		fmt.Println("Expected error:", err)
	}
}
