package set1

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"github.com/senecal-jjs/crypto-pals/util"
)

func Challenge7() {
	file, err := os.ReadFile("set1/input7.txt")
	util.PanicOnErr(err)
	input := strings.Join(strings.Split(string(file), "\n"), "")

	cipherText, err := base64.StdEncoding.DecodeString(input)
	util.PanicOnErr(err)

	key := []byte("YELLOW SUBMARINE")
	plainText := decryptAesECB(cipherText, key)

	fmt.Println("DECRYPTED: ", string(plainText))
}

func decryptAesECB(cipherText []byte, key []byte) []byte {
	cipher, err := aes.NewCipher(key)
	util.PanicOnErr(err)

	plainText := []byte{}
	r := make([]byte, len(cipherText))

	cipherChunks := chunkByteArray(cipherText, len(key))

	for i := range cipherChunks {
		cipher.Decrypt(r, cipherChunks[i])
		plainText = append(plainText, r...)
	}

	return plainText
}
