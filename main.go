package main

import (
	"log"
	"os"
)

func main() {
	bytes, err := os.ReadFile("test.lg")
	if err != nil {
		log.Fatal(err)
	}

	tokens := Tokenize(string(bytes))

	for _, token := range tokens {
		token.Debug()
	}
}
