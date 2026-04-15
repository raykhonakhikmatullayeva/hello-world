package main

import "fmt"

func Mask(number string) string {
	return number[0:4] + " **** **** " + number[12:16]
