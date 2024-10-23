package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Println("Enter revenue")
	fmt.Scan(&revenue)
	fmt.Println("Enter expenses")
	fmt.Scan(&expenses)
	fmt.Println("Enter tax rate")
	fmt.Scan(&taxRate)

	earningsBeforeTax := revenue - expenses
	earningsAfterTax := earningsBeforeTax - (1 - taxRate/100)
	ratio := earningsBeforeTax / earningsAfterTax

	fmt.Println("EBT")
	fmt.Println(earningsBeforeTax)

	fmt.Println("Profit")
	fmt.Println(earningsAfterTax)

	fmt.Println("Ratio:")
	fmt.Println(ratio)
}
