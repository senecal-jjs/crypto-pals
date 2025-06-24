package main

import (
	"os"

	"github.com/senecal-jjs/crypto-pals/set1"
)

func main() {
	challenge := ""

	if len(os.Args) >= 2 {
		challenge = os.Args[1]
	}

	switch challenge {
	case "1":
		set1.Challenge1()

	case "2":
		set1.Challenge2()

	case "3":
		set1.Challenge3()

	default:
		set1.Challenge1()
	}
}
