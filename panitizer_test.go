package panitizer_test

import (
	"testing"

	"github.com/NickPresta/panitizer"
)

func TestMaskCreditCardNumber(t *testing.T) {
	creditCardTests := []struct {
		in       string
		expected string
	}{
		{"4242 4242 4242 4242", "**** **** **** 4242"},
		{"4242   4242   4242   4242", "****   ****   ****   4242"},
		{"4242-4242-4242-4242", "****-****-****-4242"},
		{"42-42 42-42 42-42 42-42", "**-** **-** **-** *2-42"},
		{"4242424242424242", "************4242"},
		{"4242.4242.4242.4242", "4242.4242.4242.4242"},
		{"4242x4242x4242x4242", "4242x4242x4242x4242"},
		{"4242 4242 4242 4242\n\nhello\n\nhey", "**** **** **** 4242\n\nhello\n\nhey"},
		{"4242 4242 4242 4242 4242 4242", "**** **** **** 4242 4242 4242"},
		{"foo", "foo"},
		{"morethanfour", "morethanfour"},
		{"four", "four"},
	}
	for _, tt := range creditCardTests {
		if actual := panitizer.Replace(tt.in); tt.expected != actual {
			t.Errorf("Expected %q given %q but got %q", tt.expected, tt.in, actual)
		}
	}
}

func TestContainsPAN(t *testing.T) {
	creditCardTests := []struct {
		in       string
		expected bool
	}{
		{"4242 4242 4242 4242", true},
		{"42-42 42-42 42-42 42-42", true},
		{"4242.4242.4242.4242", false},
		{"morethanfour", false},
	}
	for _, tt := range creditCardTests {
		if actual := panitizer.ContainsPAN(tt.in); tt.expected != actual {
			t.Errorf("Expected %v given %q but got %v", tt.expected, tt.in, actual)
		}
	}
}
