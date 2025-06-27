package set2

import (
	"fmt"
	"strings"

	"github.com/senecal-jjs/crypto-pals/cryptom/aes"
	"github.com/senecal-jjs/crypto-pals/encoding/pkcs7"
	"github.com/senecal-jjs/crypto-pals/util"
)

func Challenge16() {
	key := randomBytes(BLOCK_SIZE)
	iv := randomBytes(BLOCK_SIZE)
	input := "_admin-true"
	cipherText := encrypt(input, iv, key)

	// determine the replacement ciphertext to achieved the desired bit flip from "_ -> ;" and "- > ="
	// C_i = P*_i+1 ^ (P_i+1 ^ C_i)
	semicolon := util.Xor([]byte(";"), util.Xor([]byte("_"), cipherText[16:17]))
	equality := util.Xor([]byte("="), util.Xor([]byte("-"), cipherText[22:23]))

	copy(cipherText[16:], semicolon)
	copy(cipherText[22:], equality)

	fmt.Println(decrypt(cipherText, iv, key))
}

func encrypt(input string, iv []byte, key []byte) []byte {
	prefix := "comment1=cooking%20MCs;userdata="
	fmt.Println(len(prefix))
	suffix := ";comment2=%20like%20a%20pound%20of%20bacon"
	input = encode(input)
	input = prefix + input + suffix
	inputBytes := pkcs7.Pad([]byte(input), BLOCK_SIZE)
	return aes.EncryptAesCBC(inputBytes, iv, key)
}

func decrypt(cipherText []byte, iv []byte, key []byte) bool {
	plainText := string(aes.DecryptAesCBC(cipherText, iv, key))

	fmt.Println(plainText)

	return strings.Contains(string(plainText), ";admin=true;")
}

func encode(input string) string {
	input = strings.ReplaceAll(input, ";", "")
	input = strings.ReplaceAll(input, "=", "")
	return input
}

func decode(input string) string {
	input = strings.ReplaceAll(input, "\";\"", ";")
	input = strings.ReplaceAll(input, "\"=\"", "=")
	return input
}
