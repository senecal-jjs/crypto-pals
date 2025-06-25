package set1

import (
	"encoding/hex"
	"fmt"

	"github.com/senecal-jjs/crypto-pals/util"
)

func Challenge3() {
	cipherText := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	cipherBytes, err := hex.DecodeString(cipherText)

	if err != nil {
		panic(err)
	}

	_, _, plainText := decryptSingleByteXor(cipherBytes)

	fmt.Println(plainText)
}

func decryptSingleByteXor(cipher []byte) (int, byte, string) {
	currentScore := -1
	currentPlainText := ""
	currentKey := byte(0)

	// loop over all possible bytes
	for i := 0; i <= 255; i++ {
		keySlice := make([]byte, len(cipher))

		// construct a full length key
		for j := range keySlice {
			keySlice[j] = byte(i)
		}

		// xor the full key with the cipher text
		plainBytes := util.Xor(keySlice, cipher)
		newScore := score(plainBytes)

		if newScore > currentScore {
			currentScore = newScore
			currentPlainText = string(plainBytes)
			currentKey = byte(i)
		}
	}

	return currentScore, currentKey, currentPlainText
}

func score(buf []byte) int {
	count := 0

	// checking that the output is an english character is enough to recover the key
	// frequency analysis would be a fancier (not needed) method
	for i := 0; i < len(buf); i++ {
		if buf[i] >= 'A' && buf[i] <= 'Z' {
			count++
		} else if buf[i] >= 'a' && buf[i] <= 'z' {
			count++
		} else if buf[i] == ' ' {
			count++
		}
	}

	return count
}
