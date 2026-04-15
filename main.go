package main

import "fmt"

func main() {
	sender := "Sherali"
	receiver := "Alisher"
	amount := 500000

	commission := amount / 100
	total := amount + commission
	fmt.Println("============================")
	fmt.Printf("Sender: %s\n", sender)
	fmt.Printf("Receiver: %s\n", receiver)
	fmt.Printf("Amount: %d sum\n", commission)
	fmt.Printf("Commission: %d sum\n", commission)
	fmt.Printf("Total: %d sum\n")
	fmt.Println("=================")
}
