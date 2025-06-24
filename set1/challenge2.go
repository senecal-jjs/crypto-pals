package set1 

import (
	"fmt"

	"github.com/senecal-jjs/crypto-pals/util"
	"encoding/hex"
)

func Challenge2() {
	input1 := "1c0111001f010100061a024b53535009181c"
	input2 := "686974207468652062756c6c277320657965"

	bytes1, err1 := hex.DecodeString(input1)
	bytes2, err2 := hex.DecodeString(input2)

	if err1 != nil {
		panic(err1)
	}

	if err2 != nil {
		panic(err2)
	}

	xor := hex.EncodeToString(util.Xor(bytes1, bytes2))

	fmt.Printf("FixedXor(%q, %q) = %q\n", input1, input2, xor)

	if xor != "746865206b696420646f6e277420706c6179" {
		panic("fail")
	}

	fmt.Println(xor)
}