package set2

import (
	"encoding/hex"
	"fmt"

	"github.com/senecal-jjs/crypto-pals/encoding/pkcs7"
)

func Challenge9() {
	input := "YELLOW SUBMARINE"
	buf := pkcs7.Pad([]byte(input), 20)
	hexString := hex.EncodeToString(buf)
	fmt.Println(hexString)
	fmt.Println(string(buf))
}
