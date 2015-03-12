package panitizer_test

import (
	"testing"

	"github.com/NickPresta/panitizer"
)

func TestLuhnCheck(t *testing.T) {
	creditCardTests := []struct {
		in       string
		expected bool
	}{
		{"4242 4242 4242 4242", true},
		{"49927398716", true},
		{"49927398717", false},
		{"1234567812345678", false},
		{"1234567812345670", true},
		{"378282246310005", true},
		{"371449635398431", true},
		{"378734493671000", true},
		{"5610591081018250", true},
		{"30569309025904", true},
		{"38520000023237", true},
		{"6011111111111117", true},
		{"6011000990139424", true},
		{"3530111333300000", true},
		{"3566002020360505", true},
		{"5555555555554444", true},
		{"5105105105105100", true},
		{"4111111111111111", true},
		{"4012888888881881", true},
		{"4222222222222", true},
		{"foo", false},
	}
	for _, tt := range creditCardTests {
		if actual := panitizer.PassesLuhnCheck(tt.in); tt.expected != actual {
			t.Errorf("Expected %v given %q but got %v", tt.expected, tt.in, actual)
		}
	}
}
