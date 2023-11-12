package main

import "fmt"

var productName [4]string

func main() {

	productName[0] = "Milk"
	productName[1] = "Bread"
	productName[2] = "Eggs"
	productName[3] = "Butter"

	price := [4]float32{2.50, 1.50, 0.75, 1.20}

	fmt.Println(productName)
	fmt.Println(price)
}
