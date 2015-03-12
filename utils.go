package panitizer

import "regexp"

var spacesAndDashesRegexp = regexp.MustCompile(`[ -]*`)
var numberRegexp = regexp.MustCompile(`\d`)

// maskCreditCardNumber will return a masked credit card number with
func maskCreditCardNumberWithSymbol(number, symbol string) string {
	numLen := len(number)
	if numLen < 4 {
		return number
	}
	last4 := numLen - 4
	return numberRegexp.ReplaceAllString(number, symbol)[:last4] + number[last4:]
}

// MaskCreditCardNumber will mask a credit card number keeping all but the last 4 digits, replacing all others with '*' characters
func MaskCreditCardNumber(number string) string {
	return maskCreditCardNumberWithSymbol(number, "*")
}

// PassesLuhnCheck will return true if the passed in credit card number passes the Luhn Algorithm.
func PassesLuhnCheck(number string) bool {
	// Based on http://rosettacode.org/wiki/Luhn_test_of_credit_card_numbers#Go
	normalized := spacesAndDashesRegexp.ReplaceAllString(number, "")

	var mapping = [...]int{0, 2, 4, 6, 8, 1, 3, 5, 7, 9}
	sum := 0
	odd := len(normalized)%2 == 1

	for i, c := range normalized {
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
