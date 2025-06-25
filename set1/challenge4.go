package set1

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/senecal-jjs/crypto-pals/util"
)

func Challenge4() {
	file, err := os.ReadFile("set1/input4.txt")

	util.PanicOnErr(err)

	inputs := strings.Split(string(file), "\n")
	bestScore := -1
	bestText := ""
	cipherHex := ""

	for i := range inputs {
		cipherBytes, err := hex.DecodeString(inputs[i])

		util.PanicOnErr(err)

		score, _, plainText := decryptSingleByteXor(cipherBytes)

		if score > bestScore {
			bestScore = score
			bestText = plainText
			cipherHex = inputs[i]
		}
	}

	fmt.Printf("Plain Text: %q\nHex String: %q\n", bestText, cipherHex)
}
