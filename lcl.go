package main

import (
	"fmt"
	"log"

	"github.com/johnernaut/lcl/parser"
)

func main() {
	originalKeys, err := parser.ParseLocaleFile("/Users/john/code/employee-ios/mobile/en.lproj/Localizable.strings")
	if err != nil {
		log.Fatalf("Error parsing original file: %s", err.Error())
	}

	for k := range originalKeys {
		fmt.Println(k)
	}
}
