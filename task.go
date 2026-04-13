package main

import (
	"flag"
	"fmt"
	"strings"
)

func formatAmount(amount string) string {
	var builder strings.Builder
	n := len(amount)
	for i, ch := range amount {
		builder.WriteRune(ch)
		if (n-i-1)%3 == 0 && i != n-1 {
			builder.WriteRune(' ')
		}
	}
	return builder.String()
}

func main() {
	senderFlag := flag.String("sender", "", "имя отправителя")
	amountFlag := flag.String("amount", "", "сумма")
	flag.Parse()

	sender := strings.ToUpper(*senderFlag)
	amount := formatAmount(*amountFlag)

	fmt.Printf("\n  %s сум поступили 🎉\n", amount)
	fmt.Printf("  %s отправил вам деньги\n", sender)
}
