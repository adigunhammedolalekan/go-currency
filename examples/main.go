package main

import (
	"fmt"
	"github.com/adigunhammedolalekan/go-currency"
)

func main() {

	m := currency.Format(currency.NGN, 10000000000000)
	fmt.Println(m)
}
