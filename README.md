# panitizer (PAN sanitizer)

Go library (and command line tool) for scrubbing strings of PAN (Personal Account Number) data.

## Installation

```
go get github.com/NickPresta/panitizer
```

## Usage

```go
package main

import (
    "github.com/NickPresta/panitizer"
)

func main() {
    cardNumber := "4242424242424242"
    fmt.Printf(panitizer.Replace(cardNumber)) // Prints ************4242
}
```

## How Personal Account Numbers are identified


A string of digits is identified and masked as a credit card number if all of the following apply:

1. The digits follow the regular expression found [here](http://www.regular-expressions.info/creditcard.html), which covers Visa, MasterCard, American Express, Diners Club, Discover, and JCB numbers.
2. The digits pass the [Luhn Check](http://en.wikipedia.org/wiki/Luhn_algorithm).
3. The digits are separated only by spaces and/or dashes, or not separated at all. This allows for different sizes of groupings of digits.
Note: A number broken by a line break is not detected as a credit card number.

## Supported formats

`*` is used as the masking symbol in the following examples

| original | credit card detected?   | sanitized |
|----------|:-----------------------:|-----------|
| `4242 4242 4242 4242` | Yes | `**** **** **** 4242` |
| `4242-4242-4242-4242` | Yes | `**** **** **** 4242` |
| `42-42 42-42 42-42 42-42` | Yes | `**-** **-** **-** *2-42` |
| `4242.4242.4242.4242` | No | `4242.4242.4242.4242` |
| `4242x4242x4242x4242` | No | `4242x4242x4242x4242` |
| `4242 4242 4242 4242 4242 4242` | Yes | `**** **** **** 4242 4242 4242` |

Check out [the tests](./panitizer_test.go) for more examples.
