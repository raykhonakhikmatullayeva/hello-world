package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Number of your card: ")
	card := flag.String("card", "", "Card number")
	flag.Parse()
	if *card == "" {
		fmt.Println("Error: please provide a card number using -card")
		return
	}
	normalized := strings.ReplaceAll(*card, " ", "")
	normalized = strings.ReplaceAll(normalized, "-", "")
	if len(normalized) != 16 {
		fmt.Println("Wrong card number!")
		return
	}
	fmt.Println("Card number: ", normalized[:16])
	fmt.Printf("Length: %d", len(normalized))
	masked := normalized[:4] + "********" + normalized[12:]
	fmt.Println("Card: ", masked)
}
func allDigits(s string) bool {
	for _, ch := range result {
		if ch < '0' || ch > '9' {
			fmt.Println("Wrong card number!")}


