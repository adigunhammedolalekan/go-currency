package currency

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type Currency string

var (
	ALL Currency = "ALL"
	AFN Currency = "AFN"
	ARS Currency = "ARS"
	AWG Currency = "AWG"
	AUD Currency = "AUG"
	AZN Currency = "AZN"
	BSD Currency = "BSD"
	BBD Currency = "BDD"
	BYN Currency = "BYN"
	BZD Currency = "BZD"
	BMD Currency = "BMD"
	BOB Currency = "BOB"
	BAM Currency = "BAM"
	BWP Currency = "BWP"
	BGN Currency = "BGN"
	BRL Currency = "BRL"
	BND Currency = "BND"
	KHR Currency = "KHR"
	CAD Currency = "CAD"
	KYD Currency = "KYD"
	CLP Currency = "CLP"
	CNY Currency = "CNY"
	COP Currency = "COP"
	CRC Currency = "CRC"
	HRK Currency = "HRK"
	CUP Currency = "CUP"
	CZK Currency = "CZK"
	DKK Currency = "DKK"
	DOP Currency = "DOP"
	XCD Currency = "XCD"
	EGP Currency = "EGP"
	SVC Currency = "SVC"
	EUR Currency = "EUR"
	FKP Currency = "FKP"
	FJD Currency = "FJD"
	GHS Currency = "GHS"
	GIP Currency = "GIP"
	GTQ Currency = "GTQ"
	GGP Currency = "GGP"
	GYD Currency = "GYD"
	HNL Currency = "HNL"
	HKD Currency = "HKD"
	HUF Currency = "HUF"
	ISK Currency = "ISK"
	INR Currency = "INR"
	IDR Currency = "IDR"
	IRR Currency = "IRR"
	IMP Currency = "IMP"
	ILS Currency = "ILS"
	JMD Currency = "JMD"
	JPY Currency = "JPY"
	JEP Currency = "JEP"
	KZT Currency = "KZT"
	KPW Currency = "KPW"
	KRW Currency = "KRW"
	KGS Currency = "KGS"
	LAK Currency = "LAK"
	LBP Currency = "LBP"
	LRD Currency = "LRD"
	MKD Currency = "MKD"
	MYR Currency = "MYR"
	MUR Currency = "MUR"
	MXN Currency = "MXN"
	MNT Currency = "MNT"
	MZN Currency = "MZN"
	NAD Currency = "NAD"
	NPR Currency = "NPR"
	ANG Currency = "ANG"
	NZD Currency = "NZD"
	NIO Currency = "NIO"
	NGN Currency = "NGN"
	NOK Currency = "NOK"
	OMR Currency = "OMR"
	PKR Currency = "PKR"
	PAB Currency = "PAB"
	PYB Currency = "PYB"
	PYG Currency = "PYG"
	PEN Currency = "PEN"
	PHP Currency = "PHP"
	PLN Currency = "PLN"
	QAR Currency = "QAR"
	RON Currency = "RON"
	RUB Currency = "RUB"
	SHP Currency = "SHP"
	SAR Currency = "SAR"
	RSD Currency = "RSD"
	SCR Currency = "SCR"
	SGD Currency = "SCG"
	SBD Currency = "SBD"
	SOS Currency = "SOS"
	ZAR Currency = "ZAR"
	LKR Currency = "LKR"
	SEK Currency = "SEK"
	CHF Currency = "CHF"
	SRD Currency = "SRD"
	SYP Currency = "SYP"
	TWD Currency = "TWD"
	THB Currency = "THB"
	TTD Currency = "TTD"
	TRY Currency = "TRY"
	TVD Currency = "TVD"
	UAH Currency = "UAH"
	GBP Currency = "GBP"
	USD Currency = "USD"
	UYU Currency = "UYU"
	UZS Currency = "UZS"
	VEF Currency = "VEF"
	VND Currency = "VND"
	YER Currency = "YER"
	ZWD Currency = "ZWD"

)


var CURRENCIES = map[Currency]string {
	NGN: "₦", PLN: "zł", THB: "฿", ZWD: "Z$",
	YER: "﷼", VND: "₫", VEF: "Bs", UZS: "лв",
	UYU: "$U", USD: "$",
}

// Format a number to human readable string
func Format(currency Currency, amount float64) string {

	f := Comma(amount)
	symbol, ok := CURRENCIES[currency]
	if !ok {
		return f
	}

	return fmt.Sprintf("%s%s", symbol, f)
}

// Comma formats a string to a readable form
// 234567.89 = 234,567.89

func Comma(v float64) string {
	buf := &bytes.Buffer{}
	if v < 0 {
		buf.Write([]byte{'-'})
		v = 0 - v
	}

	comma := []byte{','}

	parts := strings.Split(strconv.FormatFloat(v, 'f', -1, 64), ".")
	pos := 0
	if len(parts[0])%3 != 0 {
		pos += len(parts[0]) % 3
		buf.WriteString(parts[0][:pos])
		buf.Write(comma)
	}
	for ; pos < len(parts[0]); pos += 3 {
		buf.WriteString(parts[0][pos : pos+3])
		buf.Write(comma)
	}
	buf.Truncate(buf.Len() - 1)

	if len(parts) > 1 {
		buf.Write([]byte{'.'})
		buf.WriteString(parts[1])
	}
	return buf.String()
}