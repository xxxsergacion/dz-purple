package main

import (
	"fmt"
)

func main() {
	const USDEUR = 0.8539
	const USDRUB = 84.92
	var EURRUB float64 = USDRUB / USDEUR
	fmt.Print(EURRUB)
}
