package main

import (
	"fmt"
	"github.com/adigunhammedolalekan/go-currency"
)

func main() {

	m := currency.Format(currency.NGN, 100000)
	fmt.Println(m) // ₦100,000
}
