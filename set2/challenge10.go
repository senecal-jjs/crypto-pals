package set2

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"github.com/senecal-jjs/crypto-pals/cryptom/aes"
	"github.com/senecal-jjs/crypto-pals/encoding/pkcs7"
	"github.com/senecal-jjs/crypto-pals/util"
)

func Challenge10() {
	file, err := os.ReadFile("set2/input10.txt")
	util.PanicOnErr(err)
	input := strings.Join(strings.Split(string(file), "\n"), "")
	buf, err := base64.StdEncoding.DecodeString(input)
	util.PanicOnErr(err)

	key := []byte("YELLOW SUBMARINE")
	iv := make([]byte, 16)

	buf = aes.DecryptAesCBC(buf, iv, key)
	buf, err = pkcs7.Unpad(buf, 16)
	util.PanicOnErr(err)

	fmt.Println(string(buf))
}
