package set2

import (
	"bytes"
	"encoding/base64"
	"math/rand"

	"github.com/senecal-jjs/crypto-pals/util"
)

func Challenge14() {
	key := randomBytes(16)
	unknownString, err := base64.StdEncoding.DecodeString("Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK")
	util.PanicOnErr(err)
	unknownPrefix := randomBytes(rand.Intn(16))

	prefixSize := getPrefixSize(unknownPrefix, unknownString, key)

	// we want our attacker controlled bytes + unknown prefix bytes
	// to equal one less than the size of a block
	mbs := 0

	if prefixSize == 0 {
		mbs = BLOCK_SIZE - 1
	} else {
		mbs = BLOCK_SIZE - 1 - prefixSize
	}

	attackingBytes := append(unknownPrefix, bytes.Repeat([]byte("A"), mbs)...)

	crackByteAtATime(BLOCK_SIZE, unknownString, attackingBytes)
}

func getPrefixSize(unknownPrefix []byte, unknownString []byte, key []byte) int {
	dupes := 0
	count := 1

	// increase length of "myString" until we get a duplicate block
	// count mod block size should give the size of the unknown prefix
	for {
		if dupes > 0 {
			break
		}

		myString := bytes.Repeat([]byte("A"), count)
		input := bytes.Join([][]byte{unknownPrefix, myString, unknownString}, []byte(""))
		input = oracle(input, key)
		dupes = countDuplicates(util.ChunkByteArray(input, BLOCK_SIZE))

		count++
	}

	return BLOCK_SIZE - ((count - 1) % BLOCK_SIZE)
}
