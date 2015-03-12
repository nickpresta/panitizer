package panitizer

import "regexp"

var spacesAndDashesRegexp = regexp.MustCompile(`[ -]*`)
var numberRegexp = regexp.MustCompile(`\d`)

func normalize(s string) string {
	return spacesAndDashesRegexp.ReplaceAllString(s, "")
}

// maskCreditCardNumber will return a masked credit card number with
func maskCreditCardNumberWithSymbol(number, symbol string) string {
	numLen := len(number)
	if numLen < 4 {
		return number
	}
	last4 := numLen - 4
	return numberRegexp.ReplaceAllString(number, symbol)[:last4] + number[last4:]
}

// PassesLuhnCheck will return true if the passed in credit card number passes the Luhn Algorithm.
func PassesLuhnCheck(s string) bool {
	// Based on http://rosettacode.org/wiki/Luhn_test_of_credit_card_numbers#Go
	number := normalize(s)

	var mapping = [...]int{0, 2, 4, 6, 8, 1, 3, 5, 7, 9}
	sum := 0
	odd := len(number)%2 == 1

	for i, c := range number {
		if c < '0' || c > '9' { // Check that we only have digits
			return false
		}

		if (i%2 == 1) == odd {
			sum += mapping[c-'0']
		} else {
			sum += int(c - '0')
		}
	}

	return sum%10 == 0
}
