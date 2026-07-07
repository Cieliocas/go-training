package main

import "fmt"

func main() {
	fmt.Println("Bank Account Simulation")

	// User Story 1: open an account and check its balance.
	alice, err := NewAccount("Alice", 100)
	if err != nil {
		fmt.Println("failed to open account:", err)
		return
	}
	fmt.Printf("Opened account for %s with balance %d\n", alice.Owner(), alice.Balance())
}
