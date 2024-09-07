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

func XOR(hex1 string, hex2 string) string {
	bytes1, err1 := hex.DecodeString(hex1)
	bytes2, err2 := hex.DecodeString(hex2)

	if err1 != nil {
		panic(err1)
	}

	if err2 != nil {
		panic(err2)
	}

	if len(bytes1) != len(bytes2) {
		panic("Byte arrays must have equal size")
	}

	output := make([]byte, len(bytes1))

	for i := range bytes1 {
		output[i] = bytes1[i] ^ bytes2[i]
	}

	return hex.EncodeToString(output)
}
