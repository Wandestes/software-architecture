package main

import (
	"fmt"
	"log"
)

func main() {
	expression := "+ 1 * 2 3"
	result, err := ConvertPrefixToLisp(expression)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
