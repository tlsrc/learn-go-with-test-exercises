package roman

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic int) string {
	var res strings.Builder

	for _, num := range allRomanNumerals {
		for arabic >= num.Value {
			res.WriteString(num.Symbol)
			arabic -= num.Value
		}
	}
	return res.String()
}

func ConvertToArabic(roman string) int {
	result := 0

	for range roman {
		result++
	}

	return result
}
