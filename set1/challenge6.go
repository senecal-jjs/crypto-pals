package set1

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/senecal-jjs/crypto-pals/util"
)

func Challenge6() {
	file, err := os.ReadFile("set1/input6.txt")
	util.PanicOnErr(err)
	input := strings.Join(strings.Split(string(file), "\n"), "")

	cipherText, err := base64.StdEncoding.DecodeString(input)
	util.PanicOnErr(err)

	fmt.Println(string(breakCipher(cipherText)))
}

func breakCipher(cipherText []byte) []byte {
	keySize := getKeySize(cipherText)
	fmt.Printf("Key Size %d\n", keySize)

	key := findKey(cipherText, keySize)
	fmt.Printf("Key: %q\n", string(key))

	return decryptRepeatingKeyXor(cipherText, key)
}

func decryptRepeatingKeyXor(cipherText []byte, key []byte) []byte {
	fullKey := bytes.Repeat(key, len(cipherText))[:len(cipherText)]
	return util.Xor(fullKey, cipherText)
}

func findKey(cipherText []byte, keySize int) []byte {
	fullKey := []byte{}

	for i := range keySize {
		b := buildBuffer(cipherText, i, keySize)
		_, key, _ := decryptSingleByteXor(b)
		fullKey = append(fullKey, key)
	}

	return fullKey
}

func buildBuffer(buf []byte, offset, keySize int) []byte {
	r := make([]byte, 0, len(buf)/keySize)
	for i := offset; i < len(buf); i += keySize {
		r = append(r, buf[i])
	}
	return r
}

func getKeySize(buf []byte) int {
	curKeySize := 0
	minDist := -1.00

	for keySize := 2; keySize <= 40; keySize++ {
		dist := normalizedHammingDistance(buf, keySize)

		// fmt.Printf("Key Size: %d, Dist: %f\n", keySize, dist)

		if minDist == -1.00 || dist < minDist {
			curKeySize = keySize
			minDist = dist
		}
	}

	return curKeySize
}

// 4 blocks normalized hamming distance
func normalizedHammingDistance(buf []byte, keySize int) float64 {
	chunks := chunkByteArray(buf, keySize)[:4]
	distance := 0.00

	for i := range 4 {
		for j := range 4 {
			distance += float64(hammingDistance(chunks[i], chunks[j])) / float64(keySize)
		}
	}

	return float64(distance) / 4.00
}

func chunkByteArray(data []byte, chunkSize int) [][]byte {
	var chunks [][]byte

	for i := 0; i < len(data); i += chunkSize {
		end := i + chunkSize
		if end > len(data) {
			end = len(data)
		}
		chunks = append(chunks, data[i:end])
	}

	return chunks
}

func hammingDistance(buf1 []byte, buf2 []byte) int {
	bitString1 := ""

	for i := range len(buf1) {
		bitString1 += fmt.Sprintf("%08b", buf1[i])
	}

	bitString2 := ""

	for i := range len(buf1) {
		bitString2 += fmt.Sprintf("%08b", buf2[i])
	}

	maxLen := 0

	if len(bitString1) > len(bitString2) {
		maxLen = len(bitString1)
	} else {
		maxLen = len(bitString2)
	}

	editCount := 0

	for i := range maxLen {
		if bitString1[i] != bitString2[i] {
			editCount++
		}
	}

	editCount += int(math.Abs(float64(len(bitString1) - len(bitString2))))

	return editCount
}
