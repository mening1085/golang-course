package main

import "fmt"

var product = make(map[string]float32)

func main() {
	fmt.Println("product", product)

	// add new item
	product["milk"] = 2.4
	product["coffee"] = 3.2
	product["tea"] = 1.2
	product["sugar"] = 0.8

	fmt.Println("product", product)
}
