package set2

import (
	"fmt"
	"strings"

	"github.com/senecal-jjs/crypto-pals/cryptom/aes"
	"github.com/senecal-jjs/crypto-pals/encoding/pkcs7"
)

var BLOCK_SIZE = 16

func Challenge13() {
	cutAndPaste()
}

func cutAndPaste() {
	key := randomBytes(BLOCK_SIZE)

	// craft a block that is just encrypted string "admin" + padding to fill the block
	// email=xxxxxxxxxxadmin<----padding--->&uid=10&role=user
	//                 ^ start 2nd block    ^ start 3rd block
	remaining := BLOCK_SIZE - len("email=")
	dummyEmail := strings.Repeat("x", remaining)
	dummyEmail += string(pkcs7.Pad([]byte("admin"), BLOCK_SIZE))

	eProfile := profileFor(dummyEmail, key)

	// encrypted form of admin should be the second block of cipher text
	elevatedRole := eProfile[BLOCK_SIZE : BLOCK_SIZE*2]

	// craft an email that places "role=" at the end of a block
	remaining = (BLOCK_SIZE * 2) - len("email=&uid=10&role=")
	dummyEmail = strings.Repeat("x", remaining)

	eProfile = profileFor(dummyEmail, key)

	// cut and paste the 3rd block to replace the role
	elevatedCookie := append(eProfile[:BLOCK_SIZE*2], elevatedRole...)

	decryptedElevatedCookie := aes.DecryptAesECB(elevatedCookie, key)

	fmt.Println(string(decryptedElevatedCookie))
}

func decodeCookie(cookie string) map[string]string {
	kvs := strings.Split(cookie, "&")
	m := make(map[string]string)

	for _, kv := range kvs {
		split := strings.Split(kv, "=")
		m[split[0]] = split[1]
	}

	return m
}

func encodeCookie(cookie map[string]string) string {
	encodedCookie := ""

	for k, v := range cookie {
		encodedCookie = encodedCookie + fmt.Sprintf("%s=%s&", k, v)
	}

	return strings.Trim(encodedCookie, "&")
}

func profileFor(email string, key []byte) []byte {
	cleanEmail := strings.ReplaceAll(email, "&", "")
	cleanEmail = strings.ReplaceAll(cleanEmail, "=", "")

	cookie := fmt.Sprintf("email=%s&uid=10&role=user", cleanEmail)

	return aes.EncryptAesECB([]byte(cookie), key)
}
