package roman

import (
	"errors"
	"strings"
)

type RomanNumeral struct {
	Value  int
	Symbol string
}

type RomanNumerals []RomanNumeral

var ErrInvalidRomanNumeral = errors.New("Invalid roman numeral")

func (r RomanNumerals) ValueOf(symbols ...byte) (int, error) {
	symbol := string(symbols)
	for _, n := range r {
		if n.Symbol == symbol {
			return n.Value, nil
		}
	}
	return 0, ErrInvalidRomanNumeral
}

var allRomanNumerals = RomanNumerals{
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

	for i := 0; i < len(roman); i++ {
		symbol := roman[i]
		symbolValue, err := allRomanNumerals.ValueOf(symbol)
		if err != nil {
			panic("Invalid roman numeral")
		}

		atTheEnd := i+1 >= len(roman)
		if atTheEnd {
			result += symbolValue
			break
		}

		nextSymbol := roman[i+1]
		candidate, err := allRomanNumerals.ValueOf(symbol, nextSymbol)
		if err != nil {
			result += symbolValue
		} else {
			result += candidate
			i++
		}
	}
	return result
}
