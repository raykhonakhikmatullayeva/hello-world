package main

func Mask(number string) string {
	return number[0:4] + " **** **** " + number[12:]
}
