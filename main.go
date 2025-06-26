package main

import (
	"os"

	"github.com/senecal-jjs/crypto-pals/set1"
	"github.com/senecal-jjs/crypto-pals/set2"
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

	case "4":
		set1.Challenge4()

	case "5":
		set1.Challenge5()

	case "6":
		set1.Challenge6()

	case "7":
		set1.Challenge7()

	case "8":
		set1.Challenge8()

	case "9":
		set2.Challenge9()

	case "10":
		set2.Challenge10()

	case "11":
		set2.Challenge11()

	case "12":
		set2.Challenge12()

	case "13":
		set2.Challenge13()

	case "14":
		set2.Challenge14()

	default:
		set1.Challenge1()
	}
}
