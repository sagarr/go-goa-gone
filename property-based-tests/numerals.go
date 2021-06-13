package propertybasedtests

import "strings"

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ValueOf(symbols ...byte) uint16 {
	symbol := string(symbols)
	for _, r := range allRomanNumerals {
		if r.Symbol == symbol {
			return r.Value
		}
	}
	return 0
}

func ConvertToRoman(num uint16) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for num >= numeral.Value {
			result.WriteString(numeral.Symbol)
			num -= numeral.Value
		}
	}
	return result.String()
}

func ConvertToNumber(roman string) uint16 {
	var total uint16

	for i := 0; i < len(roman); i++ {
		symbol := roman[i]

		if couldBeSubstractive(i, symbol, roman) {
			nextSymbol := roman[i+1]

			if value := ValueOf(symbol, nextSymbol); value != 0 {
				total += value
				i++
			} else {
				total += ValueOf(symbol)
			}
		} else {
			total += ValueOf(symbol)
		}
	}

	return total

}

func couldBeSubstractive(i int, symbol uint8, roman string) bool {
	isSubstractive := symbol == 'I' || symbol == 'X' || symbol == 'C'
	return isSubstractive && i+1 < len(roman)
}
