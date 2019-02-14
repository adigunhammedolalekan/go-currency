## go-currency
Human readable money format

This is a tiny library to format number into any currency.

* Very tiny
* Supports all earth's currency

#### How it works

* Grab the package
`go get github.com/adigunhammedolalekan/go-currency`

* Use it
```Go

    formattedStringNaira := currency.Format(currency.NGR, 2560) // outputs: ₦2,560

    formattedStringPounds := currency.Format(currency.EUR, 300) // outputs: £2,560

    formattedStringGhanaCedis := currency.Format(currency.GHS, 2340.90) // outputs: ¢2,340.9
```