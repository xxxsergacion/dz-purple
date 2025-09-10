package main

import (
	"fmt"
)

func main() {
	const USDEUR = 0.8539
	const USDRUB = 84.92
	EURRUB := USDRUB / USDEUR
	fmt.Print(EURRUB)
}
