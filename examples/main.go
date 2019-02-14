package main

import (
	"fmt"
	"github.com/adigunhammedolalekan/go-currency"
)

func main() {

	m := currency.Format(currency.NGN, 100000)
	fmt.Println(m) // ₦100,000

	ghs := currency.Format(currency.GHS, 2345.90)
	fmt.Println(ghs)
}
