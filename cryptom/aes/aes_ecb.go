package aes

import (
	"crypto/aes"

	"github.com/senecal-jjs/crypto-pals/encoding/pkcs7"
	"github.com/senecal-jjs/crypto-pals/util"
)

// Electronic Codebook Mode

func DecryptAesECB(buf []byte, key []byte) []byte {
	cipher, err := aes.NewCipher(key)
	util.PanicOnErr(err)

	plainText := []byte{}
	r := make([]byte, len(buf))

	cipherChunks := util.ChunkByteArray(buf, cipher.BlockSize())

	for i := range cipherChunks {
		cipher.Decrypt(r, cipherChunks[i])
		plainText = append(plainText, r...)
	}

	return plainText
}

func EncryptAesECB(buf []byte, key []byte) []byte {
	cipher, err := aes.NewCipher(key)
	util.PanicOnErr(err)

	cipherText := []byte{}
	r := make([]byte, len(key))

	plainChunks := util.ChunkByteArray(pkcs7.Pad(buf, len(key)), len(key))

	for i := range plainChunks {
		cipher.Encrypt(r, plainChunks[i])
		cipherText = append(cipherText, r...)
	}

	return cipherText
}
