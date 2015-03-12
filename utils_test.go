package panitizer

import "testing"

func TestMaskCreditCardNumber(t *testing.T) {
	creditCardTests := []struct {
		in       string
		expected string
	}{
		{"4242 4242 4242 4242", "**** **** **** 4242"},
		{"4242-4242-4242-4242", "****-****-****-4242"},
		{"4242424242424242", "************4242"},
		{"foo", "foo"},
		{"morethanfour", "morethanfour"},
		{"four", "four"},
	}
	for _, tt := range creditCardTests {
		if actual := MaskCreditCardNumber(tt.in); tt.expected != actual {
			t.Errorf("Expected '%v' given '%v' but got '%v'", tt.expected, tt.in, actual)
		}
	}
}

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
		if actual := PassesLuhnCheck(tt.in); tt.expected != actual {
			t.Errorf("Expected '%v' given '%v' but got '%v'", tt.expected, tt.in, actual)
		}
	}
}
