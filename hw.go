package main

import "fmt"

func main() {
	product := "iPhone 15 Pro"
	brend := "Apple"
	price := 1290000
	isAvailble := true
	month := 12
	discount := price / month
	fmt.Printf("Product: %s\nBrand: %s\nPrice: %d\nAvailability: %t\nDiscount: %d\n", product, brend, price, isAvailble, discount)
}
