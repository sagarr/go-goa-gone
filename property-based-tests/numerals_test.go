package propertybasedtests

import (
	"fmt"
	"log"
	"testing"
	"testing/quick"
)

var tests = []struct {
	Number uint16
	Roman  string
}{
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{4, "IV"},
	{5, "V"},
	{8, "VIII"},
	{9, "IX"},
	{10, "X"},
	{18, "XVIII"},
	{20, "XX"},
	{39, "XXXIX"},
	{40, "XL"},
	{47, "XLVII"},
	{49, "XLIX"},
	{50, "L"},
}

func TestRomanNumerals(t *testing.T) {
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d converted to %q", test.Number, test.Roman), func(t *testing.T) {
			got := ConvertToRoman(test.Number)

			if got != test.Roman {
				t.Errorf("want %q, got %q", test.Roman, got)
			}
		})
	}
}

func TestConvertingRomanToNumber(t *testing.T) {
	for _, test := range tests {
		t.Run(fmt.Sprintf("%s converted to %d", test.Roman, test.Number), func(t *testing.T) {
			got := ConvertToNumber(test.Roman)

			if got != test.Number {
				t.Errorf("want %q, got %d", test.Roman, got)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(number uint16) bool {
		if number > 3999 {
			return true
		}
		log.Println("testing: ", number)
		roman := ConvertToRoman(number)
		numberFromRoman := ConvertToNumber(roman)
		return numberFromRoman == number
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error("failed checks", err)
	}
}
