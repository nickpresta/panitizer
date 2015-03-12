package panitizer

import "regexp"

// This is the expanded regexp below
// (?:4[0-9]{12}(?:[0-9]{3})?            # Visa
//   |  5[1-5][0-9]{14}                  # MasterCard
//   |  3[47][0-9]{13}                   # American Express
//   |  3(?:0[0-5]|[68][0-9])[0-9]{11}   # Diners Club
//   |  6(?:011|5[0-9]{2})[0-9]{12}      # Discover
//   |  (?:2131|1800|35\d{3})\d{11}      # JCB
// )
// This regexp was taken from http://www.regular-expressions.info/creditcard.html but is modified
// to support a variable number of spaces and/or dashes between groups.
var creditCardRegexp = regexp.MustCompile(`(?:4[ -]*([0-9][ -]*){11}[0-9](?:[ -]*([0-9][ -]*){2}[0-9])?` +
	`|5[ -]*[1-5][ -]*([0-9][ -]*){13}[0-9]` +
	`|3[ -]*[47][ -]*([0-9][ -]*){12}[0-9]` +
	`|3[ -]*(?:0[ -]*[0-5][ -]*|[68][ -]*[0-9][ -]*)([0-9][ -]*){10}[0-9]` +
	`|6[ -]*(?:0[ -]*1[ -]*1[ -]*|5[ -]*([0-9][ -]*){2})([0-9][ -]*){11}[0-9]` +
	`|(?:2[ -]*1[ -]*3[ -]*1[ -]*|1[ -]*8[ -]*0[ -]*0[ -]*|3[ -]*5[ -]*(\d[ -]*){3})(\d[ -]*){10}\d)`)

// Replace returns a copy of pan, replacing Personal Account Numbers with '*' characters.
func Replace(pan string) string {
	return ReplaceWithSymbol(pan, "*")
}

// ReplaceWithSymbol returns a copy of pan, replacing Personal Account Numbers with symbol characters.
func ReplaceWithSymbol(pan, symbol string) string {
	return creditCardRegexp.ReplaceAllStringFunc(pan, func(match string) string {
		if PassesLuhnCheck(match) {
			match = maskCreditCardNumberWithSymbol(match, symbol)
		}
		return match
	})
}

// ContainsPAN returns whether or not pan contains Personal Account Numbers.
func ContainsPAN(pan string) bool {
	return creditCardRegexp.MatchString(pan)
}
