package set2

import (
	"fmt"

	"github.com/senecal-jjs/crypto-pals/encoding/pkcs7"
)

func Challenge15() {
	valid := "ICE ICE BABY\x04\x04\x04\x04"
	invalid := "ICE ICE BABY\x05\x05\x05\x05"

	fmt.Println([]byte(valid))

	fmt.Println(pkcs7.Unpad([]byte(valid), 16))
	fmt.Println(pkcs7.Unpad([]byte(invalid), 16))
}
