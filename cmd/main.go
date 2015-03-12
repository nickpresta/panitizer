package main

import (
	"fmt"
	"os"

	"github.com/NickPresta/panitizer"
)

func main() {
	for _, arg := range os.Args[1:] {
		cleansedContent := panitizer.Replace(arg)
		fmt.Printf("%q => %q\n", arg, cleansedContent)
	}
}
