package propertybasedtests

import (
	"fmt"
	"testing"
)

func TestRomanNumerals(t *testing.T) {
	tests := []struct {
		num           int
		expectedRoman string
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
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d gets converted to %q", test.num, test.expectedRoman), func(t *testing.T) {
			got := ConvertToRoman(test.num)

			if got != test.expectedRoman {
				t.Errorf("want %q, got %q", test.expectedRoman, got)
			}
		})

	}

}
