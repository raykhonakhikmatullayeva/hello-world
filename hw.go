package main

import (
	"fmt"
)

func main() {
	name := "iPhone 15 Pro"
	brand := "Apple"
	price := int64(12_990_000)
	inStock := true
	months := 12
	monthlyPayment := price / int64(months)

	fmt.Println("===== Alifshop =====")
	fmt.Printf("Товар: %s\n", name)
	fmt.Printf("Бренд: %s\n", brand)
	fmt.Printf("Цена: %d сум\n", price)
	fmt.Printf("В наличии: %t\n", inStock)
	fmt.Printf("Рассрочка: %d мес → %d сум/мес\n", months, monthlyPayment)
	fmt.Println("====================")
}
