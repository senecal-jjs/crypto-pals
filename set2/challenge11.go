package set2

import (
	"bytes"
	crand "crypto/rand"
	"fmt"
	mrand "math/rand"

	"github.com/senecal-jjs/crypto-pals/cryptom/aes"
	"github.com/senecal-jjs/crypto-pals/util"
)

func Challenge11() {
	cipherText := encryptWithRandomMode([]byte("yellow submarineyellow submarineyellow submarineyellow submarineyellow submarine"))
	mode := detectAesMode(cipherText, 16)

	fmt.Printf("Detected mode: %q\n", mode)
}

func detectAesMode(cipherText []byte, blockSize int) string {
	mode := "CBC"
	chunks := util.ChunkByteArray(cipherText, blockSize)
	dupes := countDuplicates(chunks)

	if dupes > 0 {
		mode = "ECB"
	} else {
		mode = "CBC"
	}

	return mode
}

func encryptWithRandomMode(plainText []byte) []byte {
	padByteCount := mrand.Intn(6) + 5
	newBytes := randomBytes(padByteCount)
	plainText = append(plainText, newBytes...)
	plainText = append(newBytes, plainText...)

	// randomly choose encryption mode
	mode := mrand.Intn(2)

	var cipherText []byte

	if mode == 0 {
		cipherText = aes.EncryptAesCBC(plainText, randomBytes(16), randomBytes(16))
		fmt.Println("Encrypted with CBC mode")
	} else {
		cipherText = aes.EncryptAesECB(plainText, randomBytes(16))
		fmt.Println("Encrypted with ECB mode")
	}

	return cipherText
}

func randomBytes(size int) []byte {
	bytes := make([]byte, size)
	_, err := crand.Read(bytes)
	util.PanicOnErr(err)
	return bytes
}

func countDuplicates(bufs [][]byte) int {
	repeats := 0

	for i, chunk1 := range bufs {
		for j, chunk2 := range bufs {
			if i != j {
				if bytes.Equal(chunk1, chunk2) {
					repeats++
				}
			}
		}
	}

	return repeats
}
