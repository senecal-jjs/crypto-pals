package set2

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/senecal-jjs/crypto-pals/cryptom/aes"
	"github.com/senecal-jjs/crypto-pals/util"
)

var key = randomBytes(16)

func Challenge12() {
	blockSize := detectBlockSize()
	mode := detectAesMode(oracle(bytes.Repeat([]byte("A"), 32), key), blockSize)
	fmt.Printf("AES mode: %q\n", mode)
	fmt.Printf("Block size: %d\n", blockSize)

	unknownString, err := base64.StdEncoding.DecodeString("Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK")
	util.PanicOnErr(err)

	crack(blockSize, unknownString)
}

func crack(blockSize int, unknownString []byte) {
	solve := []byte{}

	for i := range len(unknownString) {
		solve = append(solve, crackByte(blockSize, unknownString[i:]))
	}

	fmt.Println(string(solve))
}

func crackByte(blockSize int, unknownString []byte) byte {
	prefix := bytes.Repeat([]byte("A"), blockSize-1)
	dict := constructDictionary(prefix, 16)

	plainText := append(prefix, unknownString...)
	cipherText := oracle(plainText, key)
	cipherBlock := cipherText[:blockSize]

	return byte(dict[hex.EncodeToString(cipherBlock)])
}

func constructDictionary(prefix []byte, blockSize int) map[string]int {
	m := make(map[string]int)

	for i := range 256 {
		input := append(prefix, byte(i))
		m[hex.EncodeToString(oracle(input, key)[:blockSize])] = i
	}

	return m
}

func detectBlockSize() int {
	// we know block size but assume it maxes out at 48 / 2 bytes
	for i := 2; i < 48; i++ {
		plainText := bytes.Repeat([]byte("A"), i)
		cipherText := oracle(plainText, key)

		if bytes.Equal(cipherText[0:i/2], cipherText[i/2:i]) {
			return i / 2
		}
	}

	panic("Could not detect block size")
}

func oracle(plainText []byte, key []byte) []byte {
	cipherText := aes.EncryptAesECB(plainText, key)
	return cipherText
}
