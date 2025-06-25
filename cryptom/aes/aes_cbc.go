package aes

import (
	"crypto/aes"

	"github.com/senecal-jjs/crypto-pals/encoding/pkcs7"
	"github.com/senecal-jjs/crypto-pals/util"
)

// Cipher Block Chaining mode

func DecryptAesCBC(buf []byte, iv []byte, key []byte) []byte {
	cipher, err := aes.NewCipher(key)
	util.PanicOnErr(err)

	plainText := []byte{}
	ciMinus1 := iv

	cipherChunks := util.ChunkByteArray(buf, len(key))

	for i := range cipherChunks {
		output := make([]byte, cipher.BlockSize())
		cipher.Decrypt(output, cipherChunks[i])
		plainText = append(plainText, util.Xor(output, ciMinus1)...)
		ciMinus1 = cipherChunks[i]
	}

	return plainText
}

func EncryptAesCBC(buf []byte, iv []byte, key []byte) []byte {
	cipher, err := aes.NewCipher(key)
	util.PanicOnErr(err)

	cipherText := []byte{}
	ci := iv

	plainChunks := util.ChunkByteArray(pkcs7.Pad(buf, len(key)), len(key))

	for _, chunk := range plainChunks {
		output := make([]byte, cipher.BlockSize())
		cipher.Encrypt(output, util.Xor(ci, chunk))
		cipherText = append(cipherText, output...)
		ci = output
	}

	return cipherText
}
