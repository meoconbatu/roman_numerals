package romannumerals

import "errors"

const testVersion = 3

type romanNumeral struct {
	arabic int
	roman  string
}

var romanNumerals = []romanNumeral{
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

func ToRomanNumeral(arabic int) (string, error) {
	if arabic <= 0 || arabic >= 4000 {
		return " ", errors.New("Error")
	}
	for i := 0; i < len(romanNumerals); i++ {
		a := romanNumerals[i].arabic
		r := romanNumerals[i].roman
		if a == arabic {
			return r, nil
		} else if arabic > a {
			t, err := ToRomanNumeral(arabic - a)
			return r + t, err
		}
	}
	return " ", nil
}
