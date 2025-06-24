package set1

import (
	"fmt"

	"github.com/senecal-jjs/crypto-pals/util"
)

func Challenge1() {
	base64 := util.HexToBase64(
		"49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d",
	)
	fmt.Println(base64)
}
