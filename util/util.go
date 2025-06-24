package util

import (
	"encoding/base64"
	"encoding/hex"
)

func HexToBase64(input string) string {
	bytes, err := hex.DecodeString(input)

	if err != nil {
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(bytes)
}

// Xor returns a new buffer with content set to bytes1 xor bytes2
func Xor(bytes1, bytes2 []byte) []byte {
	if len(bytes1) != len(bytes2) {
		panic("Byte arrays must have equal size")
	}

	output := make([]byte, len(bytes1))

	for i := range bytes1 {
		output[i] = bytes1[i] ^ bytes2[i]
	}

	return output
}
