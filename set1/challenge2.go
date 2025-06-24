package set1 

import (
	"fmt"

	"github.com/senecal-jjs/crypto-pals/util"
	"encoding/hex"
)

func Challenge2() {
	bytes1, err1 := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	bytes2, err2 := hex.DecodeString("686974207468652062756c6c277320657965")

	if err1 != nil {
		panic(err1)
	}

	if err2 != nil {
		panic(err2)
	}

	xor := util.Xor(bytes1, bytes2)

	fmt.Println(hex.EncodeToString(xor))
}