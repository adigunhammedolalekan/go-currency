package currency

import "testing"

func TestComma(t *testing.T) {

	value := 123456.78
	wants := "123,456.78"

	result := Comma(value)

	if wants != result {
		t.Fatalf("Expected %s, %s gotten", wants, result)
	}

	value1 := 123456789.99
	wants2 := "123,456,789.99"

	result2 := Comma(value1)
	if wants2 != result2 {
		t.Fatalf("Expected %s, %s gotten", wants2, result2)
	}
}

func TestFormat(t *testing.T) {

	value := 123456
	expected := "#123,456"

	result := Format(NGR, float64(value))
	if expected != result {
		t.Fatalf("Expected %s, gotten %s", expected, result)
	}
}
