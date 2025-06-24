package set1

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/senecal-jjs/crypto-pals/util"
)

func Challenge5() {
	plainText := `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`
	key := "ICE"

	expectedCipher := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

	cipherText := encryptRepeatingKeyXor([]byte(key), []byte(plainText))

	fmt.Printf("encryptXor(%q, %q) = %q\n", plainText, key, cipherText)

	if cipherText != expectedCipher {
		panic("fail")
	}
}

func encryptRepeatingKeyXor(key []byte, plainText []byte) string {
	fullKey := bytes.Repeat(key, len(plainText))[:len(plainText)]
	return hex.EncodeToString(util.Xor(fullKey, plainText))
}
