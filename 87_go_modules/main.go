package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	dec, err := decimal.NewFromString("136.02")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", dec)
}
