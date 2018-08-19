package xlutil

import (
	"math"
	"strings"
)

const (
	Billion  = 1000000000
	Million  = 1000000
	Thousand = 1000
	Hundred  = 100
)

func NumberToWordsEN(number int) string {
	if number == 0 {
		return "zero"
	}

	if number < 0 {
		newNumber := math.Abs(float64(number))
		return "minus " + NumberToWordsEN(int(newNumber))
	}

	words := ""

	if (number / Billion) > 0 {
		words += NumberToWordsEN(number/Billion) + " billion "
		number %= Billion
	}

	if (number / Million) > 0 {
		words += NumberToWordsEN(number/Million) + " million "
		number %= Million
	}

	if (number / Thousand) > 0 {
		words += NumberToWordsEN(number/Thousand) + " thousand "
		number %= Thousand
	}

	if (number / Hundred) > 0 {
		words += NumberToWordsEN(number/Hundred) + " hundred "
		number %= Hundred
	}

	if number > 0 {
		if words != "" {
			words += "and "
		}

		var unitsMap = []string{
			"zero", "one", "two", "three", "four", "five",
			"six", "seven", "eight", "nine", "ten",
			"eleven", "twelve", "thirteen", "fourteen", "fifteen",
			"sixteen", "seventeen", "eighteen", "nineteen",
		}
		var tensMap = []string{
			"zero", "ten", "twenty", "thirty", "forty", "fifty",
			"sixty", "seventy", "eighty", "ninety",
		}

		if number < 20 {
			words += unitsMap[number]
		} else {
			words += tensMap[number/10]
			if (number % 10) > 0 {
				words += "-" + unitsMap[number%10]
			}
		}
	}

	return words
}

func NumberToWordsFA(number int) string {
	if number == 0 {
		return "صفر"
	}

	if number < 0 {
		newNumber := math.Abs(float64(number))
		return "منفی " + NumberToWordsFA(int(newNumber))
	}

	words := ""

	if (number / Billion) > 0 {
		words += NumberToWordsFA(number/Billion) + " میلیارد "
		if number %= Billion; number != 0 {
			words += " و "
		}
	}

	if (number / Million) > 0 {
		words += NumberToWordsFA(number/Million) + " میلیون "
		if number %= Million; number != 0 {
			words += " و "
		}
	}

	if (number / Thousand) > 0 {
		words += NumberToWordsFA(number/Thousand) + " هزار "
		if number %= Thousand; number != 0 {
			words += " و "
		}
	}

	if (number / Hundred) > 0 {
		thirdNumber := number / Hundred
		switch thirdNumber {
		case 1:
			words += " یکصد "
		case 2:
			words += " دویست "
		case 3:
			words += " سیصد "
		case 4:
			words += " چهارصد "
		case 5:
			words += " پانصد "
		case 6:
			words += " ششصد "
		case 7:
			words += " هفتصد "
		case 8:
			words += " هشتصد "
		case 9:
			words += " نهصد "
		default:
			words += NumberToWordsFA(number/Hundred) + " صد "
		}
		number %= Hundred
	}

	if number > 0 {
		if words != "" {
			words += "و "
		}

		var unitsMap = []string{
			"صفر", "یک", "دو", "سه", "چهار", "پنج", "شش", "هفت", "هشت", "نه", "ده",
			"یازده", "دوازده", "سیزده", "چهارده", "پانزده", "شانزده", "هفده", "هجده", "نوزده",
		}
		var tensMap = []string{
			"صفر", "ده", "بیست", "سی", "چهل", "پنچاه", "شصت", "هفتاد", "هشتاد", "نود",
		}

		if number < 20 {
			words += unitsMap[number]
		} else {
			words += tensMap[number/10]
			if (number % 10) > 0 {
				words += " و " + unitsMap[number%10]
			}
		}
	}
	words = strings.Replace(words, "  ", " ", -1)
	return words
}
