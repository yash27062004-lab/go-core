package main

import (
	"fmt"

	// The import path must match the module declared in go.mod.
	// This example assumes your module name is "go-core".

	"github.com/yash27062004-lab/go-core/encapsulation/internal/user"
)

func main() {
	// Create a User through the constructor.
	u, err := user.New(
		"U-101",
		"Asha",
		"asha@example.com",
		"+91-9876543210",
	)
	if err != nil {
		fmt.Println("User creation failed:", err)
		return
	}

	// Read the user's information.
	fmt.Println("Before:", u.Summary())

	// Update the email through a controlled method.
	if err := u.ChangeEmail("asha.dev@example.com"); err != nil {
		fmt.Println("Email update failed:", err)
		return
	}

	// Display the updated state.
	fmt.Println("After: ", u.Summary())
}
